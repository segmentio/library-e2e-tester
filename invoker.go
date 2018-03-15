package tester

import (
	"context"
	"os"
	"os/exec"
)

// Invoker wraps the basic invoke method.
type Invoker func(ctx context.Context, args ...string) error

// NewCLIInvoker returns an invoker that calls the binary at the given path.
func NewCLIInvoker(path string) Invoker {
	return Invoker(func(ctx context.Context, args ...string) error {
		cmd := exec.CommandContext(ctx, path, args...)
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		return cmd.Run()
	})
}
