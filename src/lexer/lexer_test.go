package lexer

import (
	"dynamite/src/tokens"
	"testing"
)

func TestNextToken(t *testing.T) {
	input := `let five = 5;
let ten = 10;

let add = func(x, y) {
  x + y;
};

let result = add(five, ten);
!-/*5;
5 < 10 > 5;

if (5 < 10) {
	return true;
} else {
	return false;
}

10 == 10;
10 != 9;
5 ** 5;
`

	expectedResult := []struct {
		expectedType    tokens.TokenType
		expectedLiteral string
		ln              int
	}{
		{tokens.LET, "let", 1},
		{tokens.IDENT, "five", 1},
		{tokens.ASSIGN, "=", 1},
		{tokens.INT, "5", 1},
		{tokens.SEMICOLON, ";", 1},
		{tokens.LET, "let", 2},
		{tokens.IDENT, "ten", 2},
		{tokens.ASSIGN, "=", 2},
		{tokens.INT, "10", 2},
		{tokens.SEMICOLON, ";", 2},
		{tokens.LET, "let", 4},
		{tokens.IDENT, "add", 4},
		{tokens.ASSIGN, "=", 4},
		{tokens.FUNCTION, "func", 4},
		{tokens.LPAREN, "(", 4},
		{tokens.IDENT, "x", 4},
		{tokens.COMMA, ",", 4},
		{tokens.IDENT, "y", 4},
		{tokens.RPAREN, ")", 4},
		{tokens.LBRACE, "{", 4},
		{tokens.IDENT, "x", 5},
		{tokens.PLUS, "+", 5},
		{tokens.IDENT, "y", 5},
		{tokens.SEMICOLON, ";", 5},
		{tokens.RBRACE, "}", 6},
		{tokens.SEMICOLON, ";", 6},
		{tokens.LET, "let", 8},
		{tokens.IDENT, "result", 8},
		{tokens.ASSIGN, "=", 8},
		{tokens.IDENT, "add", 8},
		{tokens.LPAREN, "(", 8},
		{tokens.IDENT, "five", 8},
		{tokens.COMMA, ",", 8},
		{tokens.IDENT, "ten", 8},
		{tokens.RPAREN, ")", 8},
		{tokens.SEMICOLON, ";", 8},
		{tokens.BANG, "!", 9},
		{tokens.MINUS, "-", 9},
		{tokens.SLASH, "/", 9},
		{tokens.ASTERISK, "*", 9},
		{tokens.INT, "5", 9},
		{tokens.SEMICOLON, ";", 9},
		{tokens.INT, "5", 10},
		{tokens.LT, "<", 10},
		{tokens.INT, "10", 10},
		{tokens.GT, ">", 10},
		{tokens.INT, "5", 10},
		{tokens.SEMICOLON, ";", 10},
		{tokens.IF, "if", 12},
		{tokens.LPAREN, "(", 12},
		{tokens.INT, "5", 12},
		{tokens.LT, "<", 12},
		{tokens.INT, "10", 12},
		{tokens.RPAREN, ")", 12},
		{tokens.LBRACE, "{", 12},
		{tokens.RETURN, "return", 13},
		{tokens.TRUE, "true", 13},
		{tokens.SEMICOLON, ";", 13},
		{tokens.RBRACE, "}", 14},
		{tokens.ELSE, "else", 14},
		{tokens.LBRACE, "{", 14},
		{tokens.RETURN, "return", 15},
		{tokens.FALSE, "false", 15},
		{tokens.SEMICOLON, ";", 15},
		{tokens.RBRACE, "}", 16},
		{tokens.INT, "10", 18},
		{tokens.EQ, "==", 18},
		{tokens.INT, "10", 18},
		{tokens.SEMICOLON, ";", 18},
		{tokens.INT, "10", 19},
		{tokens.NOT_EQ, "!=", 19},
		{tokens.INT, "9", 19},
		{tokens.SEMICOLON, ";", 19},
		{tokens.INT, "5", 20},
		{tokens.EXPONENT, "**", 20},
		{tokens.INT, "5", 20},
		{tokens.SEMICOLON, ";", 20},
		{tokens.EOF, "", 21},
	}

	l := New(input);

	for i, val := range expectedResult {
		tok := l.NextToken()
		
		// check type
		if val.expectedType != tok.Type {
			t.Errorf("Tests[%d] - wrong tokenType. Expected: %q, got: %q | Ln: %d : col : %d", i, val.expectedType, tok.Type, tok.Ln, tok.Col)
		}
		// check literal
		if val.expectedLiteral != tok.Literal {
			t.Errorf("Tests[%d] - wrong tokenLiteral. Expected: %q, got: %q | Ln: %d : col : %d", i, val.expectedLiteral, tok.Literal, tok.Ln, tok.Col)
		}
		// check line number
		if val.ln != tok.Ln {
			t.Errorf("Tests[%d] - wrong line number. Expected: %d, got: %d", i, val.ln, tok.Ln)
		}
	}
}
