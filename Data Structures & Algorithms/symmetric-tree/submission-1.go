/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// In Leetcode I'll implement BFS
// In Neetcode I'll implement DFS
func isSymmetric(root *TreeNode) bool {
    var checkSymmetry func(*TreeNode, *TreeNode) bool
	checkSymmetry = func(cur1, cur2 *TreeNode) bool {
		if cur1 == nil && cur2 == nil {
			return true
		}
		if cur1 == nil || cur2 == nil {
			return false
		}
		if cur1.Val != cur2.Val {
			return false
		}
		return checkSymmetry(cur1.Left, cur2.Right) &&
				checkSymmetry(cur1.Right, cur2.Left)
	}

	return checkSymmetry(root, root)
}
