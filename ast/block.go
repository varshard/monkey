package ast

import (
	"fmt"
	"github.com/varshard/monkey/token"
	"strings"
)

type BlockStatement struct {
	Token      token.Token
	Statements []Statement
}

func (b BlockStatement) TokenLiteral() string {
	return b.Token.Literal
}

func (b BlockStatement) String() string {
	lines := make([]string, 0)
	for _, s := range b.Statements {
		lines = append(lines, s.String())
	}

	return fmt.Sprintf("{\n%s\n}", strings.Join(lines, "\n"))
}

func (b BlockStatement) statementNode() {
}

func (b *BlockStatement) PushStatement(s Statement) {
	b.Statements = append(b.Statements, s)
}
