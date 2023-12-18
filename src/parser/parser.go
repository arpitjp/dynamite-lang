package parser

import (
	"dynamite/src/ast"
	"dynamite/src/lexer"
	"dynamite/src/logger"
	"dynamite/src/tokens"
	"fmt"
)

type Parser struct {
	l *lexer.Lexer
	currToken tokens.Token
	peekToken tokens.Token
	errors []string
	// pratt parser
	// Rule: they start with their token and end with their token
	prefixParseFns map[tokens.TokenType]prefixParseFn
	infixParseFnx map[tokens.TokenType]infixParseFn
}

func New(l *lexer.Lexer) *Parser {
	p :=  &Parser{
		l: l,
		prefixParseFns: make(map[tokens.TokenType]prefixParseFn),
		infixParseFnx: make(map[tokens.TokenType]infixParseFn),
	}

	// registering functions
	p.registerPrefixParseFn(tokens.IDENT, p.parseIdentifierExpression)

	// initializing both currToken and nextToken
	p.NextToken()
	p.NextToken()

	return p
}

// Creates root program node
func (p *Parser)ParseProgram() *ast.Program {
	programNode := &ast.Program{}

	for p.currToken.Type != tokens.EOF {
		stmt := p.parseStatement()
		if stmt != nil {
			programNode.Statements = append(programNode.Statements, stmt)
		}
		p.NextToken()
	}

	return programNode
}

func(p *Parser) NextToken() {
	p.currToken = p.peekToken
	p.peekToken = p.l.NextToken()
}

// if next token is as expected, move tokens. Else add error to program
func (p *Parser) expectPeekToken(tok tokens.TokenType) bool {
	if tok == p.peekToken.Type {
		p.NextToken()
		return true
	}
	p.expectPeekError(tok)
	return false
}

func (p *Parser) parsingError(s string, t tokens.Token) {
	errorLine := p.l.GetErrorLine(t.Ln, t.Col)
	msg := fmt.Sprintf("\n%s %s at%s", logger.Error("Parsing error:"), s, errorLine)
	p.errors = append(p.errors, msg)
}

func (p *Parser) expectPeekError(expTokenType tokens.TokenType) {
	p.parsingError(fmt.Sprintf("expected %q, but got %q", expTokenType, p.peekToken.Type), p.peekToken)
}
