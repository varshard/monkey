package ast

import (
	"github.com/varshard/monkey/token"
	"strings"
)

type Node interface {
	TokenLiteral() string
	String() string
}

type Statement interface {
	Node
	statementNode()
}

type Expression interface {
	Node
	expressionNode()
}

type Program struct {
	Statements []Statement
}

func NewProgram() *Program {
	p := &Program{
		Statements: make([]Statement, 0),
	}
	return p
}

func (p Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	}
	return ""
}

func (p *Program) PushStatement(s Statement) {
	p.Statements = append(p.Statements, s)
}

func (p *Program) String() string {
	lines := make([]string, 0)
	for _, s := range p.Statements {
		lines = append(lines, s.String())
	}

	return strings.Join(lines, "\n")
}

type ExpressionStatement struct {
	Token      token.Token
	Expression Expression
}

func (es *ExpressionStatement) statementNode() {
	panic("Implement me")
}
func (es ExpressionStatement) TokenLiteral() string {
	return es.Token.Literal
}
func (es ExpressionStatement) String() string {
	if es.Expression != nil {
		return es.Expression.String() + ";"
	}
	return ""
}
