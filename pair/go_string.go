package pair

import "fmt"

func (self Pair[T1, T2]) GoString() string {
	return fmt.Sprintf("pair.Pair[%T, %T](%#v, %#v)", self.Item1, self.Item2, self.Item1, self.Item2)
}
