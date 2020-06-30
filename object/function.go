package object

import (
	"bytes"
	"fmt"
	"github.com/varshard/monkey/ast"
	"strings"
)

type FunctionObject struct {
	Name  string
	Args  []ast.Identifier
	Body  ast.BlockStatement
	Scope *Scope
}

func (f FunctionObject) Type() ObjectType {
	return FUNCTION
}

func (f FunctionObject) String() string {
	out := bytes.Buffer{}

	out.WriteString("fn")

	if f.Name != "" {
		out.WriteString(fmt.Sprintf(" %s", f.Name))
	}

	args := make([]string, 0)
	for _, arg := range f.Args {
		args = append(args, arg.Name)
	}
	out.WriteString(fmt.Sprintf("(%s) ", strings.Join(args, ", ")))
	out.WriteString(f.Body.String())

	return out.String()
}
