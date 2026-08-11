/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func kthSmallest(root *TreeNode, k int) int {
    if root == nil || k <= 0 {
		return -1
	}
	var kNode *TreeNode
	var inOrder func(*TreeNode)
	inOrder = func(node *TreeNode) {
		if node == nil {
			return
		}
		inOrder(node.Left)
		// Process node. Just check to see if it is kth node
		k--
		if k == 0 {
			kNode = node
		}
		inOrder(node.Right)
	}
	inOrder(root)
	if kNode == nil { return -1 }
	return kNode.Val
}
