//Given the head of a singly linked list and two integers left and right where
//left <= right, reverse the nodes of the list from position left to position
//right, and return the reversed list.
//
//
// Example 1:
//
//
//Input: head = [1,2,3,4,5], left = 2, right = 4
//Output: [1,4,3,2,5]
//
//
// Example 2:
//
//
//Input: head = [5], left = 1, right = 1
//Output: [5]
//
//
//
// Constraints:
//
//
// The number of nodes in the list is n.
// 1 <= n <= 500
// -500 <= Node.val <= 500
// 1 <= left <= right <= n
//
//
//
//Follow up: Could you do it in one pass? Related Topics 链表 👍 1179 👎 0

package q0092

import (
	. "rhys.me/purecode-go/src/me/rhys/leetcode/editor/cn"
)

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func reverseBetween(head *ListNode, left int, right int) *ListNode {
	dummyNode := &ListNode{}
	dummyNode.Next = head
	pre := dummyNode
	//
	for i := 1; i < left; i++ {
		pre = pre.Next
	}

	// 代表要插入头部的元素
	var nex *ListNode
	// 代表反转区域的第一个元素
	cur := pre.Next
	for i := left; i < right; i++ {
		nex = cur.Next
		// 一定是用Next接受才是指针变更
		// 向后走一位，将cur.Next指向下一轮循环要头插的元素
		cur.Next = nex.Next
		// 头插第一步，将新头指向旧头；pre.Next可不等同于cur，pre.Next是上次的头
		nex.Next = pre.Next
		// 头插第二部，将新头放在第一位
		pre.Next = nex
	}

	return dummyNode.Next
}

//leetcode submit region end(Prohibit modification and deletion)
