package eval

import (
	"github.com/stretchr/testify/assert"
	"github.com/varshard/monkey/object"
	"github.com/varshard/monkey/parser"
	"testing"
)

func Test_Eval(t *testing.T) {
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
}

func evalCode(code string) object.Object {
	p := parser.New(code)
	return Eval(p.ParseProgram())
}
