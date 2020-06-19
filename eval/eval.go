package eval

import (
	"github.com/varshard/monkey/ast"
	"github.com/varshard/monkey/object"
)

func Eval(node ast.Node) object.Object {
	switch node := node.(type) {
	case ast.IntegerLiteral:
		return evalInteger(node)
	}

	return nil
}

func evalInteger(intNode ast.IntegerLiteral) object.IntegerObject {
	return object.IntegerObject{
		Value: intNode.Value,
	}
}
