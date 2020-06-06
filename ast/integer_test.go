package ast

import (
	"github.com/stretchr/testify/assert"
	"github.com/varshard/monkey/token"
	"testing"
)

func TestIntegerLiteral_String(t *testing.T) {
	integer := IntegerLiteral{
		Token: token.Token{
			Type:    token.Integer,
			Literal: "10",
		},
		Value: 10,
	}

	assert.Equal(t, "10", integer.String())
}
