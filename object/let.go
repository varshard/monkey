package object

import (
	"errors"
	"fmt"
)

func NewLet(variable string, scope *Scope, value Object) (Let, error) {
	if !scope.IsDeclared(variable) {
		scope.Set(variable, value)
		identifier := Identifier{
			Name:  variable,
			Scope: scope,
		}
		return Let{identifier}, nil
	}

	return Let{}, errors.New(fmt.Sprintf("Identifier %s is already declared", variable))
}

type Let struct {
	Identifier Identifier
}

func (l Let) Type() ObjectType {
	return LET
}

func (l Let) String() string {
	return fmt.Sprintf("let %s = %s", l.Identifier.Name, l.Value())
}

func (l Let) Value() Object {
	return l.Identifier.GetValue()
}
