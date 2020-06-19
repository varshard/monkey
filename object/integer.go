package object

type IntegerObject struct {
	Value int
}

func (i IntegerObject) Type() ObjectType {
	return Integer
}
