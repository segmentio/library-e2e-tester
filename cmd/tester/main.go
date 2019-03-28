package main

import (
	"fmt"
	"os"
	"time"

	"github.com/segmentio/conf"
	"github.com/segmentio/events"
	"github.com/segmentio/events/text"
	tester "github.com/segmentio/library-e2e-tester"
)

var Version = "dev"

// Config represents the options that can be supplied to the harness.
type Config struct {
	Path                string        `conf:"path"                     help:"path to the library binary" validate:"nonzero"`
	SegmentWriteKey     string        `conf:"segment-write-key"        help:"write key for the Segment project to send data to" validate:"nonzero"`
	WebhookBucket       string        `conf:"webhook-bucket"           help:"webhook bucket the Segment project sends data to" validate:"nonzero"`
	WebhookAuthUsername string        `conf:"webhook-auth-username"    help:"basic auth username for the webhook bucket the Segment project sends data to" validate:"nonzero"`
	Skip                string        `conf:"skip"                     help:"skip fixtures matching the regex. providing an empty string will skip nothing."`
	Timeout             time.Duration `conf:"timeout"                  help:"if a message does not appear in the webhook within this duration, give up. the default is 5 minutes."`
	Debug               bool          `conf:"debug"                    help:"enable debug logging"`
	Concurrency         int           `conf:"concurrency"              help:"allow upto n concurrent tests to run simultaneously. the default is to run 1 test at a time."`
}

func main() {
	config := Config{
		Timeout:     5 * time.Minute,
		Concurrency: 1,
	}
	conf.Load(&config)

	configureLogging(config.Debug)

	invoker := tester.NewCLIInvoker(config.Path)

	t := &tester.T{
		SegmentWriteKey:     config.SegmentWriteKey,
		WebhookBucket:       config.WebhookBucket,
		WebhookAuthUsername: config.WebhookAuthUsername,
		Output:              os.Stdout,
		SkipRegex:           config.Skip,
		Timeout:             config.Timeout,
		Concurrency:         config.Concurrency,
	}

	err := t.Test(invoker)
	if err != nil {
		events.Log("test error: %{error}v", err)
		os.Exit(1)
	}
}

// configureLogging enables debug logging based on the argument provided.
// It also configures logs to be printed on os.Stderr as we want to keep the
// standard output clean.
func configureLogging(debug bool) {
	prefix := fmt.Sprintf("library-e2e-tester@%s[%d]: ", Version, os.Getpid())
	events.DefaultHandler = text.NewHandler(prefix, os.Stderr)
	if debug {
		events.DefaultLogger.EnableDebug = true
		events.DefaultLogger.EnableSource = true
	} else {
		events.DefaultLogger.EnableDebug = false
		events.DefaultLogger.EnableSource = false
	}
}
