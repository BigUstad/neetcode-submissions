/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func lowestCommonAncestor(root *TreeNode, p *TreeNode, q *TreeNode) *TreeNode {
	pval := p.Val
    qval := q.Val
    node := root
    for node != nil {
        // Parent node value
        parentVal := node.Val
        if pval > parentVal && qval > parentVal {
            // search right subtree in this case
            node = node.Right
        } else if pval < parentVal && qval < parentVal {
            //search left subtree in this case
            node = node.Left
        } else {
            // parentVal is now in the middle of pval & qval
            // hence a common ancestor
            return node
        }
    }
	// Making returning root as default
    return root
}
