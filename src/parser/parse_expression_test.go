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

func TestIntegerExpression(t *testing.T) {
	input := `5;`
	expected := []int64{
		5,
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
		testIntegerLiteral(t, expStmt.Expression, expected[i], i)
	}
}

func testIntegerLiteral(t *testing.T, exp ast.Expression, no int64, i int) {
	integerExp, ok := exp.(*ast.IntegerExpressionNode)

	if !ok {
		t.Fatalf("Test[%d] failure: Expected expression to be of type *ast.IntegerExpressionNode, but got %T", i, exp)
	}

	if got := integerExp.Value; got != no {
		t.Fatalf("Test[%d] failure: Expected exp value to be %q, but got %q", i, no, got)
	}
}

func testIdentifier(t *testing.T, exp ast.Expression, value string) bool { ident, ok := exp.(*ast.Identifier)
	if !ok {
		t.Errorf("exp not *ast.Identifier. got=%T", exp)
		return false 
	}
	if ident.Value != value {
		t.Errorf("ident.Value not %s. got=%s", value, ident.Value) 
		return false
	}
	if ident.TokenLiteral() != value { 
		t.Errorf("ident.TokenLiteral not %s. got=%s", value, ident.TokenLiteral())
		return false 
	}
	return true 
}

func TestParsingPrefixExpressions(t *testing.T) {
	prefixTests := []struct {
		input        string
		operator     string
		integerValue int64
	}{
		{"!5;", "!", 5},
		{"-15;", "-", 15},
	}

	for _, test := range prefixTests {
		l := lexer.New(test.input)
		p := New(l)
		programNode := p.ParseProgram()

		testProgramNode(t, 1, programNode, p)

		for i, stmt := range programNode.Statements {
			expStmt, ok := stmt.(*ast.ExpressionStatement)
			if !ok {
				t.Fatalf("Test[%d] failure: Expected stmt to be of type *ast.ExpressionStatement, but got %T", i, stmt)
			}
			prefixExp, ok := expStmt.Expression.(*ast.PrefixExpressionNode)
			if !ok {
				t.Fatalf("Test[%d] failure: Expected expression to be of type *ast.PrefixExpressionNode, but got %T", i, expStmt.Expression)
			}
			if got := prefixExp.Operator; got != test.operator {
				t.Fatalf("Test[%d] failure: Expected operator to be %q, but got %q", i, test.operator, got)
			}
			testIntegerLiteral(t, prefixExp.Right, test.integerValue, i)
		}
	}
}

func TestParsingInfixExpressions(t *testing.T) {
	infixTests := []struct {
		input      string
		leftValue  int64
		operator   string
		rightValue int64
	}{
		{"5 + 5;", 5, "+", 5},
		{"5 - 5;", 5, "-", 5},
		{"5 * 5;", 5, "*", 5},
		{"5 / 5;", 5, "/", 5},
		{"5 > 5;", 5, ">", 5},
		{"5 < 5;", 5, "<", 5},
		{"5 == 5;", 5, "==", 5},
		{"5 != 5;", 5, "!=", 5},
	}

	for _, test := range infixTests {
		l := lexer.New(test.input)
		p := New(l)
		programNode := p.ParseProgram()

		testProgramNode(t, 1, programNode, p)

		for i, stmt := range programNode.Statements {
			expStmt, ok := stmt.(*ast.ExpressionStatement)
			if !ok {
				t.Fatalf("Test[%d] failure: Expected stmt to be of type *ast.ExpressionStatement, but got %T", i, stmt)
			}
			infixExp, ok := expStmt.Expression.(*ast.InfixExpressionNode)
			if !ok {
				t.Fatalf("Test[%d] failure: Expected expression to be of type *ast.PrefixExpressionNode, but got %T", i, expStmt.Expression)
			}
			if got := infixExp.Operator; got != test.operator {
				t.Fatalf("Test[%d] failure: Expected operator to be %q, but got %q", i, test.operator, got)
			}
			testIntegerLiteral(t, infixExp.Left, test.leftValue, i)
			testIntegerLiteral(t, infixExp.Right, test.rightValue, i)
		}
	}
}
