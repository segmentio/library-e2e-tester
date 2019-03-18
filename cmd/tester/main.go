package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"time"

	"github.com/segmentio/conf"
	"github.com/segmentio/events"
	tester "github.com/segmentio/library-e2e-tester"

	_ "github.com/segmentio/events/text"
)

/* print content of test result file on stdout */
func printTestResults(resultFile string) {
	fmt.Println("Test result file: ", resultFile)
	testResults, err := ioutil.ReadFile(resultFile) // just pass the file name
	if err != nil {
		events.Log("error print results: %{error}v", err)
	}
	fmt.Println(string(testResults))
}

func main() {
	var config struct {
		Path                string `conf:"path"                     help:"path to the library binary" validate:"nonzero"`
		SegmentWriteKey     string `conf:"segment-write-key"        help:"writekey for the Segment project to send data to" validate:"nonzero"`
		WebhookBucket       string `conf:"webhook-bucket"           help:"webhook bucket the Segment project sends data to" validate:"nonzero"`
		WebhookAuthUsername string `conf:"webhook-auth-username"    help:"basic auth username for the webhook bucket the Segment project sends data to" validate:"nonzero"`
		FailFast            bool   `conf:"failfast"                 help:"disable running additional tests after any test fails"`
		TestResultFile      string `conf:"test-result-file"         help:"file name to write test results"`
		SkipMessages        string `conf:"skip-messages"            help:"comma-separated list of message types to skip"`
		Debug               bool   `conf:"debug"                    help:"Enable Debugging"`
	}
	conf.Load(&config)

	// if TestResultFile is not specified, default to
	// test-results-YYYY-MM-DDTHH:MM:SS.txt (with current time value)
	if config.TestResultFile == "" {
		config.TestResultFile = fmt.Sprintf("test-results-%v.txt", time.Now().Format("2006-01-02T15:04:05"))
	}

	invoker := tester.NewCLIInvoker(config.Path)

	t := &tester.T{
		SegmentWriteKey:     config.SegmentWriteKey,
		WebhookBucket:       config.WebhookBucket,
		WebhookAuthUsername: config.WebhookAuthUsername,
		ReportFileName:      config.TestResultFile,
		FailFast:            config.FailFast,
		SkipMessages:        strings.Split(config.SkipMessages, ","),
	}

	if config.Debug {
		events.DefaultLogger.EnableDebug = true
		events.DefaultLogger.EnableSource = true
	} else {
		events.DefaultLogger.EnableDebug = false
		events.DefaultLogger.EnableSource = false
	}

	err := t.Test(invoker)
	printTestResults(config.TestResultFile)
	if err != nil {
		events.Log("test error: %{error}v", err)
		os.Exit(1)
	}
}
