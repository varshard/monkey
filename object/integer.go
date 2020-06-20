package object

import "strconv"

type IntegerObject struct {
	Value int
}

func (i IntegerObject) String() string {
	return strconv.Itoa(i.Value)
}

func (i IntegerObject) Type() ObjectType {
	return INTEGER
}
