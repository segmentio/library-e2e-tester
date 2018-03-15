package tester

import (
	"bufio"
	"context"
	"io"
	"text/template"
	"time"

	"github.com/segmentio/ksuid"
)

// Producer is an interface implemented by types that produces messages.
type Producer interface {
	// Produce is called to generate messages.
	//
	// Produce may be called concurrently from multiple goroutines.
	Produce(ctx context.Context, input io.Reader, output io.Writer) (err error)
}

// NewTemplatedProducer takes a templated message as a model to the `.Produce` method
// and returns a generated message using the model's template.
//
// The following functions can be used in the template:
//
//  - `{{id}}`        -- generates a new string-based KSUID
//  - `{{timestamp}}` -- generates the current ISO 8601 formated date.
func NewTemplatedProducer() Producer {
	tmpl := template.New("template-producer")
	tmpl.Funcs(map[string]interface{}{
		"id":        func() string { return ksuid.New().String() },
		"timestamp": func() string { return time.Now().UTC().Format(time.RFC3339) },
	})

	return &templatedProducer{
		tmpl: tmpl,
	}
}

type templatedProducer struct {
	tmpl *template.Template
}

func (producer *templatedProducer) Produce(ctx context.Context, in io.Reader, out io.Writer) error {
	scanner := bufio.NewScanner(in)
	for scanner.Scan() {
		tmpl, err := producer.tmpl.Parse(scanner.Text())
		if err != nil {
			return err
		}
		if err := tmpl.Execute(out, nil); err != nil {
			return err
		}
	}

	if err := scanner.Err(); err != nil {
		return err
	}

	return nil
}
