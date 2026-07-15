package scope

func (v Value[T]) Tap(act func(T)) Value[T] {
	act(v.value)
	return v
}
