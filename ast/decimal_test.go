package ast

import (
	"github.com/stretchr/testify/assert"
	"github.com/varshard/monkey/token"
	"testing"
)

func TestDecimalLiteral_String(t *testing.T) {
	decimal := DecimalLiteral{
		Token: token.Token{
			Type:    token.Decimal,
			Literal: "10.0",
		},
		Value: 10.0,
	}

	assert.Equal(t, "10.0", decimal.String())
}
