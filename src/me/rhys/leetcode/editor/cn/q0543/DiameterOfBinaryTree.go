//Given the root of a binary tree, return the length of the diameter of the
//tree.
//
// The diameter of a binary tree is the length of the longest path between any
//two nodes in a tree. This path may or may not pass through the root.
//
// The length of a path between two nodes is represented by the number of edges
//between them.
//
//
// Example 1:
//
//
//Input: root = [1,2,3,4,5]
//Output: 3
//Explanation: 3 is the length of the path [4,2,1,3] or [5,2,1,3].
//
//
// Example 2:
//
//
//Input: root = [1,2]
//Output: 1
//
//
//
// Constraints:
//
//
// The number of nodes in the tree is in the range [1, 10⁴].
// -100 <= Node.val <= 100
//
// Related Topics 树 深度优先搜索 二叉树 👍 1064 👎 0

package q0543

import (
	"math"
	. "rhys.me/purecode-go/src/me/rhys/leetcode/editor/cn"
)

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for a binary tree node.
 * 分解子问题：任意两个节点最大的长度可以认为是左侧最大深度和右侧最大深度的和
 */
var max = 0

func diameterOfBinaryTree(root *TreeNode) int {
	max = 0
	maxDepth(root)
	return max
}

func maxDepth(root *TreeNode) int {
	var res = 0
	if root == nil {
		return res
	}
	l := maxDepth(root.Left)
	r := maxDepth(root.Right)
	s := l + r
	max = int(math.Max(float64(s), float64(max)))
	res = int(math.Max(float64(l+1), float64(r+1)))
	return res
}

//leetcode submit region end(Prohibit modification and deletion)
