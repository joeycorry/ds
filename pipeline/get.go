package pipeline

func (self Pipeline[T]) Get() T {
	return self.value
}
