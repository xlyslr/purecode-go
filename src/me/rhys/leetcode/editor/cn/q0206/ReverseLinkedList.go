//Given the head of a singly linked list, reverse the list, and return the
//reversed list.
//
//
// Example 1:
//
//
//Input: head = [1,2,3,4,5]
//Output: [5,4,3,2,1]
//
//
// Example 2:
//
//
//Input: head = [1,2]
//Output: [2,1]
//
//
// Example 3:
//
//
//Input: head = []
//Output: []
//
//
//
// Constraints:
//
//
// The number of nodes in the list is the range [0, 5000].
// -5000 <= Node.val <= 5000
//
//
//
// Follow up: A linked list can be reversed either iteratively or recursively.
//Could you implement both?
// Related Topics 递归 链表 👍 2315 👎 0

package q0206

import . "rhys.me/purecode-go/src/me/rhys/leetcode/editor/cn"

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseList(head *ListNode) *ListNode {
	dummy := &ListNode{}

	//if head == nil || head.Next == nil {
	//	return head
	//}
	// == 方案一
	// 关键思路：1. 怎么返回最后的头，2. 反转有两个动作，下一个节点的指针指向自己，自己指向前一个指针，其实都是指向前一个指针
	// 递归处理思路：递归是先处理最后面的
	//last := reverseList(head.Next)
	//head.Next.Next = head
	// 这一步虽然置空，但上一层递归，会把他重新指向前一个，最终是为了第一个元素指向空
	//head.Next = nil
	//return last

	// == 方案二
	//curr := head
	//var prev *ListNode
	//for curr != nil {
	//	tmp := curr.Next
	//	curr.Next = prev
	//	prev = curr
	//	curr = tmp
	//}

	// == 头插法
	// dummy 就是pre
	dummy.Next = head
	cur := head
	var nex *ListNode
	for cur.Next != nil {
		nex = cur.Next
		cur.Next = nex.Next
		nex.Next = dummy.Next
		dummy.Next = nex
	}
	return dummy.Next
}

//leetcode submit region end(Prohibit modification and deletion)
