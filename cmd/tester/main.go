package main

import (
	"os"

	"github.com/segmentio/conf"
	"github.com/segmentio/events"
	tester "github.com/segmentio/library-e2e-tester"

	_ "github.com/segmentio/events/text"
)

func main() {
	var config struct {
		Path                string `conf:"path"                     help:"path to the library binary" validate:"nonzero"`
		SegmentWriteKey     string `conf:"segment-write-key"        help:"writekey for the Segment project to send data to" validate:"nonzero"`
		WebhookBucket       string `conf:"webhook-bucket"           help:"webhook bucket the Segment project sends data to" validate:"nonzero"`
		WebhookAuthUsername string `conf:"webhook-auth-username"    help:"basic auth username for the webhook bucket the Segment project sends data to" validate:"nonzero"`
		Debug               bool   `conf:"debug"                    help:"Enable Debugging"`
	}
	conf.Load(&config)

	invoker := tester.NewCLIInvoker(config.Path)

	t := &tester.T{
		SegmentWriteKey:     config.SegmentWriteKey,
		WebhookBucket:       config.WebhookBucket,
		WebhookAuthUsername: config.WebhookAuthUsername,
	}

	if config.Debug {
		events.DefaultLogger.EnableDebug = true
		events.DefaultLogger.EnableSource = true
	} else {
		events.DefaultLogger.EnableDebug = false
		events.DefaultLogger.EnableSource = false
	}

	err := t.Test(invoker)
	if err != nil {
		events.Log("test error: %{error}v", err)
		os.Exit(1)
	}
}
