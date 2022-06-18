//Given the root of a binary search tree, and an integer k, return the kᵗʰ
//smallest value (1-indexed) of all the values of the nodes in the tree.
//
//
// Example 1:
//
//
//Input: root = [3,1,4,null,2], k = 1
//Output: 1
//
//
// Example 2:
//
//
//Input: root = [5,3,6,2,4,null,null,1], k = 3
//Output: 3
//
//
//
// Constraints:
//
//
// The number of nodes in the tree is n.
// 1 <= k <= n <= 10⁴
// 0 <= Node.val <= 10⁴
//
//
//
// Follow up: If the BST is modified often (i.e., we can do insert and delete
//operations) and you need to find the kth smallest frequently, how would you
//optimize?
// Related Topics 树 深度优先搜索 二叉搜索树 二叉树 👍 634 👎 0

package q0230

import . "rhys.me/purecode-go/src/me/rhys/leetcode/editor/cn"

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * 思路：BST的特性，中序遍历就是有序，所以中序遍历到K就可以
 * 优化：可以在节点上记录当前节点的子节点总数，那么node.left.size就可以推断出当前第几位
 */
var idx int
var target int
var res int

func kthSmallest(root *TreeNode, k int) int {
	idx = 0
	res = 0
	target = k
	traverse(root)
	return res
}

func traverse(root *TreeNode) {
	if root == nil {
		return
	}
	traverse(root.Left)
	idx++
	if idx == target {
		res = root.Val
		return
	}
	traverse(root.Right)
}

//leetcode submit region end(Prohibit modification and deletion)
