/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type BSTIterator struct {
	stack *list.List
}

func Constructor(root *TreeNode) BSTIterator {
    itr := BSTIterator{
        stack: list.New(),
    }
    cur := root
	for cur != nil {
        itr.stack.PushFront(cur)
        cur = cur.Left
    }
    return itr
}

func (this *BSTIterator) Next() int {
    cur := this.stack.Remove(this.stack.Front()).(*TreeNode)
    val := cur.Val
    cur = cur.Right
	for cur != nil {
        this.stack.PushFront(cur)
        cur = cur.Left
    }
    return val
}

func (this *BSTIterator) HasNext() bool {
	return this.stack.Len() > 0
}

/**
 * Your BSTIterator object will be instantiated and called as such:
 * obj := Constructor(root)
 * param_1 := obj.Next()
 * param_2 := obj.HasNext()
 */
 