package main

import (
	"io"
	"os"
	"strings"
	"time"

	"github.com/segmentio/conf"
	"github.com/segmentio/events"
	tester "github.com/segmentio/library-e2e-tester"

	_ "github.com/segmentio/events/text"
)

// Config represents the options that can be supplied to the harness.
type Config struct {
	Path                string        `conf:"path"                     help:"path to the library binary" validate:"nonzero"`
	SegmentWriteKey     string        `conf:"segment-write-key"        help:"writekey for the Segment project to send data to" validate:"nonzero"`
	WebhookBucket       string        `conf:"webhook-bucket"           help:"webhook bucket the Segment project sends data to" validate:"nonzero"`
	WebhookAuthUsername string        `conf:"webhook-auth-username"    help:"basic auth username for the webhook bucket the Segment project sends data to" validate:"nonzero"`
	FailFast            bool          `conf:"failfast"                 help:"disable running additional tests after any test fails"`
	TestResultFile      string        `conf:"test-result-file"         help:"file name to write test results"`
	SkipFixtures        string        `conf:"skip-fixtures"            help:"comma-separated list of fixtures to skip"`
	Timeout             time.Duration `conf:"timeout"                    help:"Timeout before giving up checking on a message"`
	Debug               bool          `conf:"debug"                    help:"Enable Debugging"`
}

func main() {
	config := Config{
		Timeout: 1 * time.Minute,
	}
	conf.Load(&config)

	if config.Debug {
		events.DefaultLogger.EnableDebug = true
		events.DefaultLogger.EnableSource = true
	} else {
		events.DefaultLogger.EnableDebug = false
		events.DefaultLogger.EnableSource = false
	}

	invoker := tester.NewCLIInvoker(config.Path)

	t := &tester.T{
		SegmentWriteKey:     config.SegmentWriteKey,
		WebhookBucket:       config.WebhookBucket,
		WebhookAuthUsername: config.WebhookAuthUsername,
		Output:              makeOutputWriter(config.TestResultFile),
		FailFast:            config.FailFast,
		SkipFixtures:        strings.Split(config.SkipFixtures, ","),
		Timeout:             config.Timeout,
	}

	err := t.Test(invoker)
	if err != nil {
		events.Log("test error: %{error}v", err)
		os.Exit(1)
	}
}

// makeOutputWriter returns an io.Writer for the tester to write results to.
func makeOutputWriter(testResultFile string) io.Writer {
	if testResultFile == "" {
		return os.Stdout
	}

	f, err := os.Create(testResultFile)
	if err != nil {
		events.Log("error creating file %{file}v: %{error}v", testResultFile, err)
		os.Exit(1)
	}

	return io.MultiWriter(os.Stdout, f)
}
