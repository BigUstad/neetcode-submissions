/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 
func preorderTraversal(root *TreeNode) []int {
    var res []int
	if root == nil {
		return []int{}
	}
	// Iterative Tree Traversal: https://en.wikipedia.org/wiki/Tree_traversal
	// Node-Left-Right
	var stck []*TreeNode
    // NOTE: This 'nil' on the bottom signifies that the iteration needs to end.
    // Important step for pre-order.
    // nil is an indicator or a sentinal value
	stck = append(stck, nil)
	for root != nil {
		res = append(res, root.Val)
        // fmt.Println(res)
		// Putting the right onto the stack first.
		// So that it gets popped after the left
		if root.Right != nil {
			stck = append(stck, root.Right)
		}
		// Process left subtree
		if root.Left != nil {
			root = root.Left
		} else if len(stck) > 0 {
			last := len(stck)-1
			root = stck[last]
			stck = stck[:last]
		}
        // fmt.Println(stck)
	}
	return res
}
