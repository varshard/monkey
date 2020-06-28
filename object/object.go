package object

type ObjectType string

const (
	BOOLEAN    ObjectType = "BOOLEAN"
	INTEGER    ObjectType = "INTEGER"
	DECIMAL    ObjectType = "DECIMAL"
	EXPRESSION ObjectType = "EXPRESSION"
	IDENTIFIER ObjectType = "IDENTIFIER"
	ERROR      ObjectType = "ERROR"
	NULL       ObjectType = "NULL"
	LET        ObjectType = "LET"
)

type Object interface {
	Type() ObjectType
	String() string
}
