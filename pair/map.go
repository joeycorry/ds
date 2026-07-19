package pair

func (self Pair[T1, T2]) Map[U1 any, U2 any](transform func(T1, T2) (U1, U2)) Pair[U1, U2] {
	return New(transform(self.Item1, self.Item2))
}
