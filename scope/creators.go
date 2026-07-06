package scope

func New[T any](value T) Value[T] {
	return Value[T]{value: value}
}
