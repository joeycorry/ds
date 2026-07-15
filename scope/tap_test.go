package scope_test

import (
	"testing"

	"github.com/joeycorry/ds/internal/testingutil"
	"github.com/joeycorry/ds/scope"
	"github.com/stretchr/testify/assert"
)

func TestTap(t *testing.T) {
	t.Parallel()

	testingutil.Extensions.T.RunParallel(t, "executes the action with the contained value", func(t *testing.T) {
		input := 1
		actCalled := false
		act := func(value int) {
			actCalled = true
		}

		scope.New(input).Tap(act)

		assert.True(t, actCalled)
	})

	testingutil.Extensions.T.RunParallel(t, "returns a scope containing the original scope's value after the action", func(t *testing.T) {
		input := new(1)
		inputScope := scope.New(input)

		actual := inputScope.Tap(func(value *int) {})

		assert.Equal(t, input, actual.Get())
	})

	testingutil.Extensions.T.RunParallel(t, "propagates changes made to contained pointers in the action", func(t *testing.T) {
		input := 1

		actual := scope.New(&input).Tap(func(value *int) {
			*value += 1
		})

		assert.Equal(t, 2, *actual.Get())
	})

	testingutil.Extensions.T.RunParallel(t, "does not propagate changes made to contained non-pointers in the action", func(t *testing.T) {
		input := 1

		actual := scope.New(input).Tap(func(value int) {
			value += 1
		})

		assert.Equal(t, 1, actual.Get())
	})
}
