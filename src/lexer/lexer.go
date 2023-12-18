package lexer

import (
	"dynamite/src/constants"
	"dynamite/src/env"
	"dynamite/src/logger"
	"dynamite/src/tokens"
	"fmt"
	"strings"

	"golang.org/x/exp/utf8string"
)

type Lexer struct {
	input utf8string.String
	position int
	nextPosition int
	ch rune
	str string
	// meta
	ln int // updated by eatWhitespace()
	col int // updated by readChar()
}

func New(input string) *Lexer {
	cleanInput := strings.Replace(input, "\t", env.Getenv(env.ROCKET_TAB_WIDTH), -1)
	l := &Lexer{
		input: *utf8string.NewString(cleanInput),
		ln: 1,
	}
	l.readChar()
	return l
}

func (l *Lexer) readChar() {
	if l.nextPosition < l.input.RuneCount() {
		l.ch = l.input.At(l.nextPosition)
	} else {
		l.ch = constants.NULL_CHAR
	}
	l.str = string(l.ch)
	l.position = l.nextPosition
	l.nextPosition++
	l.col++
}

func(l *Lexer) peekChar() rune {
	if l.nextPosition < l.input.RuneCount() {
		return l.input.At(l.nextPosition)
	}
	return constants.NULL_CHAR
}

func(l *Lexer) eatWhitespace() {
	for l.ch == '\n' || l.ch == '\t' || l.ch == ' ' || l.ch == '\r' {
		if l.ch == '\n' {
			l.ln++
			l.col = 0
		}
		l.readChar()
	}
}

func(l *Lexer) createToken(tokenType tokens.TokenType, literal string) tokens.Token {
	return tokens.Token{
		Type: tokenType,
		Literal: literal,
		Col: l.col,
		Ln: l.ln,
	}
}

func(l *Lexer) parseNumber() string {
	s := ""
	for isDigit(l.ch) {
		s += string(l.ch)
		if !isDigit(l.peekChar()) {
			return s
		}
		l.readChar()
	}
	return s
}

func(l *Lexer) parseLiteral() string {
	s := ""
	for isLetter(l.ch) {
		s += string(l.ch)
		if !isLetter(l.peekChar()) {
			return s
		}
		l.readChar()
	}
	return s
}

func (l *Lexer) NextToken() tokens.Token {
	var tok tokens.Token
	l.eatWhitespace()

	switch(l.ch) {
	// single char tokens
	case constants.NULL_CHAR:
		tok = l.createToken(tokens.EOF, "")
	case '/':
		tok = l.createToken(tokens.SLASH, string(l.ch))
	case '<':
		tok = l.createToken(tokens.LT, string(l.ch))
	case '>':
		tok = l.createToken(tokens.GT, string(l.ch))
	case ',':
		tok = l.createToken(tokens.COMMA, string(l.ch))
	case ';':
		tok = l.createToken(tokens.SEMICOLON, string(l.ch))
	case ':':
		tok = l.createToken(tokens.COLON, string(l.ch))
	case '(':
		tok = l.createToken(tokens.LPAREN, string(l.ch))
	case ')':
		tok = l.createToken(tokens.RPAREN, string(l.ch))
	case '{':
		tok = l.createToken(tokens.LBRACE, string(l.ch))
	case '}':
		tok = l.createToken(tokens.RBRACE, string(l.ch))
	case '[':
		tok = l.createToken(tokens.LBRACKET, string(l.ch))
	case ']':
		tok = l.createToken(tokens.RBRACKET, string(l.ch))

	// multi char tokens
	case '+':
		var tt tokens.TokenType = tokens.PLUS
		literal := string(l.ch)
		if l.peekChar() == '+' {
			l.readChar()
			tt = tokens.INCREMENT
			literal += string(l.ch)
		}
		tok = l.createToken(tt, literal)
	case '-':
		var tt tokens.TokenType = tokens.MINUS
		literal := string(l.ch)
		if l.peekChar() == '-' {
			l.readChar()
			tt = tokens.DECREMENT
			literal += string(l.ch)
		}
		tok = l.createToken(tt, literal)
	case '*':
		var tt tokens.TokenType = tokens.ASTERISK
		literal := string(l.ch)
		if l.peekChar() == '*' {
			l.readChar()
			tt = tokens.EXPONENT
			literal += string(l.ch)
		}
		tok = l.createToken(tt, literal)
	case '=':
		var tt tokens.TokenType = tokens.ASSIGN
		literal := string(l.ch)
		if l.peekChar() == '=' {
			l.readChar()
			tt = tokens.EQ
			literal += string(l.ch)
		}
		tok = l.createToken(tt, literal)
	case '!':
		var tt tokens.TokenType = tokens.BANG
		literal := string(l.ch)
		if l.peekChar() == '=' {
			l.readChar()
			tt = tokens.NOT_EQ
			literal += string(l.ch)
		}
		tok = l.createToken(tt, literal)

	// identifiers (a-zA-Z_), keywords, integers
	default:
		if isLetter(l.ch) {
			literal := l.parseLiteral()
			var tt tokens.TokenType
			if keyword, ok := tokens.KeywordsMap[literal]; ok {
				// keyword
				tt = keyword
			} else {
				// identifier
				tt = tokens.IDENT
			}
			tok = l.createToken(tt, literal)
		} else if isDigit(l.ch) {
			// integer
			numLiteral := l.parseNumber()
			tok = l.createToken(tokens.INT, numLiteral)
		} else {
			// illegal token
			tok = l.createToken(tokens.ILLEGAL, string(l.ch))
		}
	}
	l.readChar()
	return tok
}

// other utils -------

func isDigit(ch rune) bool {
	if ch >= '0' && ch <= '9' {
		return true
	}
	return false
}

func isLetter(ch rune) bool {
	if (ch >= 'a' && ch <= 'z') || (ch >= 'A' && ch <= 'Z') || ch == '_' {
		return true
	}
	return false
}

func (l *Lexer) GetErrorLine(ln int, col int) string {
	str := "\n"
	lineStart := 0
	for line := 0; line < ln-1 && lineStart < l.input.RuneCount()-1; lineStart++ {
		if l.input.At(lineStart) == '\n' {
			line++
		}
	}
	lineEnd := lineStart
	for l.input.At(lineEnd) != '\n' && lineEnd < l.input.RuneCount()-1 {
		lineEnd++
	}
	str += logger.Warn(fmt.Sprintf("%d:%d\t", ln, col))
	str += l.input.Slice(lineStart, lineEnd)
	str += "\n\t"
	str += strings.Repeat(" ", col-1)
	str += logger.Warn("^")
	return str
}