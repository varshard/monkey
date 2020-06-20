package ast

import (
	"fmt"
	"github.com/varshard/monkey/token"
)

type Assignment struct {
	Token      token.Token
	Identifier *Identifier
	Value      Expression
}

func (a Assignment) TokenLiteral() string {
	return a.Token.Literal
}

func (a Assignment) String() string {
	return fmt.Sprintf("%s = %s", a.Identifier.String(), a.Value.String())
}

func (a Assignment) expressionNode() {
	panic("implement me")
}
