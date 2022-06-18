//Given the root of a binary tree, determine if it is a valid binary search
//tree (BST).
//
// A valid BST is defined as follows:
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
//Input: root = [2,1,3]
//Output: true
//
//
// Example 2:
//
//
//Input: root = [5,1,4,null,null,3,6]
//Output: false
//Explanation: The root node's value is 5 but its right child's value is 4.
//
//
//
// Constraints:
//
//
// The number of nodes in the tree is in the range [1, 10⁴].
// -2³¹ <= Node.val <= 2³¹ - 1
//
// Related Topics 树 深度优先搜索 二叉搜索树 二叉树 👍 1629 👎 0

package q0098

import . "rhys.me/purecode-go/src/me/rhys/leetcode/editor/cn"

//leetcode submit region begin(Prohibit modification and deletion)
/**
https://www.wolai.com/rhys/m7PcJAfh7sdSmssMD2cS7W#iPXqcdwk75uRrzc4RWw1wB
*/
func isValidBST(root *TreeNode) bool {
	return validBST(root, nil, nil)
}

func validBST(root, min, max *TreeNode) bool {
	if root == nil {
		return true
	}
	if min != nil && root.Val <= min.Val {
		return false
	}
	if max != nil && root.Val >= max.Val {
		return false
	}

	l := validBST(root.Left, min, root)
	r := validBST(root.Right, root, max)
	return l && r
}

//leetcode submit region end(Prohibit modification and deletion)
