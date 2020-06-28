package object

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_NewScope(t *testing.T) {
	scope := NewScope()

	scope.Set("x", IntegerObject{Value: 7})
	obj, ok := scope.Get("x").(IntegerObject)

	assert.True(t, ok)
	assert.Equal(t, 7, obj.Value)
}
