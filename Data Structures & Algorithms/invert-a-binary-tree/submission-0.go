/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func invertTree(root *TreeNode) *TreeNode {
    if root == nil ||
        root.Left == nil && root.Right == nil {
        return root
    }
    queue := list.New()
    queue.PushBack(root)
    for queue.Len() > 0 {
        cur := queue.Front().Value.(*TreeNode)
        queue.Remove(queue.Front())
        // swap
        cur.Left, cur.Right = cur.Right, cur.Left
        if cur.Left != nil { queue.PushBack(cur.Left) }
        if cur.Right != nil { queue.PushBack(cur.Right) }
    }
    return root
}
