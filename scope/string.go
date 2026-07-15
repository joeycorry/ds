package scope

import "fmt"

func (v Value[T]) String() string {
	return fmt.Sprintf("(%v)", v.value)
}
