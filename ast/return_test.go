package ast

import (
	"github.com/stretchr/testify/assert"
	"github.com/varshard/monkey/token"
	"testing"
)

func TestReturnStatement(t *testing.T) {
	t.Run("Test return;", func(t *testing.T) {
		ret := ReturnStatement{
			Token: token.Token{
				Literal: "return",
				Type:    token.Return,
			},
			Value: nil,
		}

		assert.Equal(t, "return;", ret.String())
	})
}
