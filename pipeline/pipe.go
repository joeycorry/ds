package pipeline

func (self Pipeline[T]) Pipe[U any](stage func(T) U) Pipeline[U] {
	return New(stage(self.value))
}
