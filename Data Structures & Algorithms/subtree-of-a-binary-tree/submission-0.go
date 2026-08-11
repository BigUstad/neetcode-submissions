/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

import "slices"

func dfsHelper(root *TreeNode, dfsOutputs *[]string) (dfsOutput string) {
	if root == nil {
		return ""
	}
	leftOutput := dfsHelper(root.Left, dfsOutputs)
	rightOutput := dfsHelper(root.Right, dfsOutputs)
	dfsOutput = strings.Join([]string{leftOutput, rightOutput, strconv.Itoa(root.Val)}, ",")
	// fmt.Print(root.Val)
	// fmt.Print(" - ")
	// fmt.Println(dfsOutput)
	*dfsOutputs = append(*dfsOutputs, dfsOutput)
	return
}

func isSubtree(root *TreeNode, subRoot *TreeNode) bool {
    if root == nil && subRoot == nil { return true }
	if root == nil && subRoot != nil { return false }
	var rootDfsOutputs, subrootDfsOutputs []string
	_ = dfsHelper(root, &rootDfsOutputs)
	// fmt.Println("xxxxxxxxxx")
	subrootDfsOutput := dfsHelper(subRoot, &subrootDfsOutputs)
	// fmt.Println(len(rootDfsOutputs))
	// fmt.Println(strings.Join(rootDfsOutputs, "\n"))
	// fmt.Println(subrootDfsOutput)
	// return strings.Contains(rootDfsOutput, subrootDfsOutput)
	return slices.Contains(rootDfsOutputs, subrootDfsOutput)
}
