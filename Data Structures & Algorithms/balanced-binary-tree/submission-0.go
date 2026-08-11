/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func getHeight(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return 1 + max(getHeight(root.Left), getHeight(root.Right))
}

func isBalanced(root *TreeNode) bool {
    if root == nil { return true }
	// if root.Left == nil && root.Right == nil { return true }
	leftHeight := getHeight(root.Left)
	rightHeight := getHeight(root.Right)
	balanced := math.Abs(float64(leftHeight - rightHeight)) <= float64(1)
	if balanced && isBalanced(root.Left) && isBalanced(root.Right) {
		return true
	}
	return false
}
