package q0785

import (
	"fmt"
	"testing"
)

func TestName(t *testing.T) {
	// [[1,3],[0,2],[1,3],[0,2]]
	testCase := [][]int{
		{1, 3}, {0, 2}, {1, 3}, {0, 2},
	}
	//
	fmt.Println(isBipartite(testCase))
}
