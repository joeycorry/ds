package pipeline

func New[T any](value T) Pipeline[T] {
	return Pipeline[T]{
		value: value,
	}
}
