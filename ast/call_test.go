package ast

import (
	"github.com/stretchr/testify/assert"
	"github.com/varshard/monkey/token"
	"testing"
)

func TestCallExpression_String(t *testing.T) {
	t.Run("Test calling a named function", func(t *testing.T) {
		call := CallExpression{
			Token: token.Token{
				Type:    token.Lparen,
				Literal: "(",
			},
			Function: Identifier{
				Token: token.Token{
					Type:    token.Identifier,
					Literal: "add",
				},
				Name: "add",
			},
			Parameters: []Expression{
				Identifier{
					Token: token.Token{
						Type:    token.Identifier,
						Literal: "x",
					},
					Name: "x",
				},
				PrefixExpression{
					Token: token.Token{
						Type:    token.Minus,
						Literal: "-",
					},
					Operator: "-",
					Right: IntegerLiteral{
						Token: token.Token{
							Type:    token.Integer,
							Literal: "2",
						},
						Value: 2,
					},
				},
			},
		}

		assert.Equal(t, "add(x, (-2))", call.String())
	})

	t.Run("Test calling a function literal", func(t *testing.T) {
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
		call := CallExpression{
			Token: token.Token{
				Type:    token.Lparen,
				Literal: "(",
			},
			Function: FunctionLiteral{
				Token: token.Token{
					Type:    token.Function,
					Literal: "fn",
				},
				Parameters: []Identifier{a, b},
				Body: &BlockStatement{
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
								Left:     a,
								Right:    b,
								Token: token.Token{
									Type:    token.Plus,
									Literal: "+",
								},
							},
						},
					},
				},
			},
			Parameters: []Expression{
				Identifier{
					Token: token.Token{
						Type:    token.Identifier,
						Literal: "x",
					},
					Name: "x",
				},
				PrefixExpression{
					Token: token.Token{
						Type:    token.Minus,
						Literal: "-",
					},
					Operator: "-",
					Right: IntegerLiteral{
						Token: token.Token{
							Type:    token.Integer,
							Literal: "2",
						},
						Value: 2,
					},
				},
			},
		}

		assert.Equal(t, "fn(a, b) {\nreturn (a + b);\n}(x, (-2))", call.String())
	})
}
