package ast

import "github.com/varshard/monkey/token"

type DecimalLiteral struct {
	Token token.Token
	Value float64
}

func (i DecimalLiteral) TokenLiteral() string {
	return i.Token.Literal
}

func (i DecimalLiteral) String() string {
	return i.TokenLiteral()
}

func (i *DecimalLiteral) expressionNode() {
	panic("implement me")
}
