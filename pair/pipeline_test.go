package pair_test

import (
	"testing"

	"github.com/joeycorry/ds/internal/testingx"
	"github.com/joeycorry/ds/pair"
	"github.com/stretchr/testify/assert"
)

func TestPipeline(t *testing.T) {
	t.Parallel()

	testingx.RunParallel(t, "returns a pipeline containing the tuple", func(t *testing.T) {
		input := pair.New(1, "hello")

		actual := input.Pipeline().Get()

		assert.Equal(t, input, actual)
	})
}

func TestPipelineMut(t *testing.T) {
	t.Parallel()

	testingx.RunParallel(t, "returns a pipeline containing a pointer to the tuple", func(t *testing.T) {
		input := pair.New(1, "hello")

		actual := input.PipelineMut().Get()

		assert.Equal(t, &input, actual)
	})

	testingx.RunParallel(t, "propagates changes made to the tuple through the pipeline", func(t *testing.T) {
		input := pair.New(1, "hello")

		input.PipelineMut().Tap(func(value *pair.Pair[int, string]) {
			value.Item1 = 2
		})

		assert.Equal(t, 2, input.Item1)
	})
}
