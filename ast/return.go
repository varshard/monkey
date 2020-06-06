package ast

import (
	"bytes"
	"github.com/varshard/monkey/token"
)

type ReturnStatement struct {
	Token token.Token
	Value Expression
}

func (r ReturnStatement) TokenLiteral() string {
	return r.Token.Literal
}

func (r *ReturnStatement) statementNode() {}

func (r ReturnStatement) String() string {
	var out bytes.Buffer

	out.WriteString(r.TokenLiteral())
	if r.Value != nil {
		out.WriteString(" " + r.Value.String())
	}
	out.WriteString(";")
	return out.String()
}
