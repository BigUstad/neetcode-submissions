/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func bfsHelper(root *TreeNode) []int {
	var res []int
	var prev *TreeNode
	queue := list.New()
	queue.PushBack(root)
	queue.PushBack(&TreeNode{Val: 1001, Left: nil, Right: nil,})
	for queue.Len() > 0 {
		curEle := queue.Front()
		cur := curEle.Value.(*TreeNode)
		queue.Remove(curEle)
		if cur.Val == 1001 {
			// fmt.Println(level)
			if prev != nil {
				res = append(res, prev.Val)
			}
			// Mark another level end
			// Except for the final node in bfs
			if queue.Len() > 0 {
				queue.PushBack(&TreeNode{Val: 1001, Left: nil, Right: nil,})
			}
			continue
		}
		// NA
		if cur.Left != nil {
			queue.PushBack(cur.Left)
		}
		if cur.Right != nil {
			queue.PushBack(cur.Right)
		}
		prev = cur
	}
    queue = nil
	return res
}

func rightSideView(root *TreeNode) []int {
    if root == nil {
		return []int{}
	}
	if root.Left == nil && root.Right == nil {
		return []int{root.Val}
	}
	return bfsHelper(root)
}
