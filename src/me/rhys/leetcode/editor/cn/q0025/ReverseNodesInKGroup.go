//Given the head of a linked list, reverse the nodes of the list k at a time,
//and return the modified list.
//
// k is a positive integer and is less than or equal to the length of the
//linked list. If the number of nodes is not a multiple of k then left-out nodes, in
//the end, should remain as it is.
//
// You may not alter the values in the list's nodes, only nodes themselves may
//be changed.
//
//
// Example 1:
//
//
//Input: head = [1,2,3,4,5], k = 2
//Output: [2,1,4,3,5]
//
//
// Example 2:
//
//
//Input: head = [1,2,3,4,5], k = 3
//Output: [3,2,1,4,5]
//
//
//
// Constraints:
//
//
// The number of nodes in the list is n.
// 1 <= k <= n <= 5000
// 0 <= Node.val <= 1000
//
//
//
// Follow-up: Can you solve the problem in O(1) extra memory space?
// Related Topics 递归 链表 👍 1694 👎 0

package q0025

import (
	. "rhys.me/purecode-go/src/me/rhys/leetcode/editor/cn"
)

//leetcode submit region begin(Prohibit modification and deletion)
type Stack struct {
	arr []int
	idx int
}

func NewStack(size int) *Stack {
	return &Stack{
		arr: make([]int, size),
		idx: -1,
	}
}

func (s *Stack) Push(val int) {
	s.idx++
	s.arr[s.idx] = val
}

func (s *Stack) Pop() int {
	val := s.arr[s.idx]
	s.idx--
	return val
}

func (s *Stack) isEmpty() bool {
	if s.idx == -1 {
		return true
	}
	return false
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	dummy := &ListNode{}
	res := dummy
	tmp := head
	count := 0
	stack := NewStack(k)
	for tmp != nil {
		if count == k {
			count = 0
			for !stack.isEmpty() {
				res.Next = &ListNode{Val: stack.Pop()}
				res = res.Next
			}
		}
		count++
		stack.Push(tmp.Val)
		tmp = tmp.Next
	}
	for !stack.isEmpty() {
		res.Next = &ListNode{Val: stack.Pop()}
		res = res.Next
	}
	return dummy.Next
}

//leetcode submit region end(Prohibit modification and deletion)
