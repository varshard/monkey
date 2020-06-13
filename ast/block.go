package ast

import (
	"bytes"
	"fmt"
	"github.com/varshard/monkey/token"
)

type BlockStatement struct {
	Token      token.Token
	Statements []Statement
}

func (b BlockStatement) TokenLiteral() string {
	return b.Token.Literal
}

func (b BlockStatement) String() string {
	buffer := bytes.Buffer{}

	buffer.WriteString("{\n")
	for _, s := range b.Statements {
		buffer.WriteString(fmt.Sprintf("%s\n", s.String()))
	}
	buffer.WriteString("}")

	return buffer.String()
}

func (b BlockStatement) statementNode() {
}

func (b *BlockStatement) PushStatement(s Statement) {
	b.Statements = append(b.Statements, s)
}
