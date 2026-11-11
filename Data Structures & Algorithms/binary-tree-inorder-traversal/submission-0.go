/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func inorderTraversal(root *TreeNode) []int {
	var res []int
	if root == nil {
		return []int{}
	}
	if root.Left == nil && root.Right == nil {
		return []int{root.Val}
	}
	// Tree Traversal: https://en.wikipedia.org/wiki/Tree_traversal
	var stck []*TreeNode
	cur := root
	
	for len(stck) > 0 || cur != nil {
		// Process the left subtree
		// Move to the leftmost node
		// Push all nodes you encounter to the stack
		for cur != nil {
			stck = append(stck, cur)
			cur = cur.Left
		}
		// Process the node on top
		last := len(stck) - 1
		cur = stck[last]
		// TODO: Can be handled with top & bottom pointers to skip stack shrinking & growing
		stck = stck[:last]
		res = append(res, cur.Val)
		cur = cur.Right
	}
	return res
}
