/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// Choosing Post order as it processes both children nodes
// Before root
func diameterBTHelper(root *TreeNode, diameter *int) int {
	if root == nil { return 0 }
	// fmt.Println(diameter)
	leftLength := diameterBTHelper(root.Left, diameter)
	rightLength := diameterBTHelper(root.Right, diameter)
	/* fmt.Print(root.Val)
	fmt.Print(" - ")
	fmt.Print(leftLength)
	fmt.Print(",")
	fmt.Println(rightLength) */
	*diameter = max(*diameter, (leftLength + rightLength))
	// + 1 accounts for node/root
	return max(leftLength, rightLength) + 1 
}

func diameterOfBinaryTree(root *TreeNode) int {
    if root == nil { return 0 }
    if root.Left == nil && root.Right == nil { return 0 }
	var diameter int
	fmt.Println(diameterBTHelper(root, &diameter))
	return diameter
}
