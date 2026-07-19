package pipeline

import "fmt"

func (v Pipeline[T]) GoString() string {
	return fmt.Sprintf("pipeline.Pipeline[%T](%#v)", v.value, v.value)
}
