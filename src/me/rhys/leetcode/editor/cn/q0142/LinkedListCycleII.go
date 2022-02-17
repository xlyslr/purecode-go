//Given the head of a linked list, return the node where the cycle begins. If
//there is no cycle, return null.
//
// There is a cycle in a linked list if there is some node in the list that can
//be reached again by continuously following the next pointer. Internally, pos is
//used to denote the index of the node that tail's next pointer is connected to (0
//-indexed). It is -1 if there is no cycle. Note that pos is not passed as a
//parameter.
//
// Do not modify the linked list.
//
//
// Example 1:
//
//
//Input: head = [3,2,0,-4], pos = 1
//Output: tail connects to node index 1
//Explanation: There is a cycle in the linked list, where tail connects to the
//second node.
//
//
// Example 2:
//
//
//Input: head = [1,2], pos = 0
//Output: tail connects to node index 0
//Explanation: There is a cycle in the linked list, where tail connects to the
//first node.
//
//
// Example 3:
//
//
//Input: head = [1], pos = -1
//Output: no cycle
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
// Related Topics 哈希表 链表 双指针 👍 1375 👎 0

package q0142

import . "rhys.me/purecode-go/src/me/rhys/leetcode/editor/cn"

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func DetectCycle(head *ListNode) *ListNode {
	a, b := head, head
	flag := false
	// 首先找到环的位置，然后让b从环的位置出发，a从头出发，同速前进，遇到相同的点，则是环的起点
	for b != nil && b.Next != nil {
		a = a.Next
		b = b.Next.Next
		if a == b {
			flag = true
			break
		}
	}

	if flag {
		a = head
		for a != b {
			a = a.Next
			b = b.Next
		}
		return a
	}

	return nil
}

//leetcode submit region end(Prohibit modification and deletion)
