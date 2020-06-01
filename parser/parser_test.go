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

			assert.True(t, len(program.Statements) == 1)
			_, ok := program.Statements[0].(ast.LetStatement)

			assert.True(t, ok)
		})
	})
}
