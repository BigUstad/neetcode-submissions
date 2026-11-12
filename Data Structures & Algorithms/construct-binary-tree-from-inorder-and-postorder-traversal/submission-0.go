/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */


func buildTree(inorder []int, postorder []int) *TreeNode {
	inordermap := make(map[int]int)
	var buildInOrderMap func()

	buildInOrderMap = func() {
		for i, n := range inorder {
			inordermap[n] = i
		}
	}
	var root *TreeNode

	var dfsHelper func(int, int) *TreeNode
	// Let me start from the end for list.
	// For preorder I started from the start of the list
	// Post-Order is LRN
	postOrderIdx := len(postorder) - 1
	dfsHelper = func(l, r int) *TreeNode {
		if l > r {
			return nil
		}
		cur := &TreeNode{Val: postorder[postOrderIdx]}
		postOrderIdx--
		if root == nil {
			root = cur
		}
		inorderIdx := inordermap[cur.Val]
		// The root of the right sub-tree is the right node of cur
		// PostOrder is LRN. We got the Node, we get the right
		cur.Right = dfsHelper(inorderIdx+1, r)
		cur.Left = dfsHelper(l, inorderIdx-1)
		return cur
	}

	buildInOrderMap()
	dfsHelper(0, len(inorder)-1)

	return root
}
