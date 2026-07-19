package pipeline_test

import (
	"testing"

	"github.com/joeycorry/ds/internal/testingx"
	"github.com/joeycorry/ds/pipeline"
	"github.com/stretchr/testify/assert"
)

func TestGet(t *testing.T) {
	t.Parallel()

	testingx.RunParallel(t, "returns the contained value", func(t *testing.T) {
		input := "hello"

		actual := pipeline.New(input).Get()

		assert.Equal(t, input, actual)
	})
}
