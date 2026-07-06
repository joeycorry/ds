package scope

import "fmt"

func (v Value[T]) Also(act func(T)) T {
	return v.Tap(act).Get()
}

func (v Value[T]) Get() T {
	return v.value
}

func (v Value[T]) GoString() string {
	return fmt.Sprintf("scope.Value[%T](%#v)", v.value, v.value)
}

func (v Value[T]) Let[U any](transform func(T) U) U {
	return v.Map(transform).Get()
}

func (v Value[T]) Map[U any](transform func(T) U) Value[U] {
	return Value[U]{
		value: transform(v.value),
	}
}

func (v Value[T]) String() string {
	return fmt.Sprintf("(%v)", v.value)
}

func (v Value[T]) Tap(act func(T)) Value[T] {
	act(v.value)
	return v
}
