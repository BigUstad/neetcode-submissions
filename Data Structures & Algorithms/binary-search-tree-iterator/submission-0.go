/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type BSTIterator struct {
    inorder []*TreeNode
    inIndex int
}

func getInorder(root *TreeNode) []*TreeNode {
    stack := list.New()
    var inorder []*TreeNode
    cur := root
    for stack.Len() > 0 || cur != nil {
        if cur != nil {
            stack.PushFront(cur)
            cur = cur.Left
        } else {
            if stack.Len() == 0 { continue }
            cur = stack.Remove(stack.Front()).(*TreeNode)
            // fmt.Print(cur.Val, "  ")
            inorder = append(inorder, cur)
            cur = cur.Right
        }
    }
    // fmt.Println()

    return inorder
}

func Constructor(root *TreeNode) BSTIterator {
    itr := BSTIterator{
        inorder: nil,
        inIndex: 0,
    }
    itr.inorder = getInorder(root)
    return itr
}


func (this *BSTIterator) Next() int {
    ret := this.inorder[this.inIndex].Val
    this.inIndex++
    if this.inIndex > len(this.inorder) {
        this.inIndex = len(this.inorder)
    }
    return ret
}


func (this *BSTIterator) HasNext() bool {
    return (this.inIndex < len(this.inorder))
}


/**
 * Your BSTIterator object will be instantiated and called as such:
 * obj := Constructor(root)
 * param_1 := obj.Next()
 * param_2 := obj.HasNext()
 */
 