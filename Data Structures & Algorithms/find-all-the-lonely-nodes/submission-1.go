/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func getLonelyNodes(root *TreeNode) []int {
    q := []*TreeNode{root}
    res := make([]int, 0)
    head := 0
    for  head < len(q) {
        cur := q[head]
        head++
        if cur.Left != nil && cur.Right != nil {
            q = append(q, cur.Left)
            q = append(q, cur.Right)
            continue
        }
        if cur.Left == nil && cur.Right == nil {
            continue
        }
        if cur.Left != nil {
            q = append(q, cur.Left)
            res = append(res, cur.Left.Val)
        } else {
            q = append(q, cur.Right)
            res = append(res, cur.Right.Val)
        }
    }
    return res
}
