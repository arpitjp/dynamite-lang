package ast

import (
	"dynamite/src/tokens"
	"strings"
)

// LET statement --------
type LetStatement struct {
	Token tokens.Token
	Name *IdentifierExpNode
	Value Expression
}
func(node *LetStatement) TokenLiteral() string {
	return node.Token.Literal
}
func(node *LetStatement) String() string {
	var str strings.Builder
	str.WriteString(node.TokenLiteral())
	str.WriteString(" ")
	str.WriteString(node.Name.String())
	str.WriteString(" = ")
	str.WriteString(node.Value.String())
	return str.String()
}
func(node *LetStatement) statementNode() {}