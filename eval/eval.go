package eval

import (
	"github.com/varshard/monkey/ast"
	"github.com/varshard/monkey/object"
	"github.com/varshard/monkey/token"
)

func Eval(node ast.Node) object.Object {
	switch node := node.(type) {
	case ast.IntegerLiteral:
		return evalInteger(node)
	case ast.DecimalLiteral:
		return evalDecimal(node)
	case ast.InfixExpression:
		return evalInfixExpression(node)
	}

	return nil
}

func evalInfixExpression(node ast.InfixExpression) object.Object {
	left := Eval(node.Left)
	right := Eval(node.Right)

	switch node.Token.Type {
	case token.Plus:
		// TODO: support decimal
		//leftType := reflect.TypeOf(left)
		//rightType := reflect.TypeOf(right)
		//
		//isFloat := false
		//var leftObj object.Object
		//var rightObj object.Object
		//if leftType.ConvertibleTo(reflect.TypeOf(object.DecimalObject{})) {
		//	isFloat = true
		//	leftObj = left.(object.DecimalObject)
		//} else {
		//	leftObj = left.(object.IntegerObject)
		//}
		//if rightType.ConvertibleTo(reflect.TypeOf(object.DecimalObject{})) {
		//	isFloat = true
		//	rightObj = right.(object.DecimalObject)
		//} else {
		//	rightObj = right.(object.IntegerObject)
		//}
		//
		//if isFloat {
		//	return object.DecimalObject{
		//		Value: leftObj.Value + rightObj.Value,
		//	}
		//}
		return object.IntegerObject{
			Value: left.(object.IntegerObject).Value + right.(object.IntegerObject).Value,
		}
	}
	return nil
}

func evalInteger(intNode ast.IntegerLiteral) object.IntegerObject {
	return object.IntegerObject{
		Value: intNode.Value,
	}
}

func evalDecimal(node ast.DecimalLiteral) object.DecimalObject {
	return object.DecimalObject{Value: node.Value}
}
