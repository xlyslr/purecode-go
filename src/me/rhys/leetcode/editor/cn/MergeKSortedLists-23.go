//You are given an array of k linked-lists lists, each linked-list is sorted in
//ascending order.
//
// Merge all the linked-lists into one sorted linked-list and return it.
//
//
// Example 1:
//
//
//Input: lists = [[1,4,5],[1,3,4],[2,6]]
//Output: [1,1,2,3,4,4,5,6]
//Explanation: The linked-lists are:
//[
//  1->4->5,
//  1->3->4,
//  2->6
//]
//merging them into one sorted list:
//1->1->2->3->4->4->5->6
//
//
// Example 2:
//
//
//Input: lists = []
//Output: []
//
//
// Example 3:
//
//
//Input: lists = [[]]
//Output: []
//
//
//
// Constraints:
//
//
// k == lists.length
// 0 <= k <= 10⁴
// 0 <= lists[i].length <= 500
// -10⁴ <= lists[i][j] <= 10⁴
// lists[i] is sorted in ascending order.
// The sum of lists[i].length will not exceed 10⁴.
//
// Related Topics 链表 分治 堆（优先队列） 归并排序 👍 1746 👎 0

package leetcode

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeKLists(lists []*ListNode) *ListNode {
	dummy := &ListNode{}
	p := dummy
	tmp := make([]*ListNode, len(lists))
	//tmp := [...]*ListNode{}
	// 寻找最小的一个
	for i, list := range lists {
		tmp[i] = list
	}
	for minNode := findMin(tmp); minNode != nil; minNode = findMin(tmp) {
		p.Next = minNode
		p = p.Next
	}

	return dummy.Next
}

func findMin(tmp []*ListNode) *ListNode {
	var minNode *ListNode
	minIdx := -1
	for i, node := range tmp {
		if node == nil {
			continue
		}
		if minNode == nil || node.Val <= minNode.Val {
			minNode = node
			minIdx = i
		}
	}
	if minIdx != -1 && minNode != nil {
		tmp[minIdx] = minNode.Next
	}
	return minNode
}

//leetcode submit region end(Prohibit modification and deletion)
