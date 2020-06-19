package eval

import (
	"github.com/stretchr/testify/assert"
	"github.com/varshard/monkey/ast"
	"github.com/varshard/monkey/object"
	"github.com/varshard/monkey/token"
	"testing"
)

func Test_Eval(t *testing.T) {
	t.Run("Test eval literal values", func(t *testing.T) {
		t.Run("Test eval an integer literal", func(t *testing.T) {
			node := ast.IntegerLiteral{
				Token: token.Token{
					Type:    token.Integer,
					Literal: "5",
				},
				Value: 5,
			}
			intObj, ok := Eval(node).(object.IntegerObject)

			assert.True(t, ok)
			assert.Equal(t, 5, intObj.Value)
			assert.Equal(t, object.Integer, intObj.Type())
		})

		t.Run("Test eval a decimal literal", func(t *testing.T) {
			node := ast.DecimalLiteral{
				Token: token.Token{
					Type:    token.Decimal,
					Literal: "5.0",
				},
				Value: 5.0,
			}
			obj, ok := Eval(node).(object.DecimalObject)

			assert.True(t, ok)
			assert.Equal(t, 5.0, obj.Value)
			assert.Equal(t, object.Decimal, obj.Type())
		})
	})

	t.Run("Test eval infix", func(t *testing.T) {
		t.Run("Test eval int + int", func(t *testing.T) {
			node := ast.InfixExpression{
				Operator: "+",
				Left: ast.IntegerLiteral{
					Token: token.Token{
						Literal: "3",
						Type:    token.Integer,
					},
					Value: 3,
				},
				Right: ast.IntegerLiteral{
					Token: token.Token{
						Literal: "2",
						Type:    token.Integer,
					},
					Value: 2,
				},
				Token: token.Token{
					Literal: "+",
					Type:    token.Plus,
				},
			}

			obj, ok := Eval(node).(object.IntegerObject)

			assert.True(t, ok)
			assert.Equal(t, 5, obj.Value)
		})

		t.Run("Test eval decimal + int", func(t *testing.T) {
			node := ast.InfixExpression{
				Operator: "+",
				Left: ast.DecimalLiteral{
					Token: token.Token{
						Literal: "3.5",
						Type:    token.Decimal,
					},
					Value: 3.5,
				},
				Right: ast.IntegerLiteral{
					Token: token.Token{
						Literal: "2",
						Type:    token.Integer,
					},
					Value: 2,
				},
				Token: token.Token{
					Literal: "+",
					Type:    token.Plus,
				},
			}

			obj, ok := Eval(node).(object.DecimalObject)

			assert.True(t, ok)
			assert.Equal(t, 5.5, obj.Value)
		})

		t.Run("Test eval decimal + decimal", func(t *testing.T) {
			node := ast.InfixExpression{
				Operator: "+",
				Left: ast.DecimalLiteral{
					Token: token.Token{
						Literal: "3.5",
						Type:    token.Decimal,
					},
					Value: 3.5,
				},
				Right: ast.DecimalLiteral{
					Token: token.Token{
						Literal: "2.5",
						Type:    token.Decimal,
					},
					Value: 2.5,
				},
				Token: token.Token{
					Literal: "+",
					Type:    token.Plus,
				},
			}

			obj, ok := Eval(node).(object.DecimalObject)

			assert.True(t, ok)
			assert.Equal(t, 6.0, obj.Value)
		})

		t.Run("Test eval int + decimal", func(t *testing.T) {
			node := ast.InfixExpression{
				Operator: "+",
				Left: ast.IntegerLiteral{
					Token: token.Token{
						Literal: "1",
						Type:    token.Integer,
					},
					Value: 1,
				},
				Right: ast.DecimalLiteral{
					Token: token.Token{
						Literal: "2.53",
						Type:    token.Decimal,
					},
					Value: 2.53,
				},
				Token: token.Token{
					Literal: "+",
					Type:    token.Plus,
				},
			}

			obj, ok := Eval(node).(object.DecimalObject)

			assert.True(t, ok)
			assert.Equal(t, 3.53, obj.Value)
		})
	})
}
