//Given head, the head of a linked list, determine if the linked list has a
//cycle in it.
//
// There is a cycle in a linked list if there is some node in the list that can
//be reached again by continuously following the next pointer. Internally, pos is
//used to denote the index of the node that tail's next pointer is connected to.
//Note that pos is not passed as a parameter.
//
// Return true if there is a cycle in the linked list. Otherwise, return false.
//
//
//
// Example 1:
//
//
//Input: head = [3,2,0,-4], pos = 1
//Output: true
//Explanation: There is a cycle in the linked list, where the tail connects to
//the 1st node (0-indexed).
//
//
// Example 2:
//
//
//Input: head = [1,2], pos = 0
//Output: true
//Explanation: There is a cycle in the linked list, where the tail connects to
//the 0th node.
//
//
// Example 3:
//
//
//Input: head = [1], pos = -1
//Output: false
//Explanation: There is no cycle in the linked list.
//
//
//
// Constraints:
//
//
// The number of the nodes in the list is in the range [0, 10⁴].
// -10⁵ <= Node.val <= 10⁵
// pos is -1 or a valid index in the linked-list.
//
//
//
// Follow up: Can you solve it using O(1) (i.e. constant) memory?
// Related Topics 哈希表 链表 双指针 👍 1350 👎 0

package leetcode

//leetcode submit region begin(Prohibit modification and deletion)

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func hasCycle(head *ListNode) bool {
	if head == nil {
		return false
	}
	a, b := head, head
	// 如果b的出口是nil则说明没有环，这里只需要判断快指针即可
	// b.Next 判空实际是为了继续走一步，因为快指针一次循环走两步，实际是把循环里的判断挪到了循环外边
	for b != nil && b.Next != nil {
		// 头肯定相等，所以要先起跑，再判断，如果有环就会有相等
		a = a.Next
		b = b.Next.Next
		// 注意判断的是对象不是值
		if a == b {
			return true
		}
	}
	return false
}

//leetcode submit region end(Prohibit modification and deletion)
