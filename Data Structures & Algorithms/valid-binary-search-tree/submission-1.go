/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// I assume the space & time performance is affected by "container/list"

func getInorder(root *TreeNode) []int {
    stack := list.New()
	var inorder []int
	cur := root
	for stack.Len() > 0 || cur != nil {
		if cur != nil {
			stack.PushFront(cur)
			cur = cur.Left
		} else {
			cur = stack.Front().Value.(*TreeNode)
			stack.Remove(stack.Front())
			inorder = append(inorder, cur.Val)
            cur = cur.Right
		}
		// fmt.Println(stack.Len())
	}
	stack = nil
	return inorder
}

func isValidBST(root *TreeNode) bool {
	if root == nil { return true }
	if root.Right == nil && root.Left == nil { return true }
	inorder := getInorder(root)
	l := len(inorder) - 1
	// inorder. compare item at i to i + 1.
	// Ascending order - valid bst
	for i, n := range inorder {
		// last item. no next item to compare
		if i == l {
			break
		}
		if n >= inorder[i+1] {
			return false
		}
	}
	return true
}
