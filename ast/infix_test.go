package ast

import (
	"github.com/stretchr/testify/assert"
	"github.com/varshard/monkey/token"
	"testing"
)

func TestInfixExpression(t *testing.T) {
	t.Run("Test 1 + 2", func(t *testing.T) {
		infix := InfixExpression{
			Operator: "+",
			Left: &IntegerLiteral{
				Token: token.Token{
					Type:    token.Integer,
					Literal: "1",
				},
				Value: 1,
			},
			Right: &IntegerLiteral{
				Token: token.Token{
					Type:    token.Integer,
					Literal: "2",
				},
			},
			Token: token.Token{
				Type:    token.Plus,
				Literal: "+",
			},
		}

		assert.Equal(t, "(1 + 2)", infix.String())
	})
}
