package q0021

import (
	"fmt"
	"rhys.me/purecode-go/src/me/rhys/leetcode/editor/cn"
	"testing"
)

func TestMergeTwoLists(t *testing.T) {
	l1 := &leetcode.ListNode{
		Val: 1,
		Next: &leetcode.ListNode{
			Val: 2,
			Next: &leetcode.ListNode{
				Val:  4,
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
	l := mergeTwoLists(l1, l2)
	fmt.Printf("%+v\n", l)
}
