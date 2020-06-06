package ast

import (
	"bytes"
	"github.com/varshard/monkey/token"
)

type LetStatement struct {
	Token    token.Token
	Variable *Identifier
	Value    Expression
}

func (s LetStatement) TokenLiteral() string {
	return s.Token.Literal
}

func (s *LetStatement) statementNode() {}

func (s LetStatement) String() string {
	var out bytes.Buffer
	out.WriteString(s.TokenLiteral() + " ")
	out.WriteString(s.Variable.String())
	if s.Value != nil {
		out.WriteString(" = " + s.Value.String())
	}
	out.WriteString(";")

	return out.String()
}
