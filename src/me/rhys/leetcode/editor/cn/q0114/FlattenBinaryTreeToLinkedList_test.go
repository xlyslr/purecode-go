package q0114

import (
	"fmt"
	. "rhys.me/purecode-go/src/me/rhys/leetcode/editor/cn"
	"testing"
)

func TestFlatten(t *testing.T) {
	testCase := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val:   3,
				Left:  nil,
				Right: nil,
			},
			Right: &TreeNode{
				Val:   4,
				Left:  nil,
				Right: nil,
			},
		},
		Right: &TreeNode{
			Val:  5,
			Left: nil,
			Right: &TreeNode{
				Val:   6,
				Left:  nil,
				Right: nil,
			},
		},
	}
	flatten(testCase)
	fmt.Printf("%#v\n", testCase)
}
