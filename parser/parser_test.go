package parser

import (
	"github.com/stretchr/testify/assert"
	"github.com/varshard/monkeyinterpreter/ast"
	"testing"
)

func TestParser(t *testing.T) {
	t.Run("Test New", func(t *testing.T) {
		p := New("")
		assert.NotNil(t, p.lexer)
	})

	t.Run("Test parsing let statement", func(t *testing.T) {
		t.Run("Test let integer", func(t *testing.T) {
			statement := "let x = 10;"

			p := New(statement)
			assert.NotNil(t, p.lexer)

			program := p.ParseProgram()

			assert.Equal(t, 1, len(program.Statements))
			_, ok := program.Statements[0].(ast.LetStatement)

			assert.True(t, ok)
		})

		t.Run("Test error invalid let", func(t *testing.T) {
			statement := "let;"

			p := New(statement)

			p.ParseProgram()

			assert.Equal(t, 1, len(p.errors))

			e := p.errors[0]
			// TODO: support column
			assert.Equal(t, "Expected identifier at 1:3", e.Error())
		})
	})
}
