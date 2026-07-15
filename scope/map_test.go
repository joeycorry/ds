package scope_test

import (
	"testing"

	"github.com/joeycorry/ds/internal/testingutil"
	"github.com/joeycorry/ds/scope"
	"github.com/stretchr/testify/assert"
)

func TestMap(t *testing.T) {
	t.Parallel()

	testingutil.Extensions.T.RunParallel(t, "does not modify the original scope", func(t *testing.T) {
		inputScope := scope.New(1)

		inputScope.Map(func(value int) int {
			return value + 1
		})

		assert.Equal(t, 1, inputScope.Get())
	})

	testingutil.Extensions.T.RunParallel(t, "returns a scope that wraps the output of the transformation", func(t *testing.T) {
		input := 1

		actual := scope.New(input).Map(func(value int) int {
			return value + 1
		})

		assert.Equal(t, 2, actual.Get())
	})

	testingutil.Extensions.T.RunParallel(t, "propagates changes made to contained pointers in the transformation", func(t *testing.T) {
		input := 1

		scope.New(&input).Map(func(value *int) int {
			*value += 1
			return *value
		})

		assert.Equal(t, 2, input)
	})

	testingutil.Extensions.T.RunParallel(t, "does not propagate changes made to contained non-pointers in the transformation", func(t *testing.T) {
		input := 1

		scope.New(input).Map(func(value int) int {
			value += 1
			return value
		})

		assert.Equal(t, 1, input)
	})
}
