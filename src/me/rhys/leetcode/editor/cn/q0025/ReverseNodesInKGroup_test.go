package q0025

import (
	"fmt"
	. "rhys.me/purecode-go/src/me/rhys/leetcode/editor/cn"
	"testing"
)

func TestName(t *testing.T) {
	testcase := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val: 4,
					Next: &ListNode{
						Val: 5,
						Next: &ListNode{
							Val: 6,
							Next: &ListNode{
								Val:  7,
								Next: &ListNode{Val: 8},
							},
						},
					},
				},
			},
		},
	}
	fmt.Printf("%#v\n", testcase)
	kGroup := reverseKGroup(testcase, 3)
	fmt.Printf("%#v\n", kGroup)

}
