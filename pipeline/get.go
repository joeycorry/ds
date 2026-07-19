package pipeline

func (v Pipeline[T]) Get() T {
	return v.value
}
