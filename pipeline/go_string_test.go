package pipeline_test

import (
	"testing"

	"github.com/joeycorry/ds/internal/testingx"
	"github.com/joeycorry/ds/pipeline"
	"github.com/stretchr/testify/assert"
)

func TestGoString(t *testing.T) {
	t.Parallel()

	testingx.RunParallel(t, "returns a go-string representation of the pipeline", func(t *testing.T) {
		actual := pipeline.New(1).GoString()

		assert.Equal(t, "pipeline.Pipeline[int](1)", actual)
	})
}
