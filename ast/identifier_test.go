package ast

import (
	"github.com/stretchr/testify/assert"
	"github.com/varshard/monkey/token"
	"testing"
)

func TestIdentifier(t *testing.T) {
	t.Run("Test x", func(t *testing.T) {
		identifierTok := token.Token{
			Type:    token.Identifier,
			Literal: "x",
		}

		identifier := Identifier{
			Token: identifierTok,
			Name:  identifierTok.Literal,
		}

		assert.Equal(t, "x", identifier.String())
	})
}
