package pair

import "github.com/joeycorry/ds/pipeline"

func (self Pair[T1, T2]) Pipeline() pipeline.Pipeline[Pair[T1, T2]] {
	return pipeline.New(self)
}

func (self *Pair[T1, T2]) PipelineMut() pipeline.Pipeline[*Pair[T1, T2]] {
	return pipeline.New(self)
}
