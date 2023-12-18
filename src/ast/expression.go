package ast

import "dynamite/src/tokens"

type IdentifierExpNode struct {
	Token tokens.Token
	Value string
}
func(node *IdentifierExpNode) TokenLiteral() string {
	return node.Token.Literal
}
func(node *IdentifierExpNode) String() string {
	return node.Value
}
func(node *IdentifierExpNode) expressionNode() {}