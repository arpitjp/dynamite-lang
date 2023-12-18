package parser

import (
	"dynamite/src/ast"
	"dynamite/src/lexer"
	"testing"
)

func TestIdentifierExpression(t *testing.T) {
	input := `foobar;`
	expected := []string{
		"foobar",
	}

	l := lexer.New(input)
	p := New(l)
	programNode := p.ParseProgram()

	testProgramNode(t, 1, programNode, p)

	for i, stmt := range programNode.Statements {
		expStmt, ok := stmt.(*ast.ExpressionStatement)
		if !ok {
			t.Fatalf("Test[%d] failure: Expected stmt to be of type *ast.ExpressionStatement, but got %T", i, stmt)
		}
		identifierExp, ok := expStmt.Expression.(*ast.IdentifierExpNode)
		if !ok {
			t.Fatalf("Test[%d] failure: Expected expression to be of type *ast.IdentifierExpNode, but got %T", i, expStmt.Expression)
		}
		if got := identifierExp.TokenLiteral(); got != expected[i] {
			t.Fatalf("Test[%d] failure: Expected token literal to be %q, but got %q", i, expected[i], got)
		}
		if got := identifierExp.Value; got != expected[i] {
			t.Fatalf("Test[%d] failure: Expected exp value to be %q, but got %q", i, expected[i], got)
		}
	}
}
