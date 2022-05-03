package golang_generics

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func IsSameValue[T comparable](value1 T, value2 T) bool {
	if value1 == value2 {
		return true
	} else {
		return false
	}
}

func TestIsSameValue(t *testing.T) {
	assert.Equal(t, true, IsSameValue[string]("ring", "ring"))
	assert.Equal(t, true, IsSameValue[int](26, 26))
	assert.Equal(t, true, IsSameValue[bool](false, false))
}
