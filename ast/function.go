package ast

import (
	"fmt"
	"github.com/varshard/monkey/token"
	"strings"
)

type FunctionLiteral struct {
	Token      token.Token
	Parameters []Identifier
	Body       *BlockStatement
}

func (f FunctionLiteral) TokenLiteral() string {
	return f.Token.Literal
}

func (f FunctionLiteral) String() string {

	params := make([]string, 0)

	for _, param := range f.Parameters {
		params = append(params, param.String())
	}

	return fmt.Sprintf("%s(%s) %s", f.TokenLiteral(), strings.Join(params, ", "), f.Body.String())
}

func (f FunctionLiteral) expressionNode() {
	panic("implement me")
}
