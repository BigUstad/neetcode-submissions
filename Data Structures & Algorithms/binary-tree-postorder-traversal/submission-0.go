/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func postorderTraversal(root *TreeNode) []int {
    var res []int
	var resL *list.List
	if root == nil { return []int{} }
	if root.Left == nil && root.Right == nil { return []int{root.Val} }
	var stck []*TreeNode
	var cur *TreeNode
	resL = list.New()
	// Tree Traversal: https://en.wikipedia.org/wiki/Tree_traversal#Pre-order, LRN
	stck = append(stck, root)
	for len(stck) > 0 {
		last := len(stck) - 1
		cur = stck[last]
		stck = stck[:last]
		resL.PushFront(cur)
		if cur.Left != nil {
			stck = append(stck, cur.Left)
		}
		if cur.Right != nil {
			stck = append(stck, cur.Right)
		}
	}
	for e := resL.Front(); e != nil; e = e.Next() {
		res = append(res, e.Value.(*TreeNode).Val)
	}
	return res
}
