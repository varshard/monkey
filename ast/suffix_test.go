package ast

import (
	"github.com/stretchr/testify/assert"
	"github.com/varshard/monkey/token"
	"testing"
)

func TestSuffix(t *testing.T) {
	s := SuffixExpression{
		Token: token.Token{
			Type:    token.Decrement,
			Literal: "--",
		},
		Left: &Identifier{
			Token: token.Token{
				Type:    token.Identifier,
				Literal: "x",
			},
			Name: "x",
		},
		Operator: "--",
	}

	assert.Equal(t, "(x--)", s.String())
}
