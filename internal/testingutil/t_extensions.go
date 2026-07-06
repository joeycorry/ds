package testingutil

import "testing"

func RunTParallel(t *testing.T, name string, f func(t *testing.T)) bool {
	return t.Run(name, func(t *testing.T) {
		t.Parallel()
		f(t)
	})
}
