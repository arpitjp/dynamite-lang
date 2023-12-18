package ast

type Node interface {
	TokenLiteral() string
	String() string
}

type Statement interface {
	Node // embedded interface
	statementNode()
}

type Expression interface {
	Node // embedded interface
	expressionNode()
}
