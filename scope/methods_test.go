package scope_test

import (
	"testing"

	"github.com/joeycorry/ds/internal/testingutil"
	"github.com/joeycorry/ds/scope"
	"github.com/stretchr/testify/assert"
)

func TestAlso(t *testing.T) {
	t.Parallel()

	testingutil.RunTParallel(t, "executes the action with the contained value", func(t *testing.T) {
		input := 1
		actCalled := false
		act := func(value int) {
			actCalled = true
		}

		scope.New(input).Also(act)

		assert.True(t, actCalled)
	})

	testingutil.RunTParallel(t, "returns the contained value from the original scope after the action", func(t *testing.T) {
		input := new(1)

		actual := scope.New(input).Also(func(value *int) {})

		assert.Equal(t, input, actual)
	})

	testingutil.RunTParallel(t, "propagates changes made to contained pointers in the action", func(t *testing.T) {
		input := 1

		actual := scope.New(&input).Also(func(value *int) {
			*value += 1
		})

		assert.Equal(t, 2, *actual)
	})

	testingutil.RunTParallel(t, "does not propagate changes made to contained non-pointers in the action", func(t *testing.T) {
		input := 1

		actual := scope.New(input).Also(func(value int) {
			value += 1
		})

		assert.Equal(t, 1, actual)
	})
}

func TestGet(t *testing.T) {
	t.Parallel()

	testingutil.RunTParallel(t, "returns the contained value", func(t *testing.T) {
		input := "hello"

		actual := scope.New(input).Get()

		assert.Equal(t, input, actual)
	})
}

func TestGoString(t *testing.T) {
	t.Parallel()

	testingutil.RunTParallel(t, "returns a go-string representation of the scope", func(t *testing.T) {
		actual := scope.New(1).GoString()

		assert.Equal(t, "scope.Value[int](1)", actual)
	})
}

func TestLet(t *testing.T) {
	t.Parallel()

	testingutil.RunTParallel(t, "does not modify the original scope", func(t *testing.T) {
		inputScope := scope.New(1)

		inputScope.Let(func(value int) int {
			return value + 1
		})

		assert.Equal(t, 1, inputScope.Get())
	})

	testingutil.RunTParallel(t, "returns the output of the transformation", func(t *testing.T) {
		input := 1

		actual := scope.New(input).Let(func(value int) int {
			return value + 1
		})

		assert.Equal(t, 2, actual)
	})

	testingutil.RunTParallel(t, "propagates changes made to contained pointers in the transformation", func(t *testing.T) {
		input := 1

		scope.New(&input).Let(func(value *int) int {
			*value += 1
			return *value
		})

		assert.Equal(t, 2, input)
	})

	testingutil.RunTParallel(t, "does not propagate changes made to contained non-pointers in the transformation", func(t *testing.T) {
		input := 1

		scope.New(input).Let(func(value int) int {
			value += 1
			return value
		})

		assert.Equal(t, 1, input)
	})
}

func TestMap(t *testing.T) {
	t.Parallel()

	testingutil.RunTParallel(t, "does not modify the original scope", func(t *testing.T) {
		inputScope := scope.New(1)

		inputScope.Map(func(value int) int {
			return value + 1
		})

		assert.Equal(t, 1, inputScope.Get())
	})

	testingutil.RunTParallel(t, "returns a scope that wraps the output of the transformation", func(t *testing.T) {
		input := 1

		actual := scope.New(input).Map(func(value int) int {
			return value + 1
		})

		assert.Equal(t, 2, actual.Get())
	})

	testingutil.RunTParallel(t, "propagates changes made to contained pointers in the transformation", func(t *testing.T) {
		input := 1

		scope.New(&input).Map(func(value *int) int {
			*value += 1
			return *value
		})

		assert.Equal(t, 2, input)
	})

	testingutil.RunTParallel(t, "does not propagate changes made to contained non-pointers in the transformation", func(t *testing.T) {
		input := 1

		scope.New(input).Map(func(value int) int {
			value += 1
			return value
		})

		assert.Equal(t, 1, input)
	})
}

func TestString(t *testing.T) {
	t.Parallel()

	testingutil.RunTParallel(t, "returns a string representation of the scope", func(t *testing.T) {
		actual := scope.New(1).String()

		assert.Equal(t, "(1)", actual)
	})
}

func TestTap(t *testing.T) {
	t.Parallel()

	testingutil.RunTParallel(t, "executes the action with the contained value", func(t *testing.T) {
		input := 1
		actCalled := false
		act := func(value int) {
			actCalled = true
		}

		scope.New(input).Tap(act)

		assert.True(t, actCalled)
	})

	testingutil.RunTParallel(t, "returns a scope containing the original scope's value after the action", func(t *testing.T) {
		input := new(1)
		inputScope := scope.New(input)

		actual := inputScope.Tap(func(value *int) {})

		assert.Equal(t, input, actual.Get())
	})

	testingutil.RunTParallel(t, "propagates changes made to contained pointers in the action", func(t *testing.T) {
		input := 1

		actual := scope.New(&input).Tap(func(value *int) {
			*value += 1
		})

		assert.Equal(t, 2, *actual.Get())
	})

	testingutil.RunTParallel(t, "does not propagate changes made to contained non-pointers in the action", func(t *testing.T) {
		input := 1

		actual := scope.New(input).Tap(func(value int) {
			value += 1
		})

		assert.Equal(t, 1, actual.Get())
	})
}
