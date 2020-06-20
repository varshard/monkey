package object

type ExpressionObject struct {
	Expression Object
}

func (e ExpressionObject) Type() ObjectType {
	return Expression
}
