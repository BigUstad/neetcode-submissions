/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func buildInOrderMap(inordermap map[int]int, inorder []int) {
	for i, n := range inorder {
		inordermap[n] = i
	}
}

func buildTree(preorder []int, inorder []int) *TreeNode {
    var inordermap map[int]int
	var root *TreeNode
	var buildTreeHelper func(l, r int) *TreeNode
	if len(preorder) == 0 || len(inorder) == 0 {
		return root
	}
	inordermap = make(map[int]int)
	buildInOrderMap(inordermap, inorder)
	// fmt.Println(len(inordermap))
	preorderIdx := 0
	buildTreeHelper = func(l, r int) *TreeNode {
		if l > r || preorderIdx == len(preorder) { return nil }
		p := preorder[preorderIdx]
		preorderIdx++
		cur := &TreeNode{Val: p, Left: nil, Right: nil}
		if root == nil {
			// fmt.Print("Root is: ")
			fmt.Println(p)
			root = cur
		}
		cur.Left = buildTreeHelper(l, inordermap[p] - 1)
		cur.Right = buildTreeHelper(inordermap[p] + 1, r)
		return cur
	}
	buildTreeHelper(0, len(preorder)-1)
	return root
}
