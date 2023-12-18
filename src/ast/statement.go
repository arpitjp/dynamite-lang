package ast

import (
	"dynamite/src/tokens"
	"strings"
)

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


type ReturnStatement struct {
	Token tokens.Token
	Value Expression
}
func (node *ReturnStatement) TokenLiteral() string {
	return node.Token.Literal
}
func(node *ReturnStatement) String() string {
	var str strings.Builder
	str.WriteString(node.TokenLiteral())
	str.WriteString(" ")
	str.WriteString(node.Value.String())
	return str.String()
}
func(node *ReturnStatement) statementNode() {}