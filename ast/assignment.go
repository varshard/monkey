package ast

import (
	"fmt"
	"github.com/varshard/monkey/token"
)

type AssignmentStatement struct {
	Token      token.Token
	Identifier Identifier
	Value      Expression
}

func (a AssignmentStatement) TokenLiteral() string {
	return a.Token.Literal
}

func (a AssignmentStatement) String() string {
	return fmt.Sprintf("%s = %s;", a.Identifier.String(), a.Value.String())
}

func (a AssignmentStatement) statementNode() {
	panic("implement me")
}
