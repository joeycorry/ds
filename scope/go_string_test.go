package scope_test

import (
	"testing"

	"github.com/joeycorry/ds/internal/testingx"
	"github.com/joeycorry/ds/scope"
	"github.com/stretchr/testify/assert"
)

func TestGoString(t *testing.T) {
	t.Parallel()

	testingx.RunParallel(t, "returns a go-string representation of the scope", func(t *testing.T) {
		actual := scope.New(1).GoString()

		assert.Equal(t, "scope.Value[int](1)", actual)
	})
}
