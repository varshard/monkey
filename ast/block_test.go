package ast

import (
	"github.com/stretchr/testify/assert"
	"github.com/varshard/monkey/token"
	"testing"
)

func TestBlock_String(t *testing.T) {
	identifier := Identifier{
		Token: token.Token{
			Type:    token.Identifier,
			Literal: "x",
		},
		Name: "x",
	}
	let := LetStatement{
		Token: token.Token{
			Literal: "let",
			Type:    token.Let,
		},
		Variable: &identifier,
		Value: &IntegerLiteral{
			Token: token.Token{
				Type:    token.Integer,
				Literal: "3",
			},
			Value: 3,
		},
	}

	ret := ReturnStatement{
		Token: token.Token{
			Literal: "return",
			Type:    token.Return,
		},
		Value: &identifier,
	}

	block := BlockStatement{
		Token: token.Token{
			Literal: "{",
			Type:    token.Lbrace,
		},
		Statements: []Statement{
			let,
			ret,
		},
	}

	assert.Equal(t, "{\nlet x = 3;\nreturn x;\n}", block.String())
}
