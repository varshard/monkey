package parser

import (
	"github.com/stretchr/testify/assert"
	"github.com/varshard/monkey/ast"
	"testing"
)

func parseCode(code string) (*Parser, *ast.Program) {
	p := New(code)
	program := p.ParseProgram()

	return p, program
}

type testInput struct {
	input    string
	expected string
}

func TestParser(t *testing.T) {
	t.Run("Test New", func(t *testing.T) {
		p := New("")
		assert.NotNil(t, p.lexer)
	})

	t.Run("Test parsing let statement", func(t *testing.T) {
		t.Run("Test let integer", func(t *testing.T) {
			statement := "let x = 10;"
			_, program := parseCode(statement)

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
			_, program := parseCode(statement)

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
			_, program := parseCode(statement)

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
			p, _ := parseCode("let;")

			e := p.Errors[0]
			assert.Equal(t, "expected identifier, but got ; at 1:4", e.Error())
		})

		t.Run("Test let without =", func(t *testing.T) {
			_, program := parseCode("let x;")

			_, ok := program.Statements[0].(*ast.LetStatement)

			assert.True(t, ok)
		})

		t.Run("Test let without semicolon", func(t *testing.T) {
			p, _ := parseCode("let x = 5")

			assert.Equal(t, 1, len(p.Errors))
			assert.Equal(t, "expected ;, but got Eof at 1:9", p.Errors[0].Error())
		})
	})

	t.Run("Test parsing return statement", func(t *testing.T) {
		t.Run("Test return an integer", func(t *testing.T) {
			_, program := parseCode("return 3;")

			assert.Equal(t, 1, len(program.Statements))
			ret, ok := program.Statements[0].(*ast.ReturnStatement)
			assert.True(t, ok)
			assert.Equal(t, "3", ret.Value.TokenLiteral())
		})

		t.Run("Test return a boolean", func(t *testing.T) {
			_, program := parseCode("return true;")

			assert.Equal(t, 1, len(program.Statements))
			ret, ok := program.Statements[0].(*ast.ReturnStatement)

			assert.True(t, ok)

			boolean, ok := ret.Value.(*ast.Boolean)

			assert.Equal(t, "true", boolean.String())
			assert.True(t, boolean.Value)
		})

		t.Run("Test return false", func(t *testing.T) {
			_, program := parseCode("return false;")

			assert.Equal(t, 1, len(program.Statements))
			ret, ok := program.Statements[0].(*ast.ReturnStatement)

			assert.True(t, ok)

			boolean, ok := ret.Value.(*ast.Boolean)

			assert.Equal(t, "false", boolean.String())
			assert.False(t, boolean.Value)
		})

		t.Run("Test return not", func(t *testing.T) {
			_, program := parseCode("return !false;")

			assert.Equal(t, 1, len(program.Statements))
			ret, ok := program.Statements[0].(*ast.ReturnStatement)

			assert.True(t, ok)

			prefix, ok := ret.Value.(*ast.PrefixExpression)

			assert.Equal(t, "!", prefix.Operator)
			assert.Equal(t, "false", prefix.Right.String())
		})

		t.Run("Test return ++y", func(t *testing.T) {
			tests := []testInput{
				{
					input:    "return ++y;",
					expected: "return (++y);",
				}, {
					input:    "return --x;",
					expected: "return (--x);",
				},
			}

			for _, test := range tests {
				p, program := parseCode(test.input)

				assert.Equal(t, 0, len(p.Errors))
				assert.Equal(t, 1, len(program.Statements))
				ret, ok := program.Statements[0].(*ast.ReturnStatement)

				assert.True(t, ok)

				assert.Equal(t, test.expected, ret.String())
			}
		})

		t.Run("Test return;", func(t *testing.T) {
			_, program := parseCode("return;")

			assert.Equal(t, "return;", program.Statements[0].String())
		})
	})

	t.Run("Test parsing an expression", func(t *testing.T) {
		t.Run("Test suffix", func(t *testing.T) {
			tests := []testInput{
				{
					input:    "x++;",
					expected: "(x++)",
				}, {
					input:    "y--;",
					expected: "(y--)",
				},
			}

			for _, test := range tests {
				p, program := parseCode(test.input)

				assert.Equal(t, 0, len(p.Errors))

				expression, ok := program.Statements[0].(*ast.ExpressionStatement)
				assert.True(t, ok)
				assert.Equal(t, test.expected, expression.String())
			}
		})

		t.Run("Test suffix with infix", func(t *testing.T) {
			tests := []testInput{
				{
					input:    "1 + x++;",
					expected: "(1 + (x++))",
				}, {
					input:    "y-- + -2;",
					expected: "((y--) + (-2))",
				},
			}

			for _, test := range tests {
				p, program := parseCode(test.input)

				assert.Equal(t, 0, len(p.Errors))

				expression, ok := program.Statements[0].(*ast.ExpressionStatement)
				assert.True(t, ok)
				assert.Equal(t, test.expected, expression.String())
			}
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

		t.Run("Test parsing precedence", func(t *testing.T) {
			tests := []testInput{
				{
					input:    "2 - 3 * 4;",
					expected: "(2 - (3 * 4))",
				}, {
					input:    "2 / 3 + 4;",
					expected: "((2 / 3) + 4)",
				},
			}
			for _, test := range tests {
				p, program := parseCode(test.input)

				assert.Equal(t, 0, len(p.Errors))
				es, ok := program.Statements[0].(*ast.ExpressionStatement)
				assert.True(t, ok)

				infix, ok := es.Expression.(*ast.InfixExpression)
				assert.True(t, ok)

				assert.Equal(t, test.expected, infix.String())
			}
		})

		t.Run("Test parsing grouped expressions", func(t *testing.T) {
			tests := []testInput{
				{
					input:    "(2 - 3) * 4;",
					expected: "((2 - 3) * 4)",
				},
				{
					input:    "2 / (3 + 4);",
					expected: "(2 / (3 + 4))",
				}, {
					input:    "(2 + 3) + 4;",
					expected: "((2 + 3) + 4)",
				},
				{
					input:    "(2 + 3 - 4);",
					expected: "((2 + 3) - 4)",
				},
			}
			for _, test := range tests {
				p, program := parseCode(test.input)

				assert.Equal(t, 0, len(p.Errors))
				es, ok := program.Statements[0].(*ast.ExpressionStatement)
				assert.True(t, ok)

				infix, ok := es.Expression.(*ast.InfixExpression)
				assert.True(t, ok)

				assert.Equal(t, test.expected, infix.String())
			}
		})

		t.Run("Test parsing prefix expression in a grouped expression", func(t *testing.T) {
			tests := []testInput{
				{
					input:    "(2);",
					expected: "2",
				},
			}

			for _, test := range tests {
				p, program := parseCode(test.input)

				assert.Equal(t, 0, len(p.Errors))
				es, ok := program.Statements[0].(*ast.ExpressionStatement)
				assert.True(t, ok)

				assert.Equal(t, test.expected, es.String())
			}
		})
	})
}
