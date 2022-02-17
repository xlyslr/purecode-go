package q0023

import (
	"fmt"
	"rhys.me/purecode-go/src/me/rhys/leetcode/editor/cn"
	"testing"
)

func TestMergeKLists(t *testing.T) {
	l1 := &leetcode.ListNode{
		Val: 1,
		Next: &leetcode.ListNode{
			Val: 4,
			Next: &leetcode.ListNode{
				Val:  5,
				Next: nil,
			},
		},
	}
	l2 := &leetcode.ListNode{
		Val: 1,
		Next: &leetcode.ListNode{
			Val: 3,
			Next: &leetcode.ListNode{
				Val:  4,
				Next: nil,
			},
		},
	}
	l3 := &leetcode.ListNode{
		Val: 2,
		Next: &leetcode.ListNode{
			Val:  6,
			Next: nil,
		},
	}
	lists := mergeKLists([]*leetcode.ListNode{l1, l2, l3})
	fmt.Printf("%#v\n", lists)
}
