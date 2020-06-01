package ast

import (
	"github.com/stretchr/testify/assert"
	"github.com/varshard/monkeyinterpreter/token"
	"testing"
)

func TestLet(t *testing.T) {
	t.Run("Test let x = integer", func(t *testing.T) {
		letTok := token.Token{
			Type:    token.Let,
			Literal: "let",
		}

		x := token.Token{
			Type:    token.Identifier,
			Literal: "x",
		}

		let := LetStatement{
			Token:    letTok,
			Variable: x,
		}

		assert.Equal(t, "let", let.TokenLiteral())
	})
}
