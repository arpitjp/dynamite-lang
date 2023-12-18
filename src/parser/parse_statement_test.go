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
	expected := []struct {
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

	testProgramNode(t, 3, programNode, p)

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

func testProgramNode(t *testing.T, pLen int, programNode *ast.Program, p *Parser) {
	if programNode == nil {
		t.Fatalf("ParseProgram returned nil")
	}

	if l := len(programNode.Statements); l != pLen {
		t.Fatalf("Expected %d statements in program, but got %d", pLen, l)
	}

	if l := len(p.errors); l != 0 {
		t.Fatalf("Got %d parsing errors:\n%s", l, strings.Join(p.errors, "\n"))
	}
}

func TestReturnStatements(t *testing.T) {
	input := `
   return 5;
   return 10;
   return 993322;
   `

	l := lexer.New(input)
	p := New(l)
	programNode := p.ParseProgram()

	testProgramNode(t, 3, programNode, p)

	for i, stmt := range programNode.Statements {
		if stmt == nil {
			t.Fatalf("Test[%d] failure: return statement found to be nil", i)
		}

		returnStmt, ok := stmt.(*ast.ReturnStatement)
		if !ok {
			t.Fatalf("Test[%d] failure: expected type *ast.ReturnStatement but found %T", i, stmt)
		}

		if val := returnStmt.TokenLiteral(); val != "return" {
			t.Fatalf("Test[%d] failure; expected token literal to be return but found %T", i, val)
		}
	}
}
