package pipeline

func (self Pipeline[T]) Tap(act func(T)) Pipeline[T] {
	act(self.value)
	return self
}
