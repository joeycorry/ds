package pair_test

import (
	"testing"

	"github.com/joeycorry/ds/internal/testingx"
	"github.com/joeycorry/ds/pair"
	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	t.Parallel()

	testingx.RunParallel(t, "wraps the given items", func(t *testing.T) {
		item1, item2 := pair.New(1, "hello").Get()

		assert.Equal(t, 1, item1)
		assert.Equal(t, "hello", item2)
	})
}
