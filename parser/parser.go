package parser

import (
	"errors"
	"fmt"
	"github.com/varshard/monkey/ast"
	"github.com/varshard/monkey/lexer"
	"github.com/varshard/monkey/token"
	"strconv"
)

type (
	prefixParseFn func() ast.Expression
	infixParseFn  func(expression ast.Expression) ast.Expression
)

type Parser struct {
	lexer   *lexer.Lexer
	currTok token.Token
	nextTok token.Token
	Errors  []error

	prefixParseFns map[token.TokenType]prefixParseFn
	infixParseFn   map[token.TokenType]infixParseFn
}

func New(code string) *Parser {
	parser := &Parser{
		lexer:  lexer.New(code),
		Errors: make([]error, 0),
	}

	// default, -, ++, --
	parser.prefixParseFns = map[token.TokenType]prefixParseFn{
		token.Identifier: parser.parseIdentifier,
		token.Integer:    parser.parseInteger,
		token.Minus:      parser.parseNegative,
		//	"++": func() ast.Expression {
		//
		//	},
		//	"--": func() ast.Expression {
		//
		//	},
	}
	// advance twice to set curr and next
	parser.advanceToken()
	parser.advanceToken()

	return parser
}

func (p *Parser) ParseProgram() *ast.Program {
	program := ast.NewProgram()

	for p.currTok.Type != token.Eof {
		statement := p.readTokens()

		if statement != nil {
			program.PushStatement(statement)
			p.advanceToken()
		} else {
			break
		}
	}

	return program
}

func (p *Parser) advanceToken() {
	p.currTok = p.nextTok
	p.nextTok = p.lexer.NextToken()
}

func (p *Parser) readTokens() ast.Statement {
	var statement ast.Statement
	if p.currTok.Type == token.Let {
		statement = p.parseLet()
	} else if p.currTok.Type == token.Return {
		statement = p.parseReturn()
	} else {
		statement = p.parseExpressionStatement()
	}

	return statement
}

func (p *Parser) readSemicolon() bool {
	if !p.peekToken(token.Semicolon) {
		p.peekError(token.Semicolon)
		return false
	}
	p.advanceToken()
	return true
}

func (p *Parser) parseLet() *ast.LetStatement {
	s := ast.LetStatement{
		Token: p.currTok,
	}

	if !p.peekToken(token.Identifier) {
		p.peekError(token.Identifier)
		return nil
	}
	p.advanceToken()
	if p.currTok.Type == token.Identifier {
		identifier := ast.Identifier{
			Token: p.currTok,
			Name:  p.currTok.Literal,
		}
		s.Variable = &identifier
	}

	if p.peekToken(token.Assign) {
		// Skip =
		p.advanceToken()
		p.advanceToken()
		s.Value = p.parseExpression()
	}
	if !p.readSemicolon() {
		return nil
	}

	return &s
}

func (p *Parser) parseReturn() *ast.ReturnStatement {
	s := ast.ReturnStatement{
		Token: p.currTok,
	}

	p.advanceToken()
	s.Value = p.parseExpression()
	if !p.readSemicolon() {
		return nil
	}
	return &s
}

func (p *Parser) parseExpression() ast.Expression {
	prefix := p.prefixParseFns[p.currTok.Type]
	if prefix == nil {
		return nil
	}
	leftExp := prefix()

	return leftExp
}

func (p *Parser) parseIdentifier() ast.Expression {
	return &ast.Identifier{Token: p.currTok, Name: p.currTok.Literal}
}

func (p *Parser) parseNegative() ast.Expression {
	if p.peekToken(token.Integer) {
		p.advanceToken()
		integer := p.parseInteger().(*ast.IntegerLiteral)
		integer.Value = -1 * integer.Value
		return integer
	}
	// TODO: parse float
	return nil
}

func (p *Parser) parseInteger() ast.Expression {
	value, err := strconv.Atoi(p.currTok.Literal)
	if err != nil {
		p.Errors = append(p.Errors, err)
	}
	return &ast.IntegerLiteral{Token: p.currTok, Value: value}
}

func (p *Parser) parseExpressionStatement() *ast.ExpressionStatement {
	s := ast.ExpressionStatement{
		Token:      p.currTok,
		Expression: p.parseExpression(),
	}

	return &s
}

func (p *Parser) peekToken(target token.TokenType) bool {
	return p.nextTok.Type == target
}

func (p *Parser) peekError(target token.TokenType) {
	nextTok := p.nextTok
	if nextTok.Type != target {
		p.Errors = append(p.Errors, errors.New(fmt.Sprintf("expected %s, but got %s at %d:%d", target, nextTok.Type, nextTok.Line, nextTok.Col)))
	}
}
