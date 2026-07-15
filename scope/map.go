package scope

func (v Value[T]) Map[U any](transform func(T) U) Value[U] {
	return Value[U]{
		value: transform(v.value),
	}
}
