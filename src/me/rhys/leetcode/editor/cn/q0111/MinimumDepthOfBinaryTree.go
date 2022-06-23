//Given a binary tree, find its minimum depth.
//
// The minimum depth is the number of nodes along the shortest path from the
//root node down to the nearest leaf node.
//
// Note: A leaf is a node with no children.
//
//
// Example 1:
//
//
//Input: root = [3,9,20,null,null,15,7]
//Output: 2
//
//
// Example 2:
//
//
//Input: root = [2,null,3,null,4,null,5,null,6]
//Output: 5
//
//
//
// Constraints:
//
//
// The number of nodes in the tree is in the range [0, 10⁵].
// -1000 <= Node.val <= 1000
//
// Related Topics 树 深度优先搜索 广度优先搜索 二叉树 👍 780 👎 0

package q0111

import . "rhys.me/purecode-go/src/me/rhys/leetcode/editor/cn"

//leetcode submit region begin(Prohibit modification and deletion)
/**
  使用广度遍历，找到一个叶子节点就可以返回了，因为这时深度是最小的
*/
func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	var queue []*TreeNode
	var depths []int
	queue = append(queue, root)
	depths = append(depths, 1)
	for i := 0; i < len(queue); i++ {
		node := queue[i]
		depth := depths[i]
		if node.Left == nil && node.Right == nil {
			return depth
		}
		if node.Left != nil {
			queue = append(queue, node.Left)
			depths = append(depths, depth+1)
		}
		if node.Right != nil {
			queue = append(queue, node.Right)
			depths = append(depths, depth+1)
		}
	}
	return 0
}

//leetcode submit region end(Prohibit modification and deletion)
