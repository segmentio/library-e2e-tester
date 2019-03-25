package tester

import (
	"fmt"
	"io"
	"time"
)

const (
	// TestPass is the result when a test passes.
	TestPass = "PASS"
	// TestFail is the result when a test fails.
	TestFail = "FAIL"
	// TestSkip is the result when a test is skipped.
	TestSkip = "SKIP"
	// TestError is the result when a test errors.
	TestError = "ERROR"
)

// TestRun is the records information about a test run (execution of a test case).
type TestRun struct {
	TestCaseName string
	StartTime    time.Time
	EndTime      time.Time
	Result       string
	Details      string
}

// Start begins the test.
func (tr *TestRun) Start(testcase string) {
	tr.TestCaseName = testcase
	tr.StartTime = time.Now()
}

// Skip skips the test.
func (tr *TestRun) Skip() {
	tr.end(TestSkip)
}

// Pass passes the test.
func (tr *TestRun) Pass() {
	tr.end(TestPass)
}

// Error ends the test with an error and the given message.
func (tr *TestRun) Error(details ...string) {
	tr.end(TestError)
	for _, detail := range details {
		tr.addDetail(detail)
	}
}

// Fail ends the test with a failure and the given message.
func (tr *TestRun) Fail(details ...string) {
	tr.end(TestFail)
	for _, detail := range details {
		tr.addDetail(detail)
	}
}

// Print writes information about the test run in "go test" output format.
func (tr *TestRun) Print(writer io.Writer) error {
	if tr.Result == TestSkip {
		_, err := fmt.Fprintf(writer, "=== %v %v\n", TestSkip, tr.TestCaseName)
		return err
	}
	elapsed := tr.EndTime.Sub(tr.StartTime)
	_, err := fmt.Fprintf(writer, "=== RUN %v\n", tr.TestCaseName)
	if err != nil {
		return err
	}
	_, err = fmt.Fprintf(writer, "--- %v: %v (%f seconds)\n", tr.Result, tr.TestCaseName, elapsed.Seconds())
	if tr.Details != "" {
		_, err = fmt.Fprintf(writer, "			%v\n", tr.Details)
	}
	return err
}

// end ends the test with the given result.
func (tr *TestRun) end(result string) {
	tr.EndTime = time.Now()
	tr.Result = result
}

// addDetail adds a detail about the test run.
func (tr *TestRun) addDetail(detail string) {
	tr.Details += "\n        " + detail + "\n"
}
