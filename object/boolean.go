package object

import "strconv"

type BooleanObject struct {
	Value bool
}

func (b BooleanObject) Type() ObjectType {
	return BOOLEAN
}

func (b BooleanObject) String() string {
	return strconv.FormatBool(b.Value)
}
