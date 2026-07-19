package pipeline

import "fmt"

func (self Pipeline[T]) String() string {
	return fmt.Sprintf("(%v)", self.value)
}
