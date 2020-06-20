package eval

import (
	"errors"
	"fmt"
	"github.com/varshard/monkey/ast"
	"github.com/varshard/monkey/object"
	"github.com/varshard/monkey/token"
	"reflect"
)

type evalInfixFn func(left, right object.Object) object.Object
type rightSideInfixMap map[string]evalInfixFn

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

func makeError(err error) object.Error {
	return object.Error{Error: err}
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
	prefixFns := map[token.TokenType]map[string]rightSideInfixMap{
		token.Plus: {
			"IntegerObject": map[string]evalInfixFn{
				"IntegerObject": func(left, right object.Object) object.Object {
					leftObj := left.(object.IntegerObject)
					rightObj := right.(object.IntegerObject)

					return object.IntegerObject{Value: leftObj.Value + rightObj.Value}
				},
				"DecimalObject": func(left, right object.Object) object.Object {
					leftObj := left.(object.IntegerObject)
					rightObj := right.(object.DecimalObject)

					return object.DecimalObject{Value: float64(leftObj.Value) + rightObj.Value}
				},
			},
			"DecimalObject": map[string]evalInfixFn{
				"IntegerObject": func(left, right object.Object) object.Object {
					leftObj := left.(object.DecimalObject)
					rightObj := right.(object.IntegerObject)

					return object.DecimalObject{Value: leftObj.Value + float64(rightObj.Value)}
				},
				"DecimalObject": func(left, right object.Object) object.Object {
					leftObj := left.(object.DecimalObject)
					rightObj := right.(object.DecimalObject)

					return object.DecimalObject{Value: leftObj.Value + rightObj.Value}
				},
			},
		},
		token.Minus: {
			"IntegerObject": map[string]evalInfixFn{
				"IntegerObject": func(left, right object.Object) object.Object {
					leftObj := left.(object.IntegerObject)
					rightObj := right.(object.IntegerObject)

					return object.IntegerObject{Value: leftObj.Value - rightObj.Value}
				},
				"DecimalObject": func(left, right object.Object) object.Object {
					leftObj := left.(object.IntegerObject)
					rightObj := right.(object.DecimalObject)

					return object.DecimalObject{Value: float64(leftObj.Value) - rightObj.Value}
				},
			},
			"DecimalObject": map[string]evalInfixFn{
				"IntegerObject": func(left, right object.Object) object.Object {
					leftObj := left.(object.DecimalObject)
					rightObj := right.(object.IntegerObject)

					return object.DecimalObject{Value: leftObj.Value - float64(rightObj.Value)}
				},
				"DecimalObject": func(left, right object.Object) object.Object {
					leftObj := left.(object.DecimalObject)
					rightObj := right.(object.DecimalObject)

					return object.DecimalObject{Value: leftObj.Value - rightObj.Value}
				},
			},
		},
		token.Multiply: {
			"IntegerObject": map[string]evalInfixFn{
				"IntegerObject": func(left, right object.Object) object.Object {
					leftObj := left.(object.IntegerObject)
					rightObj := right.(object.IntegerObject)

					return object.IntegerObject{Value: leftObj.Value * rightObj.Value}
				},
				"DecimalObject": func(left, right object.Object) object.Object {
					leftObj := left.(object.IntegerObject)
					rightObj := right.(object.DecimalObject)

					return object.DecimalObject{Value: float64(leftObj.Value) * rightObj.Value}
				},
			},
			"DecimalObject": map[string]evalInfixFn{
				"IntegerObject": func(left, right object.Object) object.Object {
					leftObj := left.(object.DecimalObject)
					rightObj := right.(object.IntegerObject)

					return object.DecimalObject{Value: leftObj.Value * float64(rightObj.Value)}
				},
				"DecimalObject": func(left, right object.Object) object.Object {
					leftObj := left.(object.DecimalObject)
					rightObj := right.(object.DecimalObject)

					return object.DecimalObject{Value: leftObj.Value * rightObj.Value}
				},
			},
		},
	}

	left := evalExpression(node.Left)
	right := evalExpression(node.Right)

	leftType := reflect.TypeOf(left)
	rightType := reflect.TypeOf(right)
	prefixFn, ok := prefixFns[node.Token.Type][leftType.Name()][rightType.Name()]

	if !ok {
		return makeError(errors.New(fmt.Sprintf("Operation %s between %s and %s is undefined", node.Operator, left.Type(), right.Type())))
	}
	return prefixFn(left, right)
}

func evalInteger(intNode *ast.IntegerLiteral) object.IntegerObject {
	return object.IntegerObject{
		Value: intNode.Value,
	}
}

func evalDecimal(node *ast.DecimalLiteral) object.DecimalObject {
	return object.DecimalObject{Value: node.Value}
}
