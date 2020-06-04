package parser

import (
	"errors"
	"fmt"
	"github.com/varshard/monkeyinterpreter/ast"
	"github.com/varshard/monkeyinterpreter/lexer"
	"github.com/varshard/monkeyinterpreter/token"
)

type Parser struct {
	lexer   *lexer.Lexer
	currTok token.Token
	nextTok token.Token
	Errors  []error
}

func New(code string) *Parser {
	parser := &Parser{
		lexer:  lexer.New(code),
		Errors: make([]error, 0),
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
		}
		break
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
	}

	return statement
}

func (p *Parser) readSemicolon() bool {
	// Every statement must be terminated with a ;
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
		s.Variable = p.currTok
	}

	if p.peekToken(token.Assign) {
		p.advanceToken()
		p.advanceToken()
		// TODO: read expression as Value
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
	// TODO: read expression as Value
	if !p.readSemicolon() {
		return nil
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
