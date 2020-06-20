package ast

import (
	"github.com/stretchr/testify/assert"
	"github.com/varshard/monkey/token"
	"testing"
)

func TestAssignment_String(t *testing.T) {
	a := Assignment{
		Token: token.Token{
			Literal: "=",
			Type:    token.Assign,
		},
		Identifier: &Identifier{
			Token: token.Token{
				Type:    token.Identifier,
				Literal: "a",
			},
			Name: "a",
		},
		Value: IntegerLiteral{
			Token: token.Token{
				Type:    token.Integer,
				Literal: "7",
			},
			Value: 7,
		},
	}

	assert.Equal(t, "a = 7", a.String())
}
