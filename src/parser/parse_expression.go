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

func (p *Parser) parseExpression(precedence int) ast.Expression {
	prefixFn := p.prefixParseFns[p.currToken.Type]
	if prefixFn == nil {
		p.parsingError(fmt.Sprintf("prefix function not found for %q", p.currToken.Type), p.currToken)
		return nil
	}
	leftExp := prefixFn()

	for p.peekToken.Type != tokens.SEMICOLON && precedence < p.peekPrecedence() {
		infixFn := p.infixParseFnx[p.peekToken.Type]
		if infixFn == nil {
			return leftExp
		}
		p.NextToken()
		leftExp = infixFn(leftExp)
	}

	return leftExp
}

// -------------
func (p *Parser) parseIdentifierExpression() ast.Expression {
	return &ast.IdentifierExpNode{
		Token: p.currToken,
		Value: p.currToken.Literal,
	}
}

func (p *Parser) parseIntegerExpression() ast.Expression {
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

func (p *Parser) parsePrefixExpression() ast.Expression {
	node := &ast.PrefixExpressionNode{
		Token: p.currToken,
	}
	node.Operator = node.TokenLiteral()
	p.NextToken()
	node.Right = p.parseExpression(PREFIX)
	return node
}

func (p *Parser) parseInfixExpression(left ast.Expression) ast.Expression {
	node := &ast.InfixExpressionNode{
		Token: p.currToken,
		Operator: p.currToken.Literal,
		Left: left,
	}
	precedence := p.currPrecedence()
	p.NextToken()
	node.Right = p.parseExpression(precedence)
	return node
}