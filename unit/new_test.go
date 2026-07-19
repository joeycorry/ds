package unit_test

import (
	"testing"

	"github.com/joeycorry/ds/internal/testingx"
	"github.com/joeycorry/ds/unit"
	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	t.Parallel()

	testingx.RunParallel(t, "returns a unit value", func(t *testing.T) {
		actual := unit.New()

		assert.Equal(t, unit.Unit{}, actual)
	})
}
