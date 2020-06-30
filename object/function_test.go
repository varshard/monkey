package object

import (
	"github.com/stretchr/testify/assert"
	"github.com/varshard/monkey/ast"
	"github.com/varshard/monkey/token"
	"testing"
)

func TestFunctionObject_String(t *testing.T) {
	t.Run("Test a function literal", func(t *testing.T) {
		f := FunctionObject{
			Body: ast.BlockStatement{
				Token: token.Token{
					Type:    token.Lbrace,
					Literal: "{",
				},
				Statements: []ast.Statement{
					ast.ReturnStatement{
						Token: token.Token{
							Type:    token.Return,
							Literal: "return",
						},
						Value: ast.IntegerLiteral{
							Token: token.Token{
								Type:    token.Integer,
								Literal: "99",
							},
							Value: 99,
						},
					},
				},
			},
		}

		assert.Equal(t, "fn() {\nreturn 99;\n}", f.String())
	})

	t.Run("Test a named function", func(t *testing.T) {
		x := ast.Identifier{
			Token: token.Token{
				Type:    token.Identifier,
				Literal: "x",
			},
			Name: "x",
		}
		f := FunctionObject{
			Name: "magicNumber",
			Args: []ast.Identifier{x},
			Body: ast.BlockStatement{
				Token: token.Token{
					Type:    token.Lbrace,
					Literal: "{",
				},
				Statements: []ast.Statement{
					ast.ReturnStatement{
						Token: token.Token{
							Type:    token.Return,
							Literal: "return",
						},
						Value: x,
					},
				},
			},
		}

		assert.Equal(t, "fn magicNumber(x) {\nreturn x;\n}", f.String())
	})
}
