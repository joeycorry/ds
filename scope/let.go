package scope

func (v Value[T]) Let[U any](transform func(T) U) U {
	return v.Map(transform).Get()
}
