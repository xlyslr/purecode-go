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
// Related Topics éå½ é¾è¡¨ ð 2315 ð 0

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
	// == æ¹æ¡ä¸
	// å³é®æè·¯ï¼1. æä¹è¿åæåçå¤´ï¼2. åè½¬æä¸¤ä¸ªå¨ä½ï¼ä¸ä¸ä¸ªèç¹çæéæåèªå·±ï¼èªå·±æååä¸ä¸ªæéï¼å¶å®é½æ¯æååä¸ä¸ªæé
	// éå½å¤çæè·¯ï¼éå½æ¯åå¤çæåé¢ç
	//last := reverseList(head.Next)
	//head.Next.Next = head
	// è¿ä¸æ­¥è½ç¶ç½®ç©ºï¼ä½ä¸ä¸å±éå½ï¼ä¼æä»éæ°æååä¸ä¸ªï¼æç»æ¯ä¸ºäºç¬¬ä¸ä¸ªåç´ æåç©º
	//head.Next = nil
	//return last

	// == æ¹æ¡äº
	//curr := head
	//var prev *ListNode
	//for curr != nil {
	//	tmp := curr.Next
	//	curr.Next = prev
	//	prev = curr
	//	curr = tmp
	//}

	// == å¤´ææ³
	// dummy å°±æ¯pre
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
