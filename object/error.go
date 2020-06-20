package object

type Error struct {
	Error error
}

func (e Error) Type() ObjectType {
	return ERROR
}

func (e Error) String() string {
	return e.Error.Error()
}
