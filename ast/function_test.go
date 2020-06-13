package ast

import (
	"github.com/stretchr/testify/assert"
	"github.com/varshard/monkey/token"
	"testing"
)

func TestFunctionLiteral_String(t *testing.T) {
	a := Identifier{
		Token: token.Token{
			Type:    token.Identifier,
			Literal: "a",
		},
		Name: "a",
	}

	b := Identifier{
		Token: token.Token{
			Type:    token.Identifier,
			Literal: "b",
		},
		Name: "b",
	}
	function := FunctionLiteral{
		Token: token.Token{
			Type:    token.Function,
			Literal: "func",
		},
		Parameters: []Identifier{a, b},
		Body: BlockStatement{
			Token: token.Token{
				Type:    token.Lbrace,
				Literal: "{",
			},
			Statements: []Statement{
				ReturnStatement{
					Token: token.Token{
						Type:    token.Return,
						Literal: "return",
					},
					Value: InfixExpression{
						Operator: "+",
						Left:     &a,
						Right:    &b,
						Token: token.Token{
							Type:    token.Plus,
							Literal: "+",
						},
					},
				},
			},
		},
	}

	assert.Equal(t, "func (a, b) {\nreturn (a + b);\n}", function.String())
}
