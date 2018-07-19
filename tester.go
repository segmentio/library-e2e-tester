package tester

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/kr/pretty"
	"github.com/pkg/errors"
	backo "github.com/segmentio/backo-go"
	"github.com/segmentio/events"
)

// ErrMissingInRunscope is returned when the message could not be found in Runscope.
var ErrMissingInRunscope = errors.New("message not found in runscope")

// ErrNotMatchedInRunscope is returned when the message is found in runscope but the contents do not match.
var ErrNotMatchedInRunscope = errors.New("message found in runscope, but not matched")

// T runs our tests.
type T struct {
	SegmentWriteKey     string
	WebhookBucket       string
	WebhookAuthUsername string
}

// Test invokes the test binaries.
func (t *T) Test(invoker Invoker) error {
	ctx := context.Background()

	fixturesDirectory, err := AssetDir("fixtures")
	if err != nil {
		return errors.Wrap(err, "could not read fixtures directory")
	}

	producer := NewTemplatedProducer()

	for _, dir := range fixturesDirectory {
		fixtures, err := AssetDir("fixtures/" + dir)
		if err != nil {
			return errors.Wrap(err, "could not read fixtures directory")
		}

		events.Debug("running %{count}d fixtures in %{directory}v", len(fixtures), dir)

		for _, fixture := range fixtures {
			events.Debug("running %{fixture}v", fixture)

			f, err := Asset("fixtures/" + dir + "/" + fixture)
			if err != nil {
				return errors.Wrap(err, "could not read fixture")
			}

			var buf bytes.Buffer
			if err := producer.Produce(ctx, bytes.NewReader(f), &buf); err != nil {
				return errors.Wrap(err, "could not produce messages")
			}

			var msg map[string]interface{}
			if err := json.Unmarshal(buf.Bytes(), &msg); err != nil {
				return errors.Wrap(err, "could not parse json")
			}

			msgType := msg["type"].(string)

			args := []string{
				"--writeKey=" + t.SegmentWriteKey,
				"--type=" + msgType,
				"--userId=" + msg["userId"].(string),
			}

			switch msgType {
			case "track":
				properties, err := json.Marshal(msg["properties"])
				if err != nil {
					return errors.Wrap(err, "could not marshal properties")
				}
				args = append(args, "--event="+msg["event"].(string), "--properties="+string(properties))
			case "page":
				properties, err := json.Marshal(msg["properties"])
				if err != nil {
					return errors.Wrap(err, "could not marshal properties")
				}
				args = append(args, "--name="+msg["name"].(string), "--properties="+string(properties))
			case "identify":
				traits, err := json.Marshal(msg["traits"])
				if err != nil {
					return errors.Wrap(err, "could not marshal traits")
				}
				args = append(args, "--traits="+string(traits))
			case "group":
				traits, err := json.Marshal(msg["traits"])
				if err != nil {
					return errors.Wrap(err, "could not marshal traits")
				}
				args = append(args, "--traits="+string(traits), "--groupId="+msg["groupId"].(string))
			default:
				panic("unsupported type: " + msgType)
			}

			events.Debug("invoking library for fixture %{fixture}v with %{args}v", fixture, args)

			if err := invoker(ctx, args...); err != nil {
				return errors.Wrap(err, "could not invoke command")
			}

			events.Log("sent mesage for fixture %{fixture}v", fixture)

			if err := t.testMessage(msg); err != nil {
				return err
			}
		}
	}

	return nil
}

func (t *T) testMessage(msg map[string]interface{}) error {
	var key string
	msgType := msg["type"].(string)

	switch msgType {
	case "track":
		key = "properties"
	case "page":
		key = "properties"
	case "identify":
		key = "traits"
	case "group":
		key = "traits"
	default:
		panic("unsupported type: " + msgType)
	}

	expectedID, _ := pickID(msg, key)

	backo := backo.NewBacko(2*time.Second, 2, 1, 5*time.Second)
	ticker := backo.NewTicker()
	timeout := time.After(1 * time.Minute)
	for {
		select {
		case <-ticker.C:
			events.Debug("searching for id %{id}v", expectedID)
			runscopeMsg, err := t.findMessageInRunscope(expectedID, key)
			if err != nil {
				continue
			}

			if SegmentEqual(runscopeMsg, msg) {
				events.Debug("matched: %{id}v", expectedID)
				return nil
			}

			events.Log("found id %{id}v, but could not match content", expectedID)
			pretty.Fdiff(os.Stdout, runscopeMsg, msg)
			return ErrNotMatchedInRunscope
		case <-timeout:
			events.Log("didn't find message %{id}v in runsope after timeout", expectedID)
			return ErrMissingInRunscope
		}
	}
}

func (t *T) findMessageInRunscope(expectedID, key string) (map[string]interface{}, error) {
	msgs, err := runscopeMessages(t.WebhookBucket, t.WebhookAuthUsername)
	if err != nil {
		return nil, errors.Wrap(err, "could not get runscope messages")
	}

	for _, msg := range msgs {
		gotID, ok := pickID(msg, key)
		if ok && gotID == expectedID {
			return msg, nil
		}
	}

	return nil, fmt.Errorf("could not find message with id: %s", expectedID)
}
