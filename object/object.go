package object

type ObjectType string

const (
	Integer    ObjectType = "Integer"
	Decimal    ObjectType = "Decimal"
	Expression ObjectType = "ExpressionObject"
	Identifier ObjectType = "Identifier"
)

type Object interface {
	Type() ObjectType
	String() string
}
