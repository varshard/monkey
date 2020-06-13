package ast

import (
	"fmt"
	"github.com/varshard/monkey/token"
)

type PrefixExpression struct {
	Token    token.Token
	Operator string
	Right    Expression
}

func (p PrefixExpression) TokenLiteral() string {
	return p.Token.Literal
}

func (p PrefixExpression) expressionNode() {
	panic("implement me")
}

func (p PrefixExpression) String() string {
	return fmt.Sprintf("(%s%s)", p.Operator, p.Right.String())
}
