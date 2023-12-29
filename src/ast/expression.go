package ast

import (
	"dynamite/src/tokens"
	"strings"
)

type IdentifierExpNode struct {
	Token tokens.Token
	Value string
}

func (node *IdentifierExpNode) TokenLiteral() string {
	return node.Token.Literal
}
func (node *IdentifierExpNode) String() string {
	return node.Value
}
func (node *IdentifierExpNode) expressionNode() {}

type IntegerExpressionNode struct {
	Token tokens.Token
	Value int64
}

func (node *IntegerExpressionNode) TokenLiteral() string {
	return node.Token.Literal
}
func (node *IntegerExpressionNode) String() string {
	return node.TokenLiteral()
}
func (node *IntegerExpressionNode) expressionNode() {}

type PrefixExpressionNode struct {
	Token tokens.Token // the prefix token, e.g. !
	Operator string
	Right Expression
}

func (node *PrefixExpressionNode) TokenLiteral() string {
	return node.Token.Literal
}
func (node *PrefixExpressionNode) String() string {
	var str strings.Builder

	str.WriteString("(")
	str.WriteString(node.Operator)
	str.WriteString(node.Right.String())
	str.WriteString(")")
	
	return str.String()
}
func (node *PrefixExpressionNode) expressionNode() {}

type InfixExpressionNode struct {
	Token tokens.Token // the infix or binary token, e.g. +, -
	Left Expression
	Operator string
	Right Expression
}

func (node *InfixExpressionNode) TokenLiteral() string {
	return node.Token.Literal
}
func (node *InfixExpressionNode) String() string {
	var str strings.Builder

	str.WriteString("(")
	str.WriteString(node.Left.String())
	str.WriteString(" ")
	str.WriteString(node.Operator)
	str.WriteString(" ")
	str.WriteString(node.Right.String())
	str.WriteString(")")

	return str.String()
}
func (node *InfixExpressionNode) expressionNode() {}