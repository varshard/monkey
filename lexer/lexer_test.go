package lexer

import (
	"github.com/stretchr/testify/assert"
	"github.com/varshard/monkey/token"
	"testing"
)

func TestNextToken(t *testing.T) {
	t.Run("symbols", func(t *testing.T) {
		input := "={}().,;!+- */"
		l := New(input)
		tok := l.NextToken()

		assert.Equal(t, token.Assign, tok.Type)

		tok = l.NextToken()
		assert.Equal(t, token.Lbrace, tok.Type)

		tok = l.NextToken()
		assert.Equal(t, token.Rbrace, tok.Type)

		tok = l.NextToken()
		assert.Equal(t, token.Lparen, tok.Type)

		tok = l.NextToken()
		assert.Equal(t, token.Rparen, tok.Type)

		tok = l.NextToken()
		assert.Equal(t, token.Period, tok.Type)

		tok = l.NextToken()
		assert.Equal(t, token.Comma, tok.Type)

		tok = l.NextToken()
		assert.Equal(t, token.Semicolon, tok.Type)

		tok = l.NextToken()
		assert.Equal(t, token.Bang, tok.Type)

		tok = l.NextToken()
		assert.Equal(t, token.Plus, tok.Type)

		tok = l.NextToken()
		assert.Equal(t, token.Minus, tok.Type)

		tok = l.NextToken()
		assert.Equal(t, token.Multiply, tok.Type)

		tok = l.NextToken()
		assert.Equal(t, token.Divide, tok.Type)

		tok = l.NextToken()
		assert.Equal(t, token.Eof, tok.Type)
	})

	t.Run("Equality", func(t *testing.T) {
		l := New("!===>=<=><")

		tok := l.NextToken()
		assert.Equal(t, token.NotEqual, tok.Type)

		tok = l.NextToken()
		assert.Equal(t, token.Equal, tok.Type)

		tok = l.NextToken()
		assert.Equal(t, token.MoreThanEqual, tok.Type)

		tok = l.NextToken()
		assert.Equal(t, token.LessThanEqual, tok.Type)

		tok = l.NextToken()
		assert.Equal(t, token.MoreThan, tok.Type)

		tok = l.NextToken()
		assert.Equal(t, token.LessThan, tok.Type)
	})

	t.Run("let", func(t *testing.T) {
		l := New("let")

		tok := l.NextToken()
		assert.Equal(t, token.Let, tok.Type)
	})

	t.Run("Identifier", func(t *testing.T) {
		l := New("varCount")

		tok := l.NextToken()
		assert.Equal(t, token.Identifier, tok.Type)
		assert.Equal(t, "varCount", tok.Literal)

		l = New("varCount69+")

		tok = l.NextToken()
		assert.Equal(t, token.Identifier, tok.Type)
		assert.Equal(t, "varCount69", tok.Literal)
	})

	t.Run("function", func(t *testing.T) {
		l := New("fn")

		tok := l.NextToken()
		assert.Equal(t, token.Function, tok.Type)
		assert.Equal(t, "fn", tok.Literal)
	})

	t.Run("Literal integer", func(t *testing.T) {
		t.Run("Positive", func(t *testing.T) {
			l := New("123")

			tok := l.NextToken()
			assert.Equal(t, token.Integer, tok.Type)
			assert.Equal(t, "123", tok.Literal)
		})

		t.Run("Negative", func(t *testing.T) {
			l := New("-5")

			tok := l.NextToken()
			assert.Equal(t, token.Minus, tok.Type)
			assert.Equal(t, "-", tok.Literal)

			tok = l.NextToken()
			assert.Equal(t, token.Integer, tok.Type)
			assert.Equal(t, "5", tok.Literal)
		})
	})

	t.Run("Literal floating", func(t *testing.T) {
		t.Run("Positive", func(t *testing.T) {
			l := New("0.27")

			tok := l.NextToken()
			assert.Equal(t, token.Decimal, tok.Type)
			assert.Equal(t, "0.27", tok.Literal)
		})

		t.Run("Negative", func(t *testing.T) {
			l := New("-5.0")

			tok := l.NextToken()
			assert.Equal(t, token.Minus, tok.Type)

			tok = l.NextToken()
			assert.Equal(t, token.Decimal, tok.Type)
		})
	})

	t.Run("Let statement", func(t *testing.T) {
		l := New("let x = 5;")

		tok := l.NextToken()
		assert.Equal(t, token.Let, tok.Type)

		tok = l.NextToken()
		assert.Equal(t, token.Identifier, tok.Type)

		tok = l.NextToken()
		assert.Equal(t, token.Assign, tok.Type)

		tok = l.NextToken()
		assert.Equal(t, token.Integer, tok.Type)

		tok = l.NextToken()
		assert.Equal(t, token.Semicolon, tok.Type)

		tok = l.NextToken()
		assert.Equal(t, token.Eof, tok.Type)
	})

	t.Run("Test return true", func(j *testing.T) {
		l := New("return true;")

		tok := l.NextToken()
		assert.Equal(t, token.Return, tok.Type)

		tok = l.NextToken()
		assert.Equal(t, token.True, tok.Type)

		tok = l.NextToken()
		assert.Equal(t, token.Semicolon, tok.Type)

		tok = l.NextToken()
		assert.Equal(t, token.Eof, tok.Type)
	})

	t.Run("Test return false", func(j *testing.T) {
		l := New("return false;")

		tok := l.NextToken()
		assert.Equal(t, token.Return, tok.Type)

		tok = l.NextToken()
		assert.Equal(t, token.False, tok.Type)

		tok = l.NextToken()
		assert.Equal(t, token.Semicolon, tok.Type)

		tok = l.NextToken()
		assert.Equal(t, token.Eof, tok.Type)
	})

	t.Run("Test return not true", func(j *testing.T) {
		l := New("!true")

		tok := l.NextToken()
		assert.Equal(t, token.Bang, tok.Type)

		tok = l.NextToken()
		assert.Equal(t, token.True, tok.Type)
	})

	t.Run("Test increment", func(j *testing.T) {
		l := New("++x")

		tok := l.NextToken()
		assert.Equal(t, token.Increment, tok.Type)

		tok = l.NextToken()
		assert.Equal(t, token.Identifier, tok.Type)
	})

	t.Run("Test decrement", func(j *testing.T) {
		l := New("--x")

		tok := l.NextToken()
		assert.Equal(t, token.Decrement, tok.Type)

		tok = l.NextToken()
		assert.Equal(t, token.Identifier, tok.Type)
	})
}
