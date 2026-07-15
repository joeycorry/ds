package scope_test

import (
	"testing"

	"github.com/joeycorry/ds/internal/testingutil"
	"github.com/joeycorry/ds/scope"
	"github.com/stretchr/testify/assert"
)

func TestString(t *testing.T) {
	t.Parallel()

	testingutil.Extensions.T.RunParallel(t, "returns a string representation of the scope", func(t *testing.T) {
		actual := scope.New(1).String()

		assert.Equal(t, "(1)", actual)
	})
}
