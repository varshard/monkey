package ast

import "github.com/varshard/monkey/token"

type IntegerLiteral struct {
	Token token.Token
	Value int
}

func (i IntegerLiteral) TokenLiteral() string {
	return i.Token.Literal
}

func (i IntegerLiteral) String() string {
	return i.TokenLiteral()
}

func (i *IntegerLiteral) expressionNode() {
	panic("implement me")
}
