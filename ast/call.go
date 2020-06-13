package ast

import (
	"fmt"
	"github.com/varshard/monkey/token"
	"strings"
)

type CallExpression struct {
	Token      token.Token // (
	Function   Expression  // Function name or function literal
	Parameters []Expression
}

func (c CallExpression) TokenLiteral() string {
	return c.Token.Literal
}

func (c CallExpression) String() string {
	expressions := make([]string, 0)
	for _, exp := range c.Parameters {
		expressions = append(expressions, exp.String())
	}

	return fmt.Sprintf("%s(%s)", c.Function.String(), strings.Join(expressions, ", "))
}

func (c CallExpression) expressionNode() {
	panic("implement me")
}

func (c *CallExpression) PushExpressions(exp Expression) {
	c.Parameters = append(c.Parameters, exp)
}
