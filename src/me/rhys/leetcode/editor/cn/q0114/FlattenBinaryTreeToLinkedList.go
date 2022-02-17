//Given the root of a binary tree, flatten the tree into a "linked list":
//
//
// The "linked list" should use the same TreeNode class where the right child
//pointer points to the next node in the list and the left child pointer is always
//null.
// The "linked list" should be in the same order as a pre-order traversal of
//the binary tree.
//
//
//
// Example 1:
//
//
//Input: root = [1,2,5,3,4,null,6]
//Output: [1,null,2,null,3,null,4,null,5,null,6]
//
//
// Example 2:
//
//
//Input: root = []
//Output: []
//
//
// Example 3:
//
//
//Input: root = [0]
//Output: [0]
//
//
//
// Constraints:
//
//
// The number of nodes in the tree is in the range [0, 2000].
// -100 <= Node.val <= 100
//
//
//
//Follow up: Can you flatten the tree in-place (with O(1) extra space)? Related
//Topics 栈 树 深度优先搜索 链表 二叉树 👍 1063 👎 0

package q0114

import . "rhys.me/purecode-go/src/me/rhys/leetcode/editor/cn"

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func flatten(root *TreeNode) {
	if root == nil {
		return
	}
	flatten(root.Left)
	flatten(root.Right)

	right := root.Right
	root.Right = root.Left
	root.Left = nil

	r := root
	for r.Right != nil {
		r = r.Right
	}
	r.Right = right
}

//leetcode submit region end(Prohibit modification and deletion)
