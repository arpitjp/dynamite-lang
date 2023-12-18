package ast

import (
	"dynamite/src/tokens"
	"testing"
)

func TestString(t *testing.T) {
	program := &Program{
		Statements: []Statement{
			&LetStatement{
				Token: tokens.Token{Type: tokens.LET, Literal: "let"},
				Name: &IdentifierExpNode{
					Token: tokens.Token{Type: tokens.IDENT, Literal: "myVar"},
					Value: "myVar"},
				Value: &IdentifierExpNode{
					Token: tokens.Token{Type: tokens.IDENT, Literal: "anotherVar"},
					Value: "anotherVar",
				},
			},
		}}
	if program.String() != "let myVar = anotherVar;\n" {
		t.Errorf("program.String() wrong. got=%q", program.String())
	}
}
