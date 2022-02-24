package q0704

import (
	"fmt"
	"testing"
)

func TestSearch(t *testing.T) {
	i := search([]int{-1, 0, 3, 5, 9, 12}, 9)
	fmt.Println(i)
}
