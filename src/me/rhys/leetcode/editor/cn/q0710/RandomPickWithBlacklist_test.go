package q0710

import (
	"fmt"
	"testing"
)

func TestSolution_Pick(t *testing.T) {
	solution := Constructor(7, []int{2, 3, 5})
	fmt.Printf("%#v\n", solution.Pick())
	fmt.Printf("%#v\n", solution.Pick())
	fmt.Printf("%#v\n", solution.Pick())
	fmt.Printf("%#v\n", solution.Pick())
	fmt.Printf("%#v\n", solution.Pick())
	fmt.Printf("%#v\n", solution.Pick())
	fmt.Printf("%#v\n", solution.Pick())
}
