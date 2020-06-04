package ast

import "github.com/varshard/monkeyinterpreter/token"

type LetStatement struct {
	Token    token.Token
	Variable *Identifier
	Value    Expression
}

func (s LetStatement) TokenLiteral() string {
	return s.Token.Literal
}

func (s LetStatement) statementNode() {}