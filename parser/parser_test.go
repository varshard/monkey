package parser

import (
	"github.com/stretchr/testify/assert"
	"github.com/varshard/monkey/ast"
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
			let, ok := program.Statements[0].(*ast.LetStatement)

			assert.True(t, ok)

			assert.Equal(t, "10", let.Value.TokenLiteral())
		})

		t.Run("Test let integer", func(t *testing.T) {
			statement := "let x = 10;"

			p := New(statement)
			assert.NotNil(t, p.lexer)

			program := p.ParseProgram()

			assert.Equal(t, 1, len(program.Statements))
			let, ok := program.Statements[0].(*ast.LetStatement)

			assert.True(t, ok)

			assert.Equal(t, "10", let.Value.TokenLiteral())

			integer, ok := let.Value.(*ast.IntegerLiteral)

			assert.True(t, ok)
			assert.Equal(t, 10, integer.Value)
		})

		t.Run("Test let integer", func(t *testing.T) {
			statement := "let x = -5;"

			p := New(statement)
			assert.NotNil(t, p.lexer)

			program := p.ParseProgram()

			assert.Equal(t, 1, len(program.Statements))
			let, ok := program.Statements[0].(*ast.LetStatement)

			assert.True(t, ok)

			integer, ok := let.Value.(*ast.IntegerLiteral)

			assert.True(t, ok)
			assert.Equal(t, -5, integer.Value)
			assert.Equal(t, "-5", let.Value.TokenLiteral())
		})

		t.Run("Test let identifier", func(t *testing.T) {
			statement := "let x = 10;let y = x;"

			p := New(statement)
			assert.NotNil(t, p.lexer)

			program := p.ParseProgram()

			assert.Equal(t, 2, len(program.Statements))
			for _, s := range program.Statements {
				_, ok := s.(*ast.LetStatement)
				assert.True(t, ok)
			}

			let, ok := program.Statements[1].(*ast.LetStatement)

			assert.True(t, ok)
			assert.Equal(t, "y", let.Variable.TokenLiteral())
			assert.Equal(t, "x", let.Value.TokenLiteral())
		})

		t.Run("Test error invalid let", func(t *testing.T) {
			statement := "let;"

			p := New(statement)

			p.ParseProgram()

			e := p.Errors[0]
			assert.Equal(t, "expected identifier, but got ; at 1:4", e.Error())
		})

		t.Run("Test let without semicolon", func(t *testing.T) {
			statement := "let x = 5"

			p := New(statement)
			p.ParseProgram()

			assert.Equal(t, 1, len(p.Errors))
			assert.Equal(t, "expected ;, but got Eof at 1:9", p.Errors[0].Error())
		})

		t.Run("Test let without =", func(t *testing.T) {
			statement := "let x;"

			p := New(statement)
			program := p.ParseProgram()

			_, ok := program.Statements[0].(*ast.LetStatement)

			assert.True(t, ok)
		})
	})

	t.Run("Test parsing return statement", func(t *testing.T) {
		t.Run("Test return an integer", func(t *testing.T) {
			statement := "return 3;"

			p := New(statement)

			program := p.ParseProgram()

			assert.Equal(t, 1, len(program.Statements))
			ret, ok := program.Statements[0].(*ast.ReturnStatement)

			assert.True(t, ok)
			assert.Equal(t, "3", ret.Value.TokenLiteral())
		})
	})
}
