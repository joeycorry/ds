package pair

import "github.com/joeycorry/ds/internal/clone"

func Clone[T1 clone.Cloner[T1], T2 clone.Cloner[T2]](self Pair[T1, T2]) Pair[T1, T2] {
	return New(self.Item1.Clone(), self.Item2.Clone())
}
