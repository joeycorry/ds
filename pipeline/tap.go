package pipeline

func (v Pipeline[T]) Tap(act func(T)) Pipeline[T] {
	act(v.value)
	return v
}
