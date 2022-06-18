package q0695

import (
	"fmt"
	"testing"
)

func TestMaxAreaOfIsland(t *testing.T) {
	testCase := [][]int{
		{1, 1, 1, 1, 0},
		{1, 1, 0, 1, 0},
		{1, 1, 0, 0, 0},
		{0, 0, 0, 0, 0},
	}
	island := maxAreaOfIsland(testCase)
	fmt.Printf("%#v\n", island)
}
