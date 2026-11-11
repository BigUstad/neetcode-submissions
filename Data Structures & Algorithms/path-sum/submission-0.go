/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// Recursion check. Once the recursion path reaches a leaf node,
// check if the sum is 0.
// Result will be left subtree or right subtree check
// So, the check order is Node-Left-Right (Preorder?)
func pathSumHelper(root *TreeNode, sum int) bool {
	if root == nil {
		return false
	}
	// Process Node
	sum -= root.Val
	// fmt.Println(root.Val, ", leftover sum = ", sum)
	if root.Left == nil && root.Right == nil && sum == 0 {
		return true
	}
	// root nil check has been done.
	// So root.Val is available
	return pathSumHelper(root.Left, sum) || pathSumHelper(root.Right, sum)
}

func hasPathSum(root *TreeNode, targetSum int) bool {
	return pathSumHelper(root, targetSum)
}
