package q0380

import (
	"fmt"
	"testing"
)

func TestRandomizedSet_Insert(t *testing.T) {
	rs := Constructor()
	fmt.Printf("%#v\n", rs.Insert(1))
	fmt.Printf("%#v\n", rs.Remove(2))
	fmt.Printf("%#v\n", rs.Insert(2))
	fmt.Printf("%#v\n", rs.GetRandom())
	fmt.Printf("%#v\n", rs.Remove(1))
	fmt.Printf("%#v\n", rs.Insert(2))
	fmt.Printf("%#v\n", rs.GetRandom())
}

func TestRandomizedSet_Remove(t *testing.T) {
}
func TestRandomizedSet_GetRandom(t *testing.T) {
}
