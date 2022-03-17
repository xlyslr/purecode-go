package q0206

import (
	"fmt"
	leetcode "rhys.me/purecode-go/src/me/rhys/leetcode/editor/cn"
	"testing"
)

func TestReverseList(t *testing.T) {
	list := reverseList(&leetcode.ListNode{
		1, &leetcode.ListNode{
			2, &leetcode.ListNode{
				3, &leetcode.ListNode{
					4, &leetcode.ListNode{
						5, nil,
					},
				},
			},
		},
	})

	fmt.Printf("%#v\n", list)
}
