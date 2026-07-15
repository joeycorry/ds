package testingx

import "testing"

func RunParallel(self *testing.T, name string, f func(t *testing.T)) bool {
	return self.Run(name, func(t *testing.T) {
		t.Parallel()
		f(t)
	})
}
