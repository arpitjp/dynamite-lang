package parser

import (
	"dynamite/src/ast"
	"dynamite/src/tokens"
	"fmt"
	"strconv"
)

type prefixParseFn func() ast.Expression
type infixParseFn func(ast.Expression) ast.Expression

func (p *Parser) registerPrefixParseFn(tok tokens.TokenType, f prefixParseFn) {
	p.prefixParseFns[tok] = f
}

func (p *Parser) registerInfixParseFn(tok tokens.TokenType, f infixParseFn) {
	p.infixParseFnx[tok] = f
}

const (
	_ int = iota
	LOWEST
	EQUALS  // == LESSGREATER // > or <
	SUM     // +
	PRODUCT // *
	PREFIX  // -X or !X
	CALL    // myFunction(X)
)

func (p *Parser) parseExpression(precedence int) ast.Expression {
	prefixFn := p.prefixParseFns[p.currToken.Type]
	if prefixFn == nil {
		return nil
	}
	leftExp := prefixFn()

	return leftExp
}

// -------------
func (p *Parser) parseIdentifierExpression() ast.Expression {
	return &ast.IdentifierExpNode{
		Token: p.currToken,
		Value: p.currToken.Literal,
	}
}

func (p *Parser) parseIntegerExpressionNode() ast.Expression {
	node := &ast.IntegerExpressionNode{
		Token: p.currToken,
	}
	no, err := strconv.ParseInt(p.currToken.Literal, 0, 64)
	if err != nil {
		p.parsingError(fmt.Sprintf("could not parse %q as integer\n%s", p.currToken.Literal, err.Error()), p.currToken)
		return nil
	}

	node.Value = no
	return node
}
