// SPDX-FileCopyrightText: 2021-present j e.j. sahala <jejs@sahala.org>
//
// SPDX-License-Identifier: MPL-2.0

package main

import (
	"errors"
	"fmt"
)

var (
	ErrValidation = errors.New("validation failed")
)

type stepErr struct {
	step  string
	msg   string
	cause error
}

func (s *stepErr) Error() string {
	return fmt.Sprintf("step: %q: %s: cause: %v", s.step, s.msg, s.cause)
}

func (s *stepErr) Is(target error) bool {
	t, ok := target.(*stepErr)
	if !ok {
		return false
	}

	return t.step == s.step
}

func (s *stepErr) Unwrap() error {
	return s.cause
}
