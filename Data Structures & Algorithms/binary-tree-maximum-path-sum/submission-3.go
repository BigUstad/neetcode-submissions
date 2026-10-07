/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func maxPathSum(root *TreeNode) int {
	if root == nil { return 0 }
    maxRes := root.Val
	var dfs func(*TreeNode) int
	dfs = func(root *TreeNode) int {
		if root == nil {
			return 0
		}
		leftMax := dfs(root.Left)
		rightMax := dfs(root.Right)
		// Ignoring the negative values in the path sum computation
		leftMax = max(leftMax, 0)
		rightMax = max(rightMax, 0)
		res :=  root.Val + leftMax + rightMax
		maxRes = max(maxRes, res)
		// fmt.Println(root.Val, ": ", res)
		return max(leftMax ,rightMax) + root.Val
	}
	dfs(root)
	return maxRes
}
