package tester

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/url"
	"os"
	"strings"
	"time"

	"github.com/kr/pretty"
	"github.com/pkg/errors"
	backo "github.com/segmentio/backo-go"
	"github.com/segmentio/events"
	"github.com/segmentio/library-e2e-tester/webhook"
)

// ErrMissingInWebhook is returned when the message could not be found in the webhook.
var ErrMissingInWebhook = errors.New("message not found in webhook")

// ErrNotMatchedInWebhook is returned when the message is found in the webhook but the contents do not match.
var ErrNotMatchedInWebhook = errors.New("message found in webhook, but not matched")

// T runs our tests.
type T struct {
	SegmentWriteKey     string
	WebhookBucket       string
	WebhookAuthUsername string
	ReportFileName      string
	FailFast            bool // disable running additional tests after any test fails
	SkipMessages        []string
}

// Check whether the tester should skip messages of the given message type
func (t *T) SkipMessageType(msgType string) bool {
	for _, skipType := range t.SkipMessages {
		// mesasge type comparison should be case-insensitive
		if strings.EqualFold(skipType, msgType) {
			return true
		}
	}
	return false
}

// Test invokes the test binaries.
func (t *T) Test(invoker Invoker) error {
	var res error
	ctx := context.Background()

	var reportWriter io.Writer
	if t.ReportFileName != "" {
		if reportWriter, res = os.Create(t.ReportFileName); res != nil {
			return res
		}
	} else {
		reportWriter = ioutil.Discard
	}

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

			var testrun TestRun
			testrun.Start(fixture)

			f, err := Asset("fixtures/" + dir + "/" + fixture)
			if err != nil {
				testrun.Error("could not read fixture " + fixture)
				if t.FailFast {
					return errors.Wrap(err, "could not read fixture "+fixture)
				}
				res = err
			}

			var buf bytes.Buffer
			if err := producer.Produce(ctx, bytes.NewReader(f), &buf); err != nil {
				testrun.Error(fixture + ": could not produce messages")
				if t.FailFast {
					return errors.Wrap(err, fixture+": could not produce messages")
				}
				res = err
			}

			var msg map[string]interface{}
			if err := json.Unmarshal(buf.Bytes(), &msg); err != nil {
				testrun.Error(fixture + ": could not parse json")
				if t.FailFast {
					return errors.Wrap(err, fixture+": could not parse json")
				}
				res = err
			}

			msgType := msg["type"].(string)

			if t.SkipMessageType(msgType) {
				events.Log("Skip message for fixture %{fixture}v", fixture)
				testrun.Skip()
				testrun.Print(reportWriter)
				continue
			}

			args := []string{
				"--writeKey=" + t.SegmentWriteKey,
				"--type=" + msgType,
				"--userId=" + msg["userId"].(string),
			}
			anonId, ok := msg["anonymousId"]
			if ok {
				args = append(args, "--anonymousId="+anonId.(string))
			}
			for _, key := range []string{"context", "integrations"} {
				val, ok := msg[key]
				if ok {
					jsonStr, err := json.Marshal(val)
					if err == nil {
						args = append(args, "--"+key+"="+url.PathEscape(string(jsonStr)))
					}
				}
			}

			switch msgType {
			case "track":
				properties, err := json.Marshal(msg["properties"])
				if err != nil {
					testrun.Error("could not marshal properties")
					if t.FailFast {
						return errors.Wrap(err, "could not marshal properties")
					}
					res = err
				}
				args = append(args, "--event="+msg["event"].(string), "--properties="+string(properties))
			case "screen":
				fallthrough
			case "page":
				properties, err := json.Marshal(msg["properties"])
				if err != nil {
					testrun.Error("could not marshal properties")
					if t.FailFast {
						return errors.Wrap(err, "could not marshal properties")
					}
					res = err
				}
				args = append(args, "--name="+msg["name"].(string), "--properties="+string(properties))
			case "identify":
				traits, err := json.Marshal(msg["traits"])
				if err != nil {
					testrun.Error("could not marshal traits")
					if t.FailFast {
						return errors.Wrap(err, "could not marshal traits")
					}
					res = err
				}
				args = append(args, "--traits="+string(traits))
			case "group":
				traits, err := json.Marshal(msg["traits"])
				if err != nil {
					testrun.Error("could not marshal traits")
					if t.FailFast {
						return errors.Wrap(err, "could not marshal traits")
					}
					res = err
				}
				args = append(args, "--traits="+string(traits), "--groupId="+msg["groupId"].(string))
			case "alias":
				args = append(args, "--previousId="+msg["previousId"].(string))
			default:
				panic("unsupported type: " + msgType)
			}

			events.Debug("invoking library for fixture %{fixture}v with %{args}v", fixture, args)

			if err := invoker(ctx, args...); err != nil {
				testrun.Error("could not invoke command")
				if t.FailFast {
					return errors.Wrap(err, "could not invoke command")
				}
				res = err
			}

			events.Log("sent message for fixture %{fixture}v", fixture)

			if err := t.testMessage(msg); err != nil {
				testrun.Fail(err.Error())
				testrun.AddDetails(string(buf.Bytes()))
				if t.FailFast {
					return err
				}
				res = err
			} else {
				testrun.End(TEST_PASS)
			}

			testrun.Print(reportWriter)
		}
	}

	return res
}

func (t *T) testMessage(msg map[string]interface{}) error {
	var key string
	msgType := msg["type"].(string)

	switch msgType {
	case "track":
		key = "properties"
	case "page":
		key = "properties"
	case "screen":
		key = "properties"
	case "identify":
		key = "traits"
	case "group":
		key = "traits"
	case "alias":
		key = "context"
	default:
		panic("unsupported type: " + msgType)
	}

	expectedID, _ := pickID(msg, key)

	backo := backo.NewBacko(2*time.Second, 2, 1, 5*time.Second)
	ticker := backo.NewTicker()
	timeout := time.After(3 * time.Minute)
	for {
		select {
		case <-ticker.C:
			events.Debug("searching for id %{id}v", expectedID)
			webhookMsg, err := t.findMessageInWebhook(expectedID, key)
			if err != nil {
				continue
			}

			if SegmentEqual(webhookMsg, msg) {
				events.Debug("matched: %{id}v", expectedID)
				return nil
			}

			events.Log("found id %{id}v, but could not match content", expectedID)
			pretty.Fdiff(os.Stdout, webhookMsg, msg)
			return ErrNotMatchedInWebhook
		case <-timeout:
			events.Log("didn't find message %{id}v in webhook after timeout", expectedID)
			return ErrMissingInWebhook
		}
	}
}

func (t *T) findMessageInWebhook(expectedID, key string) (map[string]interface{}, error) {
	msgs, err := webhook.GetWebhookMessages(t.WebhookBucket, t.WebhookAuthUsername)
	if err != nil {
		return nil, errors.Wrap(err, "could not get webhook messages")
	}

	for _, msg := range msgs {
		gotID, ok := pickID(msg, key)
		if ok && gotID == expectedID {
			return msg, nil
		}
	}

	return nil, fmt.Errorf("could not find message with id: %s", expectedID)
}
