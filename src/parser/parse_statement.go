package parser

import (
	"dynamite/src/ast"
	"dynamite/src/tokens"
)

func (p *Parser) parseStatement() ast.Statement {
	var stmt ast.Statement
	switch p.currToken.Type {
	case tokens.LET:
		stmt = p.parseLetStatement()
	}
	return stmt
}

func (p *Parser) parseLetStatement() *ast.LetStatement {
	node := &ast.LetStatement{Token: p.currToken}

	// identifier token
	if !p.expectPeekToken(tokens.IDENT) {
		return nil // exit early if you encounter a parsing error
	}
	node.Name = &ast.IdentifierExpNode{
		Token: p.currToken,
		Value: p.currToken.Literal,
	}

	// equals token
	if !p.expectPeekToken(tokens.ASSIGN) {
		return nil
	}

	// expression token
	// TODO: We're skipping the expressions until we // encounter a semicolon
	for p.currToken.Type != tokens.SEMICOLON {
		p.NextToken()
	}
	
	return node
}
