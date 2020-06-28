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

func ExecuteProgram(node *ast.Program) object.Object {
	scope := object.NewScope()
	var result object.Object
	for _, statement := range node.Statements {
		result = Eval(statement, &scope)
	}

	return result
}

func Eval(node ast.Node, scope *object.Scope) object.Object {
	switch node := node.(type) {
	case ast.Statement:
		return evalStatement(node, scope)
	default:
		return makeError(errors.New("statement expected"))
	}
}

func evalStatement(node ast.Statement, scope *object.Scope) object.Object {
	switch node := node.(type) {
	case *ast.LetStatement:
		return evalLet(node, scope)
	case *ast.ExpressionStatement:
		return evalExpression(node.Expression, scope)
	default:
		return nil
	}
}

func makeError(err error) object.Error {
	return object.Error{Error: err}
}

func evalExpression(node ast.Expression, scope *object.Scope) object.Object {
	switch node := node.(type) {
	case *ast.Boolean:
		return evalBoolean(node)
	case *ast.IntegerLiteral:
		return evalInteger(node)
	case *ast.DecimalLiteral:
		return evalDecimal(node)
	case *ast.Identifier:
		return evalIdentifier(node, scope)
	case *ast.PrefixExpression:
		return evalPrefix(node, scope)
	case *ast.InfixExpression:
		return evalInfix(node, scope)
	default:
		return object.Null{}
	}
}

func evalIdentifier(node *ast.Identifier, scope *object.Scope) object.Object {
	name := node.Name
	if scope.IsDeclared(name) {
		return scope.Get(name)
	}

	return makeError(errors.New(fmt.Sprintf("%s hasn't been declared", name)))
}

func evalInfix(node *ast.InfixExpression, scope *object.Scope) object.Object {
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

	left := evalExpression(node.Left, scope)
	right := evalExpression(node.Right, scope)

	leftType := reflect.TypeOf(left)
	rightType := reflect.TypeOf(right)
	prefixFn, ok := prefixFns[node.Token.Type][leftType.Name()][rightType.Name()]

	if !ok {
		return makeError(errors.New(fmt.Sprintf("Operation %s between %s and %s is undefined", node.Operator, left.Type(), right.Type())))
	}
	return prefixFn(left, right)
}

func evalPrefix(node *ast.PrefixExpression, scope *object.Scope) object.Object {
	switch node.Token.Type {
	case token.Bang:
		obj, isBool := evalExpression(node.Right, scope).(object.BooleanObject)

		if !isBool {
			return makeError(errors.New(fmt.Sprintf("! of %s doesn't exist", obj.String())))
		}
		return object.BooleanObject{Value: !obj.Value}
	case token.Minus:
		right := evalExpression(node.Right, scope)
		switch right.Type() {
		case object.INTEGER:
			obj := right.(object.IntegerObject)
			obj.Value = -obj.Value
			return obj
		case object.DECIMAL:
			obj := right.(object.DecimalObject)
			obj.Value = -obj.Value
			return obj
		default:
			return makeError(errors.New(fmt.Sprintf("expected an integer, but got %s", right.Type())))
		}
	default:
		return nil
	}
}

func evalInteger(intNode *ast.IntegerLiteral) object.IntegerObject {
	return object.IntegerObject{
		Value: intNode.Value,
	}
}

func evalDecimal(node *ast.DecimalLiteral) object.DecimalObject {
	return object.DecimalObject{Value: node.Value}
}

func evalBoolean(node *ast.Boolean) object.BooleanObject {
	return object.BooleanObject{Value: node.Value}
}

func evalLet(node *ast.LetStatement, scope *object.Scope) object.Object {
	value := evalExpression(node.Value, scope)
	let, err := object.NewLet(node.Variable.Name, scope, value)

	if err != nil {
		return makeError(err)
	}

	return let
}
