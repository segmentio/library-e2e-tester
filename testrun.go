package tester

import (
	"fmt"
	"io"
	"time"
)

/* Data structure to record information about a test run (execution of a test case) */
type TestRun struct {
	TestCaseName string
	StartTime    time.Time
	EndTime      time.Time
	Result       string
	Details      string
}

func (tr *TestRun) Start(testcase string) {
	tr.TestCaseName = testcase
	tr.StartTime = time.Now()
}

func (tr *TestRun) End(result string) {
	tr.EndTime = time.Now()
	tr.Result = result
}

func (tr *TestRun) AddDetails(details string) {
	tr.Details += "\n        " + details + "\n"
}

func (tr *TestRun) Error(errorDetails string) {
	tr.End("ERROR")
	tr.AddDetails(errorDetails)
}

func (tr *TestRun) Fail(failureDetails string) {
	tr.End("FAIL")
	tr.AddDetails(failureDetails)
}

/* Print out information about the test run in "go test" output format */
func (tr *TestRun) Print(writer io.Writer) error {
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
