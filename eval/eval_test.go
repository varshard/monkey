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
			intObjc, ok := Eval(node).(object.IntegerObject)

			assert.True(t, ok)
			assert.Equal(t, 5, intObjc.Value)
			assert.Equal(t, object.Integer, intObjc.Type())
		})
	})
}
