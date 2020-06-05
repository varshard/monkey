package ast

import (
	"github.com/stretchr/testify/assert"
	"github.com/varshard/monkey/token"
	"testing"
)

func TestLet(t *testing.T) {
	t.Run("Test let x", func(t *testing.T) {
		letTok := token.Token{
			Type:    token.Let,
			Literal: "let",
		}

		x := token.Token{
			Type:    token.Identifier,
			Literal: "x",
		}

		xIdent := Identifier{Name: x.Literal, Token: x}
		let := LetStatement{
			Token:    letTok,
			Variable: &xIdent,
		}

		assert.Equal(t, "let", let.TokenLiteral())
		assert.Equal(t, "let x;", let.String())
	})

	t.Run("Test let x = 10;", func(t *testing.T) {
		letTok := token.Token{
			Type:    token.Let,
			Literal: "let",
		}

		x := token.Token{
			Type:    token.Identifier,
			Literal: "x",
		}

		xIdent := Identifier{Name: x.Literal, Token: x}
		let := LetStatement{
			Token:    letTok,
			Variable: &xIdent,
		}

		assert.Equal(t, "let x = 10;", let.String())
	})
}
