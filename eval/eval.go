package eval

import (
	"github.com/varshard/monkey/ast"
	"github.com/varshard/monkey/object"
	"github.com/varshard/monkey/token"
	"reflect"
)

func Eval(node ast.Node) object.Object {
	switch node := node.(type) {
	case *ast.Program:
		return evalProgram(node)
	case ast.Statement:
		return evalStatement(node)
	default:
		return nil
	}
}

func evalStatement(node ast.Statement) object.Object {
	switch node := node.(type) {
	case *ast.ExpressionStatement:
		return evalExpression(node.Expression)
	default:
		return nil
	}
}

func evalExpression(node ast.Expression) object.Object {
	switch node := node.(type) {
	case *ast.IntegerLiteral:
		return evalInteger(node)
	case *ast.DecimalLiteral:
		return evalDecimal(node)
	case *ast.InfixExpression:
		return evalInfixExpression(node)
	default:
		return nil
	}
}

func evalProgram(node *ast.Program) object.Object {
	var result object.Object
	for _, statement := range node.Statements {
		result = Eval(statement)
	}

	return result
}

func evalInfixExpression(node *ast.InfixExpression) object.Object {
	left := evalExpression(node.Left)
	right := evalExpression(node.Right)

	switch node.Token.Type {
	case token.Plus:
		// TODO: support decimal
		leftType := reflect.TypeOf(left)
		rightType := reflect.TypeOf(right)

		if leftType.ConvertibleTo(reflect.TypeOf(object.DecimalObject{})) {
			leftObj := left.(object.DecimalObject)
			if rightType.ConvertibleTo(reflect.TypeOf(object.DecimalObject{})) {
				rightObj := right.(object.DecimalObject)
				return object.DecimalObject{Value: leftObj.Value + rightObj.Value}
			} else if rightType.ConvertibleTo(reflect.TypeOf(object.IntegerObject{})) {
				rightObj := right.(object.IntegerObject)
				return object.DecimalObject{Value: leftObj.Value + float64(rightObj.Value)}
			} else {
				// TODO: handle an unsupported type
				return nil
			}
		} else if leftType.ConvertibleTo(reflect.TypeOf(object.IntegerObject{})) {
			leftObj := left.(object.IntegerObject)
			if rightType.ConvertibleTo(reflect.TypeOf(object.DecimalObject{})) {
				rightObj := right.(object.DecimalObject)
				return object.DecimalObject{Value: float64(leftObj.Value) + rightObj.Value}
			} else if rightType.ConvertibleTo(reflect.TypeOf(object.IntegerObject{})) {
				rightObj := right.(object.IntegerObject)
				return object.IntegerObject{Value: leftObj.Value + rightObj.Value}
			} else {
				// TODO: handle an unsupported type
				return nil
			}
		}
		// TODO: handle an unsupported combination
	}
	return nil
}

func evalInteger(intNode *ast.IntegerLiteral) object.IntegerObject {
	return object.IntegerObject{
		Value: intNode.Value,
	}
}

func evalDecimal(node *ast.DecimalLiteral) object.DecimalObject {
	return object.DecimalObject{Value: node.Value}
}
