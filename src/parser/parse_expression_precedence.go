package parser

import "dynamite/src/tokens"

// precedence constant
const (
	_ int = iota
	LOWEST
	EQUALS      // ==
	LESSGREATER // > or <
	SUM         // + or -
	PRODUCT     // * or /
	EXPONENT    // **
	PREFIX      // -X, +X, --X, ++X, !X
	POSTFIX     // X++, X--
	CALL        // myFunction(X)
)

var precedenceMap = map[tokens.TokenType]int{
	tokens.EQ:       EQUALS,
	tokens.NOT_EQ:   EQUALS,
	tokens.LT:       LESSGREATER,
	tokens.GT:       LESSGREATER,
	tokens.PLUS:     SUM,
	tokens.MINUS:    SUM,
	tokens.SLASH:    PRODUCT,
	tokens.ASTERISK: PRODUCT,
}

func (p *Parser) peekPrecedence() int {
	tokType := p.peekToken.Type
	pre, ok := precedenceMap[tokType]
	if ok {
		return pre
	}
	return LOWEST
}

func (p *Parser) currPrecedence() int {
	tokType := p.currToken.Type
	pre, ok := precedenceMap[tokType]
	if ok {
		return pre
	}
	return LOWEST
}
