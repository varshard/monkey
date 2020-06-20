package object

import "strconv"

type DecimalObject struct {
	Value float64
}

func (d DecimalObject) Type() ObjectType {
	return DECIMAL
}

func (d DecimalObject) String() string {
	return strconv.FormatFloat(d.Value, 'e', -1, 64)
}
