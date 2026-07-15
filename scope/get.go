package scope

func (v Value[T]) Get() T {
	return v.value
}
