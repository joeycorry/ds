package pair

import "fmt"

func (self Pair[T1, T2]) String() string {
	return fmt.Sprintf("(%v, %v)", self.Item1, self.Item2)
}
