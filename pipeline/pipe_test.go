package pipeline_test

import (
	"testing"

	"github.com/joeycorry/ds/internal/testingx"
	"github.com/joeycorry/ds/pipeline"
	"github.com/stretchr/testify/assert"
)

func TestLet(t *testing.T) {
	t.Parallel()

	testingx.RunParallel(t, "does not modify the original pipeline", func(t *testing.T) {
		inputScope := pipeline.New(1)

		inputScope.Pipe(func(value int) int {
			return value + 1
		})

		assert.Equal(t, 1, inputScope.Get())
	})

	testingx.RunParallel(t, "returns the a new pipeline with the output of the transformation", func(t *testing.T) {
		input := 1

		actual := pipeline.New(input).Pipe(func(value int) int {
			return value + 1
		})

		assert.Equal(t, 2, actual.Get())
	})

	testingx.RunParallel(t, "propagates changes made to contained pointers in the transformation", func(t *testing.T) {
		input := 1

		pipeline.New(&input).Pipe(func(value *int) int {
			*value += 1
			return *value
		})

		assert.Equal(t, 2, input)
	})

	testingx.RunParallel(t, "does not propagate changes made to contained non-pointers in the transformation", func(t *testing.T) {
		input := 1

		pipeline.New(input).Pipe(func(value int) int {
			value += 1
			return value
		})

		assert.Equal(t, 1, input)
	})
}
