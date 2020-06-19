package parser

import (
	"errors"
	"fmt"
	"github.com/varshard/monkey/ast"
	"github.com/varshard/monkey/lexer"
	"github.com/varshard/monkey/token"
	"strconv"
)

// TODO: parse if, else, else if, loop
const (
	_ int = iota
	LOWEST
	EQUALS
	LESSGREATER
	SUM
	PRODUCT
	CALL
)

var precedences = map[token.TokenType]int{
	token.Equal:         EQUALS,
	token.NotEqual:      EQUALS,
	token.LessThan:      LESSGREATER,
	token.LessThanEqual: LESSGREATER,
	token.MoreThan:      LESSGREATER,
	token.MoreThanEqual: LESSGREATER,
	token.Plus:          SUM,
	token.Minus:         SUM,
	token.Multiply:      PRODUCT,
	token.Divide:        PRODUCT,
	token.Lparen:        CALL,
}

type (
	prefixParseFn func() ast.Expression
	infixParseFn  func(expression ast.Expression) ast.Expression
	suffixParseFn func(identifier *ast.Identifier) ast.Expression
)

type Parser struct {
	lexer   *lexer.Lexer
	currTok token.Token
	nextTok token.Token
	Errors  []error

	prefixParseFns map[token.TokenType]prefixParseFn
	infixParseFns  map[token.TokenType]infixParseFn
	suffixParseFns map[token.TokenType]suffixParseFn
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
		token.Decimal:    parser.parseDecimal,
		token.Minus:      parser.parsePrefix,
		token.True:       parser.parseBool,
		token.False:      parser.parseBool,
		token.Bang:       parser.parsePrefix,
		token.Increment:  parser.parsePrefix,
		token.Decrement:  parser.parsePrefix,
		token.Lparen:     parser.parseGroupedExpression,
		token.Function:   parser.parseFunction,
	}

	parser.infixParseFns = map[token.TokenType]infixParseFn{
		token.Plus:          parser.parseInfixExpression,
		token.Minus:         parser.parseInfixExpression,
		token.Multiply:      parser.parseInfixExpression,
		token.Divide:        parser.parseInfixExpression,
		token.MoreThan:      parser.parseInfixExpression,
		token.MoreThanEqual: parser.parseInfixExpression,
		token.LessThan:      parser.parseInfixExpression,
		token.LessThanEqual: parser.parseInfixExpression,
		token.Equal:         parser.parseInfixExpression,
		token.NotEqual:      parser.parseInfixExpression,
	}

	parser.suffixParseFns = map[token.TokenType]suffixParseFn{
		token.Increment: parser.parseSuffix,
		token.Decrement: parser.parseSuffix,
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
	} else if p.currTok.Type == token.Identifier && p.nextTok.Type == token.Assign {
		statement = p.parseAssignment()
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

	if !p.expectToken(token.Identifier) {
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
		s.Value = p.parseExpression(LOWEST)
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
	if !p.peekToken(token.Semicolon) {
		p.advanceToken()
		s.Value = p.parseExpression(LOWEST)
	}
	if !p.readSemicolon() {
		return nil
	}
	return &s
}

func (p *Parser) parseAssignment() *ast.AssignmentStatement {
	a := ast.AssignmentStatement{
		Identifier: ast.Identifier{
			Token: p.currTok,
			Name:  p.currTok.Literal,
		},
	}

	if !p.expectToken(token.Assign) {
		return nil
	}

	p.advanceToken()
	a.Token = p.currTok

	p.advanceToken()
	a.Value = p.parseExpression(LOWEST)

	if !p.readSemicolon() {
		return nil
	}
	return &a
}

func (p *Parser) parseExpressionStatement() *ast.ExpressionStatement {
	s := ast.ExpressionStatement{
		Token:      p.currTok,
		Expression: p.parseExpression(LOWEST),
	}

	if !p.readSemicolon() {
		return nil
	}
	return &s
}

func (p *Parser) parseFunctionCall(function ast.Expression) ast.Expression {
	call := ast.CallExpression{
		Token:    p.currTok,
		Function: function,
	}

	if p.peekToken(token.Rparen) {
		p.advanceToken()
		return call
	}

	p.advanceToken()
	call.PushExpressions(p.parseExpression(LOWEST))

	for p.peekToken(token.Comma) {
		p.advanceToken()
		p.advanceToken()
		call.PushExpressions(p.parseExpression(LOWEST))
	}

	if !p.expectToken(token.Rparen) {
		return nil
	}
	p.advanceToken()
	return call
}

func (p *Parser) parseIdentifier() ast.Expression {
	identifier := p.parseRawIdentifier()
	suffix := p.suffixParseFns[p.nextTok.Type]
	if suffix != nil {
		p.advanceToken()
		return suffix(identifier.(*ast.Identifier))
	} else {
		if p.peekToken(token.Lparen) {
			p.advanceToken()
			return p.parseFunctionCall(identifier)
		}
		return identifier
	}
}

func (p *Parser) parseExpression(precedence int) ast.Expression {
	currTok := p.currTok

	var leftExp ast.Expression
	prefix := p.prefixParseFns[currTok.Type]
	if prefix == nil {
		p.Errors = append(p.Errors, errors.New(fmt.Sprintf("expected expression, but found %s at %d:%d", currTok.Literal, currTok.Line, currTok.Col)))
		return nil
	}
	leftExp = prefix()
	// precedence > p.peekPrecedence will break the recursion and handled automatically when parseExpression is called the next time
	for !p.peekToken(token.Semicolon) && precedence < p.peekPrecedence() {
		infix := p.infixParseFns[p.nextTok.Type]
		if infix == nil {
			return leftExp
		}
		p.advanceToken()
		leftExp = infix(leftExp)
	}
	return leftExp
}

func (p *Parser) parseInfixExpression(left ast.Expression) ast.Expression {
	infix := ast.InfixExpression{
		Operator: p.currTok.Literal,
		Token:    p.currTok,
		Left:     left,
	}

	precedence := precedences[p.currTok.Type]
	p.advanceToken()
	infix.Right = p.parseExpression(precedence)

	return &infix
}

func (p *Parser) parseGroupedExpression() ast.Expression {
	// skip (
	p.advanceToken()
	exp := p.parseExpression(LOWEST)
	if !p.expectToken(token.Rparen) {
		return nil
	}

	// skip )
	p.advanceToken()
	return exp
}

func (p *Parser) parseRawIdentifier() ast.Expression {
	return &ast.Identifier{Token: p.currTok, Name: p.currTok.Literal}
}

func (p *Parser) parsePrefix() ast.Expression {
	prefix := &ast.PrefixExpression{
		Token:    p.currTok,
		Operator: p.currTok.Literal,
	}
	p.advanceToken()
	prefix.Right = p.parseExpression(LOWEST)
	return prefix
}

func (p *Parser) parseInteger() ast.Expression {
	value, err := strconv.Atoi(p.currTok.Literal)
	if err != nil {
		p.Errors = append(p.Errors, err)
	}
	return &ast.IntegerLiteral{Token: p.currTok, Value: value}
}

func (p *Parser) parseDecimal() ast.Expression {
	value, err := strconv.ParseFloat(p.currTok.Literal, 64)
	if err != nil {
		p.Errors = append(p.Errors, err)
	}
	return &ast.DecimalLiteral{Token: p.currTok, Value: value}
}

func (p *Parser) parseBool() ast.Expression {
	boolean, err := strconv.ParseBool(p.currTok.Literal)

	if err != nil {
		p.Errors = append(p.Errors, err)
		return nil
	}

	return &ast.Boolean{
		Token: p.currTok,
		Value: boolean,
	}
}

func (p *Parser) parseFunction() ast.Expression {
	function := ast.FunctionLiteral{
		Token: p.currTok,
	}

	function.Parameters = p.parseFunctionParams()
	function.Body = p.parseBlock()

	// Call function literal
	if p.peekToken(token.Lparen) {
		p.advanceToken()
		return p.parseFunctionCall(function)
	}

	return function
}

func (p *Parser) parseBlock() *ast.BlockStatement {
	if !p.expectToken(token.Lbrace) {
		return nil
	}
	p.advanceToken()
	block := ast.BlockStatement{
		Token: p.currTok,
	}

	for !p.peekToken(token.Rbrace) {
		p.advanceToken()
		statement := p.readTokens()
		if statement != nil {
			block.PushStatement(statement)
		} else {
			break
		}
	}

	if !p.expectToken(token.Rbrace) {
		return nil
	}
	p.advanceToken()

	return &block
}

func (p *Parser) parseFunctionParams() []ast.Identifier {
	params := make([]ast.Identifier, 0)

	if !p.expectToken(token.Lparen) {
		return nil
	}
	// skip (
	p.advanceToken()
	for p.peekToken(token.Identifier) {
		p.advanceToken()
		params = append(params, ast.Identifier{
			Name:  p.currTok.Literal,
			Token: p.currTok,
		})

		if p.peekToken(token.Comma) {
			// skip ,
			p.advanceToken()
		}
	}

	if !p.expectToken(token.Rparen) {
		return nil
	}
	// skip )
	p.advanceToken()

	return params
}

func (p *Parser) parseSuffix(identifier *ast.Identifier) ast.Expression {
	suffix := ast.SuffixExpression{
		Token:    p.currTok,
		Left:     identifier,
		Operator: p.currTok.Literal,
	}

	return &suffix
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

func (p *Parser) peekPrecedence() int {
	if p, ok := precedences[p.nextTok.Type]; ok {
		return p
	}

	return LOWEST
}

func (p *Parser) expectToken(target token.TokenType) bool {
	if !p.peekToken(target) {
		p.peekError(target)
		return false
	}
	return true
}
