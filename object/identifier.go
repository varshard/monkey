package object

import (
	"fmt"
)

type Identifier struct {
	Name  string
	Scope *Scope
}

func (i Identifier) Type() ObjectType {
	return IDENTIFIER
}

func (i Identifier) String() string {
	return fmt.Sprintf("%s = %s", i.Name, i.GetValue().String())
}

func (i Identifier) GetValue() Object {
	return i.Scope.Get(i.Name)
}

func (i Identifier) SetValue(value Object) {
	i.Scope.Set(i.Name, value)
}
