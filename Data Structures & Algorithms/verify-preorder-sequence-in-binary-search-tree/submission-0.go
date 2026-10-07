import ("slices")
type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func verifyPreorder(preorder []int) bool {
    inorder := slices.Clone(preorder)
    slices.Sort(inorder)
    inorderIndices := make(map[int]int)
    for i, e := range inorder {
        inorderIndices[e] = i
    }
    preorderIdx := 0
    // last index
    n := len(preorder) - 1
    var buildTreeHelper func(int, int, int, int) (*TreeNode, bool)
    buildTreeHelper = func(l, r, ll, rr int) (*TreeNode, bool) {
        if l > r || preorderIdx > n {
            return nil, true
        }
        // Get the root element from preorder (NLR)
        // Find root in inorder indices map
        // Every node after inorder index in preorder array is a root of a subtree
        //      Either it is a root node of a left or right subtree
        //      Or it is a leaf node with no children
        var lRet, rRet bool
        preorderVal := preorder[preorderIdx]
        preorderIdx++
        root := &TreeNode{Val: preorderVal, Left: nil, Right: nil}
        // For 'root', left subtree are elements in inorder range
        root.Left, lRet = buildTreeHelper(l, inorderIndices[preorderVal]-1, inorder[l], preorderVal)
        root.Right, rRet = buildTreeHelper(inorderIndices[preorderVal]+1, r, preorderVal, inorder[r])
        // fmt.Println((preorderVal > ll && preorderVal < rr), preorderVal, ll, rr)

        return root, lRet && rRet && (preorderVal >= ll && preorderVal <= rr)
    }
    // So, the island of isolation will be GC'd
    _, complete := buildTreeHelper(0, n, inorder[0], inorder[n])
    return complete
}
