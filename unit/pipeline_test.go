package unit_test

import (
	"testing"

	"github.com/joeycorry/ds/internal/testingx"
	"github.com/joeycorry/ds/unit"
	"github.com/stretchr/testify/assert"
)

func TestPipeline(t *testing.T) {
	t.Parallel()

	testingx.RunParallel(t, "returns a pipeline containing the unit", func(t *testing.T) {
		input := unit.New()

		actual := input.Pipeline().Get()

		assert.Equal(t, input, actual)
	})
}
