package main

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/segmentio/conf"
	"github.com/segmentio/events"
	"github.com/segmentio/events/text"
	tester "github.com/segmentio/library-e2e-tester"
)

// Config represents the options that can be supplied to the harness.
type Config struct {
	Path                string        `conf:"path"                     help:"path to the library binary" validate:"nonzero"`
	SegmentWriteKey     string        `conf:"segment-write-key"        help:"writekey for the Segment project to send data to" validate:"nonzero"`
	WebhookBucket       string        `conf:"webhook-bucket"           help:"webhook bucket the Segment project sends data to" validate:"nonzero"`
	WebhookAuthUsername string        `conf:"webhook-auth-username"    help:"basic auth username for the webhook bucket the Segment project sends data to" validate:"nonzero"`
	SkipFixtures        string        `conf:"skip-fixtures"            help:"comma-separated list of fixtures to skip"`
	Timeout             time.Duration `conf:"timeout"                  help:"Timeout before giving up checking on a message"`
	Debug               bool          `conf:"debug"                    help:"Enable Debugging"`
}

func main() {
	config := Config{
		Timeout: 1 * time.Minute,
	}
	conf.Load(&config)

	configureLogging(config.Debug)

	invoker := tester.NewCLIInvoker(config.Path)

	t := &tester.T{
		SegmentWriteKey:     config.SegmentWriteKey,
		WebhookBucket:       config.WebhookBucket,
		WebhookAuthUsername: config.WebhookAuthUsername,
		Output:              os.Stdout,
		SkipFixtures:        splitStringList(config.SkipFixtures),
		Timeout:             config.Timeout,
	}

	err := t.Test(invoker)
	if err != nil {
		events.Log("test error: %{error}v", err)
		os.Exit(1)
	}
}

// splitStringList slices s into all substrings separated by a comma and returns a slice of the substrings.
// It differs from strings.Split by returning an empty array if s is an empty string.
func splitStringList(s string) []string {
	if strings.TrimSpace(s) == "" {
		return []string{}
	}
	return strings.Split(s, ",")
}

// configureLogging enables debug logging based on the argument provided.
// It also configures logs to be printed on os.Stderr as we want to keep the
// standard output clean.
func configureLogging(debug bool) {
	prefix := fmt.Sprintf("library-e2e-tester[%d]: ", os.Getpid())
	events.DefaultHandler = text.NewHandler(prefix, os.Stderr)
	if debug {
		events.DefaultLogger.EnableDebug = true
		events.DefaultLogger.EnableSource = true
	} else {
		events.DefaultLogger.EnableDebug = false
		events.DefaultLogger.EnableSource = false
	}
}
