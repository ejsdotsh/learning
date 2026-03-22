// SPDX-FileCopyrightText: 2021-present j e.j. sahala <jejs@sahala.org>
//
// SPDX-License-Identifier: MPL-2.0

package main

import "errors"

var (
	ErrNotNumber        = errors.New("data is not numeric")
	ErrInvalidColumn    = errors.New("invalid column number")
	ErrNoFiles          = errors.New("no input files")
	ErrInvalidOperation = errors.New("invalid operation")
)
