package scope

import "fmt"

func (v Value[T]) GoString() string {
	return fmt.Sprintf("scope.Value[%T](%#v)", v.value, v.value)
}
