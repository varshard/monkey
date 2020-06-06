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

		t.Run("Test let negative", func(t *testing.T) {
			statement := "let x = -5;"

			p := New(statement)
			assert.NotNil(t, p.lexer)

			program := p.ParseProgram()

			assert.Equal(t, 1, len(program.Statements))
			let, ok := program.Statements[0].(*ast.LetStatement)

			assert.True(t, ok)

			prefix, ok := let.Value.(*ast.PrefixExpression)

			assert.True(t, ok)
			assert.Equal(t, "-", prefix.Operator)
			assert.Equal(t, "5", prefix.Right.String())
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

		t.Run("Test let without =", func(t *testing.T) {
			statement := "let x;"

			p := New(statement)
			program := p.ParseProgram()

			_, ok := program.Statements[0].(*ast.LetStatement)

			assert.True(t, ok)
		})

		t.Run("Test let without semicolon", func(t *testing.T) {
			statement := "let x = 5"

			p := New(statement)
			p.ParseProgram()

			assert.Equal(t, 1, len(p.Errors))
			assert.Equal(t, "expected ;, but got Eof at 1:9", p.Errors[0].Error())
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

		t.Run("Test return a boolean", func(t *testing.T) {
			statement := "return true;"

			p := New(statement)
			program := p.ParseProgram()

			assert.Equal(t, 1, len(program.Statements))
			ret, ok := program.Statements[0].(*ast.ReturnStatement)

			assert.True(t, ok)

			boolean, ok := ret.Value.(*ast.Boolean)

			assert.Equal(t, "true", boolean.String())
			assert.True(t, boolean.Value)
		})

		t.Run("Test return false", func(t *testing.T) {
			statement := "return false;"

			p := New(statement)
			program := p.ParseProgram()

			assert.Equal(t, 1, len(program.Statements))
			ret, ok := program.Statements[0].(*ast.ReturnStatement)

			assert.True(t, ok)

			boolean, ok := ret.Value.(*ast.Boolean)

			assert.Equal(t, "false", boolean.String())
			assert.False(t, boolean.Value)
		})

		t.Run("Test return not", func(t *testing.T) {
			statement := "return !false;"

			p := New(statement)
			program := p.ParseProgram()

			assert.Equal(t, 1, len(program.Statements))
			ret, ok := program.Statements[0].(*ast.ReturnStatement)

			assert.True(t, ok)

			prefix, ok := ret.Value.(*ast.PrefixExpression)

			assert.Equal(t, "!", prefix.Operator)
			assert.Equal(t, "false", prefix.Right.String())
		})

		t.Run("Test return ++y", func(t *testing.T) {
			statement := "return ++y;"

			p := New(statement)
			program := p.ParseProgram()

			assert.Equal(t, 1, len(program.Statements))
			ret, ok := program.Statements[0].(*ast.ReturnStatement)

			assert.True(t, ok)

			prefix, ok := ret.Value.(*ast.PrefixExpression)

			assert.Equal(t, "++", prefix.Operator)
			assert.Equal(t, "y", prefix.Right.String())
		})

		t.Run("Test return --x", func(t *testing.T) {
			statement := "return --x;"

			p := New(statement)
			program := p.ParseProgram()

			assert.Equal(t, 1, len(program.Statements))
			ret, ok := program.Statements[0].(*ast.ReturnStatement)

			assert.True(t, ok)

			prefix, ok := ret.Value.(*ast.PrefixExpression)

			assert.Equal(t, "--", prefix.Operator)
			assert.Equal(t, "x", prefix.Right.String())
		})
	})

	t.Run("Test parsing an expression", func(t *testing.T) {
		t.Run("Test suffix", func(t *testing.T) {
			p := New("x++;y--;")
			program := p.ParseProgram()

			assert.Equal(t, 2, len(program.Statements))
			assert.Equal(t, 0, len(p.Errors))

			expression, ok := program.Statements[0].(*ast.ExpressionStatement)
			assert.True(t, ok)

			suffix, ok := expression.Expression.(*ast.SuffixExpression)
			assert.True(t, ok)
			assert.Equal(t, "++", suffix.Operator)
			assert.Equal(t, "x", suffix.Left.String())
			assert.Equal(t, "(x++)", suffix.String())

			expression, ok = program.Statements[0].(*ast.ExpressionStatement)
			assert.True(t, ok)

			suffix, ok = expression.Expression.(*ast.SuffixExpression)
			assert.True(t, ok)
			assert.Equal(t, "--", suffix.Operator)
			assert.Equal(t, "y", suffix.Left.String())
			assert.Equal(t, "(y--)", suffix.String())
		})

		t.Run("Test parsing an invalid expression", func(t *testing.T) {
			p := New("*;")

			p.ParseProgram()

			assert.Equal(t, 1, len(p.Errors))

			assert.Equal(t, "expected expression, but found * at 1:1", p.Errors[0].Error())
		})

		t.Run("Test parsing an invalid expression", func(t *testing.T) {
			p := New("1-;")

			p.ParseProgram()

			assert.Equal(t, "expected expression, but found ; at 1:3", p.Errors[0].Error())
		})

		t.Run("Test parsing infix", func(t *testing.T) {
			p := New("1 + 2;2 - 3; x * 2; 3 * y; 3/7;")
			program := p.ParseProgram()

			checkers := []string{
				"(1 + 2)", "(2 - 3)", "(x * 2)", "(3 * y)", "(3 / 7)",
			}

			assert.Equal(t, 0, len(p.Errors))
			for i, s := range program.Statements {
				expression, ok := s.(*ast.ExpressionStatement)
				assert.True(t, ok)

				infix, ok := expression.Expression.(*ast.InfixExpression)
				assert.True(t, ok)
				assert.Equal(t, checkers[i], infix.String())
			}
		})

		t.Run("Test parsing 2 - 3 * 4", func(t *testing.T) {
			p := New("2 - 3 * 4;")
			program := p.ParseProgram()

			es, ok := program.Statements[0].(*ast.ExpressionStatement)
			assert.True(t, ok)

			infix, ok := es.Expression.(*ast.InfixExpression)
			assert.True(t, ok)

			assert.Equal(t, "2", infix.Left.String())
			assert.Equal(t, "(3 * 4)", infix.Right.String())
			assert.Equal(t, "(2 - (3 * 4))", infix.String())
		})

		t.Run("Test parsing 2 / 3 + 4", func(t *testing.T) {
			p := New("2 / 3 + 4;")
			program := p.ParseProgram()

			es, ok := program.Statements[0].(*ast.ExpressionStatement)
			assert.True(t, ok)

			infix, ok := es.Expression.(*ast.InfixExpression)
			assert.True(t, ok)

			assert.Equal(t, "(2 / 3)", infix.Left.String())
			assert.Equal(t, "4", infix.Right.String())
			assert.Equal(t, "((2 / 3) + 4)", infix.String())
		})
	})
}
