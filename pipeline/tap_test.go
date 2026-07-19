package pipeline_test

import (
	"testing"

	"github.com/joeycorry/ds/internal/testingx"
	"github.com/joeycorry/ds/pipeline"
	"github.com/stretchr/testify/assert"
)

func TestTap(t *testing.T) {
	t.Parallel()

	testingx.RunParallel(t, "executes the action with the contained value", func(t *testing.T) {
		input := 1
		actCalled := false
		act := func(value int) {
			actCalled = true
		}

		pipeline.New(input).Tap(act)

		assert.True(t, actCalled)
	})

	testingx.RunParallel(t, "returns a pipeline containing the original pipeline's value after the action", func(t *testing.T) {
		input := new(1)
		inputScope := pipeline.New(input)

		actual := inputScope.Tap(func(value *int) {})

		assert.Equal(t, input, actual.Get())
	})

	testingx.RunParallel(t, "propagates changes made to contained pointers in the action", func(t *testing.T) {
		input := 1

		actual := pipeline.New(&input).Tap(func(value *int) {
			*value += 1
		})

		assert.Equal(t, 2, *actual.Get())
	})

	testingx.RunParallel(t, "does not propagate changes made to contained non-pointers in the action", func(t *testing.T) {
		input := 1

		actual := pipeline.New(input).Tap(func(value int) {
			value += 1
		})

		assert.Equal(t, 1, actual.Get())
	})
}
