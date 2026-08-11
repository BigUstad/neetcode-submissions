/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func maxDepth(root *TreeNode) int {
    if root == nil { return 0 }
    if root.Left == nil && root.Right == nil { return 1 }
    depth := 1
    queue := list.New()
    queue.PushBack(root)
    queue.PushBack(&TreeNode{
                Val: -101,
                Left: nil,
                Right: nil,
            })

    for queue.Len() > 0 {
        curEle := queue.Front()
        cur := curEle.Value.(*TreeNode)
        queue.Remove(curEle)
        if cur.Val == -101 {
            if queue.Len() == 0 { break }
            depth++
            cur = nil
            queue.PushBack(&TreeNode{
                Val: -101,
                Left: nil,
                Right: nil,
            })
            continue
        }
        // end of previous level.
        // add nil signifying end of this level
        if cur.Left != nil {
            queue.PushBack(cur.Left)
        }
        if cur.Right != nil {
            queue.PushBack(cur.Right)
        }
    }
    return depth
}
