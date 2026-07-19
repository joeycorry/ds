package pipeline

func (self Pipeline[T]) Tap(stage func(T)) Pipeline[T] {
	stage(self.value)
	return self
}
