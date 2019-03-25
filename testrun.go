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
	Output       io.Writer
}

// NewTestRun returns an initialized TestRun struct.
func NewTestRun(testcase string, output io.Writer) TestRun {
	return TestRun{
		TestCaseName: testcase,
		Output:       output,
	}
}

// Start begins the test.
func (tr *TestRun) Start() {
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
	for _, detail := range details {
		tr.addDetail(detail)
	}
	tr.end(TestError)
}

// Fail ends the test with a failure and the given message.
func (tr *TestRun) Fail(details ...string) {
	for _, detail := range details {
		tr.addDetail(detail)
	}
	tr.end(TestFail)
}

// end ends the test with the given result and prints a report.
func (tr *TestRun) end(result string) {
	tr.EndTime = time.Now()
	tr.Result = result
	tr.report()
}

// report writes information about the test run in "go test" output format.
func (tr *TestRun) report() error {
	if tr.Result == TestSkip {
		return tr.printf("=== %v %v\n", TestSkip, tr.TestCaseName)
	}
	elapsed := tr.EndTime.Sub(tr.StartTime)
	if err := tr.printf("=== RUN %v\n", tr.TestCaseName); err != nil {
		return err
	}
	if err := tr.printf("--- %v: %v (%f seconds)\n", tr.Result, tr.TestCaseName, elapsed.Seconds()); err != nil {
		return err
	}
	if tr.Details != "" {
		tr.printf("			%v\n", tr.Details)
	}
	return nil
}

// printf writes information about the test run in "go test" output format.
func (tr *TestRun) printf(format string, a ...interface{}) error {
	_, err := fmt.Fprintf(tr.Output, format, a...)
	return err
}

// addDetail adds a detail about the test run.
func (tr *TestRun) addDetail(detail string) {
	tr.Details += "\n        " + detail + "\n"
}
