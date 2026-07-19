package unit_test

import (
	"testing"

	"github.com/joeycorry/ds/internal/testingx"
	"github.com/joeycorry/ds/unit"
	"github.com/stretchr/testify/assert"
)

func TestString(t *testing.T) {
	t.Parallel()

	testingx.RunParallel(t, "returns a string representation of the unit", func(t *testing.T) {
		actual := unit.New().String()

		assert.Equal(t, "()", actual)
	})
}
