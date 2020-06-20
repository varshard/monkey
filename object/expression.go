package object

type ExpressionObject struct {
	Expression Object
}

func (e ExpressionObject) String() string {
	return e.Expression.String()
}

func (e ExpressionObject) Type() ObjectType {
	return EXPRESSION
}
