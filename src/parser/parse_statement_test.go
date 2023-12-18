package parser

import (
	"dynamite/src/ast"
	"dynamite/src/lexer"
	"strings"
	"testing"
)

func TestLetStatements(t *testing.T) {
	input := `
	let x = 5;
	let y = 10;
	let foobar = 838383;`

	// let identifier = expression
	expected := []struct{
		identifier string
	}{
		{
			identifier: "x",
		},
		{
			identifier: "y",
		},
		{
			identifier: "foobar",
		},
	}

	l := lexer.New(input)
	p := New(l)
	programNode := p.ParseProgram()

	if programNode == nil {
		t.Fatalf("ParseProgram returned nil")
	}

	if l := len(programNode.Statements); l != 3 {
		t.Fatalf("Expected 3 statements in program, but got %d", l)
	}

	if l := len(p.errors); l != 0 {
		t.Fatalf("Got %d parsing errors:\n%s", l, strings.Join(p.errors, "\n"))
	}

	// test actual let statement
	for i, exp := range expected {
		stmt := programNode.Statements[i]
		testLetStatement(i, stmt, exp.identifier, t)
	}
}

func testLetStatement(i int, stmt ast.Statement, expIdentifier string, t *testing.T) {
	if stmt == nil {
		t.Fatalf("Test[%d] failure: let statement found to be nil", i)
	}

	letStmt, ok := stmt.(*ast.LetStatement)
	if !ok {
		t.Fatalf("Test[%d] failure: expected type *ast.LetStatement but found %T", i, stmt)
	}

	if val := letStmt.Name.Value; val != expIdentifier {
		t.Fatalf("Test[%d] failure: expected identifier value %q, but got %q", i, expIdentifier, val)
	}
	
	if val := letStmt.Name.TokenLiteral(); val != expIdentifier {
		t.Fatalf("Test[%d] failure: expected token literal value %q, but got %q", i, expIdentifier, val)
	}
}