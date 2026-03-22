// SPDX-FileCopyrightText: 2021-present j e.j. sahala <jejs@sahala.org>
//
// SPDX-License-Identifier: MPL-2.0

// ast/ast_test.go

package ast

import (
	"testing"

	"github.com/ejsdotsh/learning/monkey/token"
)

func TestString(t *testing.T) {
	program := &Program{
		Statements: []Statement{
			&LetStatement{
				Token: token.Token{Type: token.LET, Literal: "let"},
				Name: &Identifier{
					Token: token.Token{Type: token.IDENT, Literal: "myVar"},
					Value: "myVar",
				},
				Value: &Identifier{
					Token: token.Token{Type: token.IDENT, Literal: "anotherVar"},
					Value: "anotherVar",
				},
			},
		},
	}

	if program.String() != "let myVar = anotherVar;" {
		t.Errorf("program.String() wrong. got: %q", program.String())
	}
}
