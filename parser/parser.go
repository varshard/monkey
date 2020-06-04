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
	} else if p.currTok.Type == token.Return {
		statement = p.parseReturn()
	}

	return statement
}

func (p *Parser) readSemicolon() {
	// Every statement must be terminated with a ;
	if p.nextTok.Type != token.Semicolon {
		p.Errors = append(p.Errors, errors.New(fmt.Sprintf("missing a semicolon at %d:%d", p.currTok.Line, p.currTok.Col)))
	}
	p.advanceToken()
}

func (p *Parser) parseLet() ast.LetStatement {
	s := ast.LetStatement{
		Token: p.currTok,
	}

	p.advanceToken()
	if p.currTok.Type == token.Identifier {
		s.Variable = p.currTok
	} else {
		p.Errors = append(p.Errors, errors.New(fmt.Sprintf("expected identifier at %d:%d", s.Token.Line, s.Token.Col)))
	}

	if p.nextTok.Type == token.Assign {
		p.advanceToken()
		if p.currTok.Type != token.Assign {
			p.Errors = append(p.Errors, errors.New(fmt.Sprintf("expected = at %d:%d", s.Token.Line, s.Token.Col)))
		}

		p.advanceToken()
		// TODO: read expression as Value
	}
	p.readSemicolon()

	return s
}

func (p *Parser) parseReturn() ast.ReturnStatement {
	s := ast.ReturnStatement{
		Token: p.currTok,
	}

	p.advanceToken()
	// TODO: read expression as Value
	p.readSemicolon()

	return s
}
