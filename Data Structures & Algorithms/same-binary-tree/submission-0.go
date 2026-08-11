/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func bfsHelper(root *TreeNode) string {
    var bfsOutput string
    queue := list.New()
    queue.PushBack(root)
    for queue.Len() > 0 {
        curEle := queue.Front()
        cur := curEle.Value.(*TreeNode)
        queue.Remove(curEle)
        if cur.Val == 10001 {
            bfsOutput = strings.Join([]string{bfsOutput, "nullptr"}, ",")
            curEle = nil // for gc
            continue
        }
        bfsOutput = strings.Join([]string{bfsOutput, strconv.Itoa(cur.Val)}, ",")
        if cur.Left != nil {
            queue.PushBack(cur.Left)
        } else {
            queue.PushBack(&TreeNode{
                Val: 10001, Left: nil, Right: nil,
            })
        }
        if cur.Right != nil {
            queue.PushBack(cur.Right)
        } else {
            queue.PushBack(&TreeNode{
                Val: 10001, Left: nil, Right: nil,
            })
        }
    }
    return bfsOutput
}

func isSameTree(p *TreeNode, q *TreeNode) bool {
    if p == nil && q == nil {
        return true
    }
    if p == nil || q == nil {
        return false
    }
    pOutput := bfsHelper(p)
    qOutput := bfsHelper(q)
    return pOutput == qOutput
}
