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
	errors  []error
}

func New(code string) *Parser {
	parser := &Parser{
		lexer:  lexer.New(code),
		errors: make([]error, 0),
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
		}
		p.advanceToken()
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
	}

	return statement
}

func (p *Parser) parseLet() ast.LetStatement {
	s := ast.LetStatement{
		Token: p.currTok,
	}

	p.advanceToken()
	if p.currTok.Type == token.Identifier {
		s.Variable = p.currTok
	} else {
		p.errors = append(p.errors, errors.New(fmt.Sprintf("Expected identifier at %d:%d", s.Token.Line, s.Token.Col)))
	}
	p.advanceToken()
	// TODO: read expression as Value
	return s
}
