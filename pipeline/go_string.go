package pipeline

import "fmt"

func (self Pipeline[T]) GoString() string {
	return fmt.Sprintf("pipeline.Pipeline[%T](%#v)", self.value, self.value)
}
