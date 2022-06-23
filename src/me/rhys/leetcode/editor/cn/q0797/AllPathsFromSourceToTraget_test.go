package q0797

import (
	"fmt"
	"testing"
)

func TestName(t *testing.T) {
	// [[4,3,1],[3,2,4],[],[4],[]]
	testCase := [][]int{
		{4, 3, 1}, {3, 2, 4}, {}, {4}, {},
	}
	target := allPathsSourceTarget(testCase)
	fmt.Printf("%#v\n", target)
}

func Test2(t *testing.T) {
	m := make(map[string]int)
	m["xx"] = 1
	x := m["xxx"]
	fmt.Println(x)
	val, ok := m["xxx"]
	if ok {
		fmt.Println(val)
	}
}
