package object

type ObjectType string

const (
	BOOLEAN    ObjectType = "BOOLEAN"
	INTEGER    ObjectType = "INTEGER"
	DECIMAL    ObjectType = "DECIMAL"
	EXPRESSION ObjectType = "EXPRESSION"
	IDENTIFIER ObjectType = "IDENTIFIER"
	ERROR      ObjectType = "ERROR"
)

type Object interface {
	Type() ObjectType
	String() string
}
