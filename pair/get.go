package pair

func (self Pair[T1, T2]) Get() (T1, T2) {
	return self.Item1, self.Item2
}
