package ast

import "strings"

type Program struct {
	Statements []Statement
}

func(node *Program) TokenLiteral() string {
	if len(node.Statements) != 0 {
		return node.Statements[0].TokenLiteral()
	}
	return ""
}
func (node *Program) String() string {
	var str strings.Builder
	for _, stmt := range node.Statements {
		str.WriteString(stmt.String() + ";\n")
	}
	return str.String()
}