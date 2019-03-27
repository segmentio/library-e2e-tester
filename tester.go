package tester

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"regexp"
	"strings"
	"sync"
	"time"

	"github.com/f2prateek/semaphore"
	"github.com/pkg/errors"
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
	Output              io.Writer
	SkipRegex           string
	Timeout             time.Duration
	Concurrency         int
}

// shouldSkipFixture returns true if the tester should skip the given fixture.
func (t *T) shouldSkipFixture(fixture string) bool {
	// don't skip fixture if the provided regex is empty.
	if strings.TrimSpace(t.SkipRegex) == "" {
		return false
	}

	matched, err := regexp.MatchString(t.SkipRegex, fixture)
	if err != nil {
		events.Log("error matching %{fixture}s to %{regex}s", fixture, t.SkipRegex)
		return false
	}
	if matched {
		return true
	}
	return false
}

// Test invokes the test binaries.
func (t *T) Test(invoker Invoker) error {
	ctx := context.Background()

	fixturesDirectory, err := AssetDir("fixtures")
	if err != nil {
		return errors.Wrap(err, "could not read fixtures directory")
	}

	// mutex to restrict test reports to write one at a time.
	var outputMutex sync.Mutex

	s := semaphore.New(t.Concurrency)
	var wg sync.WaitGroup
	errC := make(chan error, len(fixturesDirectory))

	for _, dir := range fixturesDirectory {
		fixtures, err := AssetDir("fixtures/" + dir)
		if err != nil {
			return errors.Wrap(err, "could not read fixtures directory")
		}

		events.Debug("running %{count}d fixtures in %{directory}v", len(fixtures), dir)

		for _, fixture := range fixtures {
			wg.Add(1)
			s.Acquire(1)

			go func(directory, fixture string) {
				defer wg.Done()
				defer s.Release(1)

				var buf bytes.Buffer
				if err := t.runTestCase(ctx, directory, fixture, invoker, &buf); err != nil {
					events.Log("fixture %{fixture}s: %{error}v", fixture, err)
					errC <- errors.Wrap(err, "fixture: "+fixture)
				}

				outputMutex.Lock()
				if _, err := io.Copy(t.Output, &buf); err != nil {
					events.Log("error writing test result to output: %{error}v", err)
				}
				outputMutex.Unlock()
			}(dir, fixture)
		}
	}

	wg.Wait()
	close(errC)

	var testErrors []error
	for err := range errC {
		testErrors = append(testErrors, err)
	}
	if len(testErrors) == 0 {
		return nil
	}
	return fmt.Errorf("%v", testErrors)
}

func (t *T) runTestCase(ctx context.Context, directory, fixture string, invoker Invoker, output io.Writer) error {
	events.Debug("running %{fixture}v", fixture)

	testrun := NewTestRun(fixture, output)
	testrun.Start()

	// testError reports a test error with the given reason, and returns the error wrapped with the reason.
	testError := func(err error, reason string) error {
		testrun.Error(reason)
		return errors.Wrap(err, reason)
	}

	f, err := Asset("fixtures/" + directory + "/" + fixture)
	if err != nil {
		return testError(err, "could not read fixture")
	}

	producer := NewTemplatedProducer()
	var buf bytes.Buffer
	if err := producer.Produce(ctx, bytes.NewReader(f), &buf); err != nil {
		return testError(err, "could not produce messages")
	}

	var msg map[string]interface{}
	if err := json.Unmarshal(buf.Bytes(), &msg); err != nil {
		return testError(err, "could not parse json")
	}

	msgType := msg["type"].(string)

	if t.shouldSkipFixture(fixture) {
		events.Debug("skip fixture %{fixture}v", fixture)
		testrun.Skip()
		return nil
	}

	args := []string{
		"--writeKey=" + t.SegmentWriteKey,
		"--type=" + msgType,
		"--userId=" + msg["userId"].(string),
	}
	if anonymousID, ok := msg["anonymousId"]; ok {
		args = append(args, "--anonymousId="+anonymousID.(string))
	}
	for _, key := range []string{"context", "integrations"} {
		val, ok := msg[key]
		if ok {
			jsonStr, err := json.Marshal(val)
			if err == nil {
				args = append(args, "--"+key+"="+string(jsonStr))
			}
		}
	}

	switch msgType {
	case "track":
		properties, err := json.Marshal(msg["properties"])
		if err != nil {
			return testError(err, "could not marshal properties")
		}
		args = append(args, "--event="+msg["event"].(string), "--properties="+string(properties))
	case "screen":
		fallthrough
	case "page":
		properties, err := json.Marshal(msg["properties"])
		if err != nil {
			return testError(err, "could not marshal properties")
		}
		args = append(args, "--name="+msg["name"].(string), "--properties="+string(properties))
	case "identify":
		traits, err := json.Marshal(msg["traits"])
		if err != nil {
			return testError(err, "could not marshal traits")
		}
		args = append(args, "--traits="+string(traits))
	case "group":
		traits, err := json.Marshal(msg["traits"])
		if err != nil {
			return testError(err, "could not marshal traits")
		}
		args = append(args, "--traits="+string(traits), "--groupId="+msg["groupId"].(string))
	case "alias":
		args = append(args, "--previousId="+msg["previousId"].(string))
	default:
		panic("unsupported type: " + msgType)
	}

	events.Debug("invoking library for fixture %{fixture}v with %{args}v", fixture, args)

	if err := invoker(ctx, args...); err != nil {
		return testError(err, "could not invoke command")
	}

	events.Debug("sent message for fixture %{fixture}v", fixture)

	if err := t.testMessage(msg); err != nil {
		testrun.Fail(err.Error(), string(buf.Bytes()))
		return errors.Wrap(err, "could not test message")
	}

	testrun.Pass()
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

	ticker := time.NewTicker(5 * time.Second)
	timeout := time.After(t.Timeout)
	for {
		select {
		case <-ticker.C:
			events.Debug("searching for id %{id}v", expectedID)
			webhookMsg, err := t.findMessageInWebhook(expectedID, key)
			if err != nil {
				continue
			}

			equal, diff := SegmentEqual(webhookMsg, msg)
			if equal {
				events.Debug("matched: %{id}v", expectedID)
				return nil
			}

			events.Debug("found id %{id}v, but could not match content", expectedID)
			return errors.Wrap(ErrNotMatchedInWebhook, strings.Join(diff, ","))
		case <-timeout:
			events.Debug("didn't find message %{id}v in webhook after timeout", expectedID)
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
