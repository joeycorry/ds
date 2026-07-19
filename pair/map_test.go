package pair_test

import (
	"strconv"
	"testing"

	"github.com/joeycorry/ds/internal/testingx"
	"github.com/joeycorry/ds/pair"
	"github.com/stretchr/testify/assert"
)

func TestMap(t *testing.T) {
	t.Parallel()

	testingx.RunParallel(t, "returns a new tuple with the transformed items", func(t *testing.T) {
		input := pair.New(1, 2)

		actual := input.Map(func(item1, item2 int) (string, string) {
			return strconv.Itoa(item1), strconv.Itoa(item2)
		})

		assert.Equal(t, pair.New("1", "2"), actual)
	})

	testingx.RunParallel(t, "does not modify the original tuple", func(t *testing.T) {
		input := pair.New(1, 2)

		input.Map(func(item1, item2 int) (int, int) {
			return item1 + 1, item2 + 1
		})

		assert.Equal(t, pair.New(1, 2), input)
	})
}
