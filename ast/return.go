package ast

import "github.com/varshard/monkey/token"

type ReturnStatement struct {
	Token token.Token
	Value Expression
}

func (r ReturnStatement) TokenLiteral() string {
	return r.Token.Literal
}

func (r ReturnStatement) statementNode() {}
