package ast

type Node interface {
	TokenLiteral() string
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
