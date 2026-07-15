package scope_test

import (
	"testing"

	"github.com/joeycorry/ds/internal/testingutil"
	"github.com/joeycorry/ds/scope"
	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	t.Parallel()

	testingutil.Extensions.T.RunParallel(t, "wraps the given value", func(t *testing.T) {
		input := "hello"

		actual := scope.New(input).Get()

		assert.Equal(t, input, actual)
	})
}
