package scope_test

import (
	"testing"

	"github.com/joeycorry/ds/internal/testingx"
	"github.com/joeycorry/ds/scope"
	"github.com/stretchr/testify/assert"
)

func TestAlso(t *testing.T) {
	t.Parallel()

	testingx.RunParallel(t, "executes the action with the contained value", func(t *testing.T) {
		input := 1
		actCalled := false
		act := func(value int) {
			actCalled = true
		}

		scope.New(input).Also(act)

		assert.True(t, actCalled)
	})

	testingx.RunParallel(t, "returns the contained value from the original scope after the action", func(t *testing.T) {
		input := new(1)

		actual := scope.New(input).Also(func(value *int) {})

		assert.Equal(t, input, actual)
	})

	testingx.RunParallel(t, "propagates changes made to contained pointers in the action", func(t *testing.T) {
		input := 1

		actual := scope.New(&input).Also(func(value *int) {
			*value += 1
		})

		assert.Equal(t, 2, *actual)
	})

	testingx.RunParallel(t, "does not propagate changes made to contained non-pointers in the action", func(t *testing.T) {
		input := 1

		actual := scope.New(input).Also(func(value int) {
			value += 1
		})

		assert.Equal(t, 1, actual)
	})
}
