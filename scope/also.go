package scope

func (v Value[T]) Also(act func(T)) T {
	return v.Tap(act).Get()
}
