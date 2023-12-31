package parser

import (
	"dynamite/src/ast"
	"dynamite/src/lexer"
	"dynamite/src/logger"
	"dynamite/src/tokens"
	"fmt"
)

type Parser struct {
	l         *lexer.Lexer
	currToken tokens.Token
	peekToken tokens.Token
	errors    []string
	// pratt parser
	// Rule: they start with their token and end with their token
	prefixParseFns map[tokens.TokenType]prefixParseFn
	infixParseFnx  map[tokens.TokenType]infixParseFn
}

func New(l *lexer.Lexer) *Parser {
	p := &Parser{
		l:              l,
		prefixParseFns: make(map[tokens.TokenType]prefixParseFn),
		infixParseFnx:  make(map[tokens.TokenType]infixParseFn),
	}

	// registering functions
	// prefix
	p.registerPrefixParseFn(tokens.IDENT, p.parseIdentifierExpression)
	p.registerPrefixParseFn(tokens.INT, p.parseIntegerExpression)
	p.registerPrefixParseFn(tokens.BANG, p.parsePrefixExpression)
	p.registerPrefixParseFn(tokens.MINUS, p.parsePrefixExpression)
	// infix
	p.registerInfixParseFn(tokens.PLUS, p.parseInfixExpression) 
	p.registerInfixParseFn(tokens.MINUS, p.parseInfixExpression) 
	p.registerInfixParseFn(tokens.SLASH, p.parseInfixExpression) 
	p.registerInfixParseFn(tokens.ASTERISK, p.parseInfixExpression) 
	p.registerInfixParseFn(tokens.EQ, p.parseInfixExpression) 
	p.registerInfixParseFn(tokens.NOT_EQ, p.parseInfixExpression) 
	p.registerInfixParseFn(tokens.LT, p.parseInfixExpression) 
	p.registerInfixParseFn(tokens.GT, p.parseInfixExpression)

	// initializing both currToken and nextToken
	p.NextToken()
	p.NextToken()

	return p
}

// Creates root program node
func (p *Parser) ParseProgram() *ast.Program {
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

func (p *Parser) NextToken() {
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
