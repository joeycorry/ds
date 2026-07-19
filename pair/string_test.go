package pair_test

import (
	"testing"

	"github.com/joeycorry/ds/internal/testingx"
	"github.com/joeycorry/ds/pair"
	"github.com/stretchr/testify/assert"
)

func TestString(t *testing.T) {
	t.Parallel()

	testingx.RunParallel(t, "returns a string representation of the tuple", func(t *testing.T) {
		actual := pair.New(1, "hello").String()

		assert.Equal(t, "(1, hello)", actual)
	})
}
