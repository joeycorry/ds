package clone

type Cloner[T any] interface {
	Clone() T
}
