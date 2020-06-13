package ast

import (
	"fmt"
	"github.com/varshard/monkey/token"
)

type SuffixExpression struct {
	Token    token.Token
	Left     *Identifier
	Operator string
}

func (s SuffixExpression) TokenLiteral() string {
	return s.Token.Literal
}

func (s SuffixExpression) String() string {
	return fmt.Sprintf("(%s%s)", s.Left.String(), s.Operator)
}

func (s SuffixExpression) expressionNode() {
	panic("implement me")
}
