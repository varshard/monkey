package object

type Null struct {
}

func (n Null) Type() ObjectType {
	return NULL
}

func (n Null) String() string {
	return "null"
}
