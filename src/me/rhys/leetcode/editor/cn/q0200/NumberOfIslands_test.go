package q0200

import (
	"fmt"
	"testing"
)

func TestNumIslands(t *testing.T) {
	testCase := [][]byte{
		{1, 1, 1, 1, 0},
		{1, 1, 0, 1, 0},
		{1, 1, 0, 0, 0},
		{0, 0, 0, 0, 0},
	}
	islands := numIslands(testCase)
	fmt.Printf("%#v\n", islands)
}
