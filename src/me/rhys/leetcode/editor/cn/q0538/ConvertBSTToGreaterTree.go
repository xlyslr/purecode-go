//Given the root of a Binary Search Tree (BST), convert it to a Greater Tree
//such that every key of the original BST is changed to the original key plus the
//sum of all keys greater than the original key in BST.
//
// As a reminder, a binary search tree is a tree that satisfies these
//constraints:
//
//
// The left subtree of a node contains only nodes with keys less than the
//node's key.
// The right subtree of a node contains only nodes with keys greater than the
//node's key.
// Both the left and right subtrees must also be binary search trees.
//
//
//
// Example 1:
//
//
//Input: root = [4,1,6,0,2,5,7,null,null,null,3,null,null,null,8]
//Output: [30,36,21,36,35,26,15,null,null,null,33,null,null,null,8]
//
//
// Example 2:
//
//
//Input: root = [0,null,1]
//Output: [1,null,1]
//
//
//
// Constraints:
//
//
// The number of nodes in the tree is in the range [0, 10⁴].
// -10⁴ <= Node.val <= 10⁴
// All the values in the tree are unique.
// root is guaranteed to be a valid binary search tree.
//
//
//
// Note: This question is the same as 1038: https://leetcode.com/problems/
//binary-search-tree-to-greater-sum-tree/
// Related Topics 树 深度优先搜索 二叉搜索树 二叉树 👍 735 👎 0

package q0538

import . "rhys.me/purecode-go/src/me/rhys/leetcode/editor/cn"

//leetcode submit region begin(Prohibit modification and deletion)
/**
 */
var sum int

func convertBST(root *TreeNode) *TreeNode {
	sum = 0
	if root == nil {
		return nil
	}
	dest := &TreeNode{}
	traverse(root, dest)
	return dest
}

func traverse(root, dest *TreeNode) {
	if root.Right != nil {
		dest.Right = &TreeNode{}
		traverse(root.Right, dest.Right)
	}
	sum += root.Val
	dest.Val = sum
	if root.Left != nil {
		dest.Left = &TreeNode{}
		traverse(root.Left, dest.Left)
	}
}

//leetcode submit region end(Prohibit modification and deletion)
