package object

type ObjectType string

const (
	Integer ObjectType = "Integer"
	Decimal ObjectType = "Decimal"
)

type Object interface {
	Type() ObjectType
}
