//A path in a binary tree is a sequence of nodes where each pair of adjacent
//nodes in the sequence has an edge connecting them. A node can only appear in the
//sequence at most once. Note that the path does not need to pass through the root.
//
//
// The path sum of a path is the sum of the node's values in the path.
//
// Given the root of a binary tree, return the maximum path sum of any non-
//empty path.
//
//
// Example 1:
//
//
//Input: root = [1,2,3]
//Output: 6
//Explanation: The optimal path is 2 -> 1 -> 3 with a path sum of 2 + 1 + 3 = 6.
//
//
//
// Example 2:
//
//
//Input: root = [-10,9,20,null,null,15,7]
//Output: 42
//Explanation: The optimal path is 15 -> 20 -> 7 with a path sum of 15 + 20 + 7
//= 42.
//
//
//
// Constraints:
//
//
// The number of nodes in the tree is in the range [1, 3 * 10⁴].
// -1000 <= Node.val <= 1000
//
// Related Topics 树 深度优先搜索 动态规划 二叉树 👍 1416 👎 0

package q0124

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
var maxSum int

func maxPathSum(root *TreeNode) int {
	if root == nil {
		return 0
	}
	maxSum = -1001
	findOneSideMax(root)
	return maxSum
}

func findOneSideMax(root *TreeNode) int {
	if root == nil {
		return 0
	}
	leftPathSum := findOneSideMax(root.Left)
	rightPathSum := findOneSideMax(root.Right)
	if leftPathSum < 0 {
		leftPathSum = 0
	}
	if rightPathSum < 0 {
		rightPathSum = 0
	}
	// 第一种场景：left->node->right
	allSum := leftPathSum + rightPathSum + root.Val
	if allSum > maxSum {
		maxSum = allSum
	}
	// 第二种场景：经过node的单边分支的最优解。因为leftPathSum/RightPathSum已经处理了负数的场景，所以覆盖了只有当前node的最优解路径场景
	if leftPathSum > rightPathSum {
		return root.Val + leftPathSum
	} else {
		return root.Val + rightPathSum
	}
}

//leetcode submit region end(Prohibit modification and deletion)
