package pipeline

import "fmt"

func (v Pipeline[T]) String() string {
	return fmt.Sprintf("(%v)", v.value)
}
