package pair_test

import (
	"testing"

	"github.com/joeycorry/ds/internal/testingx"
	"github.com/joeycorry/ds/pair"
	"github.com/stretchr/testify/assert"
)

func TestGoString(t *testing.T) {
	t.Parallel()

	testingx.RunParallel(t, "returns a go-string representation of the tuple", func(t *testing.T) {
		actual := pair.New(1, "hello").GoString()

		assert.Equal(t, `pair.Pair[int, string](1, "hello")`, actual)
	})
}
