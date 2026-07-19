package unit

import "github.com/joeycorry/ds/pipeline"

func (self Unit) Pipeline() pipeline.Pipeline[Unit] {
	return pipeline.New(self)
}
