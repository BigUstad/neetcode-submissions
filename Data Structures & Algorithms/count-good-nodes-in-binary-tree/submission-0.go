/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// pursueParentMap - All nodes are <= "this" node value. No nil input check
func pursueParentMap(parentmap map[*TreeNode]*TreeNode, n *TreeNode) bool {
	nVal := n.Val
	for n != nil {
		n = parentmap[n]
		if n != nil && n.Val > nVal {
			return false
		}
	}
	return true
}

// Post-order
func buildParentMap(root *TreeNode, parentmap map[*TreeNode]*TreeNode, dfsoutput *[]*TreeNode) {
	if root == nil {
		return
	}
	buildParentMap(root.Left, parentmap, dfsoutput)
	buildParentMap(root.Right, parentmap, dfsoutput)
	// Process node
	*dfsoutput = append(*dfsoutput, root)
	if root.Left != nil { parentmap[root.Left] = root }
	if root.Right != nil { parentmap[root.Right] = root }
}

func goodNodes(root *TreeNode) int {
    if root == nil {return 0}
	if root.Right == nil && root.Left == nil {return 1}
	goodCount := 0 // Root
	parentmap := make(map[*TreeNode]*TreeNode)
	parentmap[root] = nil
	nodelist := make([]*TreeNode, 0)
	buildParentMap(root, parentmap, &nodelist)
	for _, n := range nodelist {
		if pursueParentMap(parentmap, n) {
			goodCount++
		}
	}
	// fmt.Println()
	return goodCount
}
