package eval

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/varshard/monkey/object"
	"github.com/varshard/monkey/parser"
	"testing"
)

func Test_Eval(t *testing.T) {
	t.Run("Test Eval boolean", func(t *testing.T) {
		testCases := []struct {
			input    string
			expected bool
		}{
			{"true;", true},
			{"false;", false},
			{"!true;", false},
			{"!false;", true},
		}

		for _, test := range testCases {
			obj, ok := evalCode(test.input).(object.BooleanObject)

			assert.True(t, ok)
			assert.Equal(t, test.expected, obj.Value)
		}
	})

	t.Run("Test Eval integer", func(t *testing.T) {
		testCases := []struct {
			input    string
			expected int
		}{
			{"2 + 5;", 7},
			{"2;", 2},
			{"2 - 3;", -1},
			{"2 * 3;", 6},
			//{ "4/2", 2},
		}

		for _, test := range testCases {
			obj, ok := evalCode(test.input).(object.IntegerObject)

			assert.True(t, ok)
			assert.Equal(t, test.expected, obj.Value)
		}
	})

	t.Run("Test Eval decimal", func(t *testing.T) {
		testCases := []struct {
			input    string
			expected float64
		}{
			{"2 + 5.5;", 7.5},
			{"2.5;", 2.5},
			{"2 - 0.5;", 1.5},
			{"2 * 3.0;", 6.0},
			//{ "5/2", 2.5},
		}

		for _, test := range testCases {
			obj, ok := evalCode(test.input).(object.DecimalObject)

			assert.True(t, ok)
			assert.Equal(t, test.expected, obj.Value)
		}
	})

	t.Run("Test Eval invalid infix expressions", func(t *testing.T) {
		testCases := []struct {
			input    string
			expected string
		}{
			{"2 + true;", fmt.Sprintf("Operation %s between %s and %s is undefined", "+", object.INTEGER, object.BOOLEAN)},
			{"2.0 + false;", fmt.Sprintf("Operation %s between %s and %s is undefined", "+", object.DECIMAL, object.BOOLEAN)},
			{"true - false;", fmt.Sprintf("Operation %s between %s and %s is undefined", "-", object.BOOLEAN, object.BOOLEAN)},
			{"true - 3;", fmt.Sprintf("Operation %s between %s and %s is undefined", "-", object.BOOLEAN, object.INTEGER)},
			{"3 - true;", fmt.Sprintf("Operation %s between %s and %s is undefined", "-", object.INTEGER, object.BOOLEAN)},
			{"2.5 + true;", fmt.Sprintf("Operation %s between %s and %s is undefined", "+", object.DECIMAL, object.BOOLEAN)},
			{"2.5 - !true;", fmt.Sprintf("Operation %s between %s and %s is undefined", "-", object.DECIMAL, object.BOOLEAN)},
		}

		for _, test := range testCases {
			obj, ok := evalCode(test.input).(object.Error)

			assert.True(t, ok)
			assert.Equal(t, test.expected, obj.String(), test.input)
		}
	})
}

func evalCode(code string) object.Object {
	p := parser.New(code)
	return Eval(p.ParseProgram())
}
