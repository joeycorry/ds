package pipeline_test

import (
	"testing"

	"github.com/joeycorry/ds/internal/testingx"
	"github.com/joeycorry/ds/pipeline"
	"github.com/stretchr/testify/assert"
)

func TestString(t *testing.T) {
	t.Parallel()

	testingx.RunParallel(t, "returns a string representation of the pipeline", func(t *testing.T) {
		actual := pipeline.New(1).String()

		assert.Equal(t, "(1)", actual)
	})
}
