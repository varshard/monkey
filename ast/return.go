package ast

import "github.com/varshard/monkeyinterpreter/token"

type ReturnStatement struct {
	Token token.Token
	Value Expression
}

func (r ReturnStatement) TokenLiteral() string {
	return r.Token.Literal
}
