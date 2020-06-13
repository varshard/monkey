package ast

import (
	"github.com/stretchr/testify/assert"
	"github.com/varshard/monkey/token"
	"testing"
)

func TestCallExpression_String(t *testing.T) {
	call := CallExpression{
		Token: token.Token{
			Type:    token.Identifier,
			Literal: "add",
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
}
