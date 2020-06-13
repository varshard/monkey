package ast

import "github.com/varshard/monkey/token"

// NOTE: Identifier is an expression
// let x = y;
type Identifier struct {
	Token token.Token
	Name  string
}

func (i Identifier) TokenLiteral() string {
	return i.Token.Literal
}

func (i *Identifier) expressionNode() {
	panic("implement me")
}
func (i Identifier) String() string {
	return i.TokenLiteral()
}
