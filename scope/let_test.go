package scope_test

import (
	"testing"

	"github.com/joeycorry/ds/internal/testingx"
	"github.com/joeycorry/ds/scope"
	"github.com/stretchr/testify/assert"
)

func TestLet(t *testing.T) {
	t.Parallel()

	testingx.RunParallel(t, "does not modify the original scope", func(t *testing.T) {
		inputScope := scope.New(1)

		inputScope.Let(func(value int) int {
			return value + 1
		})

		assert.Equal(t, 1, inputScope.Get())
	})

	testingx.RunParallel(t, "returns the output of the transformation", func(t *testing.T) {
		input := 1

		actual := scope.New(input).Let(func(value int) int {
			return value + 1
		})

		assert.Equal(t, 2, actual)
	})

	testingx.RunParallel(t, "propagates changes made to contained pointers in the transformation", func(t *testing.T) {
		input := 1

		scope.New(&input).Let(func(value *int) int {
			*value += 1
			return *value
		})

		assert.Equal(t, 2, input)
	})

	testingx.RunParallel(t, "does not propagate changes made to contained non-pointers in the transformation", func(t *testing.T) {
		input := 1

		scope.New(input).Let(func(value int) int {
			value += 1
			return value
		})

		assert.Equal(t, 1, input)
	})
}
