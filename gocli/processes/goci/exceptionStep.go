package main

import (
	"bytes"
	"fmt"
	"os/exec"
)

// exceptionStep extends the step type by embedding it
type exceptionStep struct {
	step
}

// newExceptionStep reuses the constructor from the embedded `step` type
func newExceptionStep(name, exe, message, proj string, args []string) exceptionStep {
	s := exceptionStep{}

	s.step = newStep(name, exe, message, proj, args)

	return s
}

// execute implements the `executer` interface
func (s exceptionStep) execute() (string, error) {
	cmd := exec.Command(s.exe, s.args...)

	var out bytes.Buffer
	cmd.Stdout = &out

	cmd.Dir = s.proj

	if err := cmd.Run(); err != nil {
		return "", &stepErr{
			step:  s.name,
			msg:   "failed to execute",
			cause: err,
		}
	}

	if out.Len() > 0 {
		return "", &stepErr{
			step:  s.name,
			msg:   fmt.Sprintf("invalid format: %s", out.String()),
			cause: nil,
		}
	}

	return s.message, nil
}
