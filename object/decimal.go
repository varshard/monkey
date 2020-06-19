package object

type DecimalObject struct {
	Value float64
}

func (d DecimalObject) Type() ObjectType {
	return Decimal
}
