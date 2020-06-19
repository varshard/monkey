package object

type ObjectType string

const (
	Integer ObjectType = "Integer"
)

type Object interface {
	Type() ObjectType
}
