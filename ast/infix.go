package ast

import (
	"fmt"
	"github.com/varshard/monkey/token"
)

type InfixExpression struct {
	Operator string
	Left     Expression
	Right    Expression
	Token    token.Token
}

func (i InfixExpression) TokenLiteral() string {
	return i.Token.Literal
}

func (i InfixExpression) String() string {
	return fmt.Sprintf("(%s %s %s)", i.Left.String(), i.Operator, i.Right.String())
}

func (i InfixExpression) expressionNode() {
	panic("implement me")
}
