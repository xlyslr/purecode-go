//Given the root of a binary tree, return its maximum depth.
//
// A binary tree's maximum depth is the number of nodes along the longest path
//from the root node down to the farthest leaf node.
//
//
// Example 1:
//
//
//Input: root = [3,9,20,null,null,15,7]
//Output: 3
//
//
// Example 2:
//
//
//Input: root = [1,null,2]
//Output: 2
//
//
//
// Constraints:
//
//
// The number of nodes in the tree is in the range [0, 10⁴].
// -100 <= Node.val <= 100
//
// Related Topics 树 深度优先搜索 广度优先搜索 二叉树 👍 1272 👎 0

package q0104

import (
	"math"
	. "rhys.me/purecode-go/src/me/rhys/leetcode/editor/cn"
)

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for a binary tree node.
 * 分解子问题：求二叉树的最大高度，就是分别求左右子树的最大高度，然后选出最大的
 */
func maxDepth(root *TreeNode) int {
	var max = 0
	if root == nil {
		return max
	}
	l := maxDepth(root.Left)
	r := maxDepth(root.Right)
	max = int(math.Max(float64(r+1), float64(l+1)))
	return max
}

//leetcode submit region end(Prohibit modification and deletion)
