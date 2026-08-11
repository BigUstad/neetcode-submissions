/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
import("slices")

func bfsHelper(root *TreeNode, level []int, levels *[][]int) {
	queue := list.New()
	queue.PushBack(root)
	queue.PushBack(&TreeNode{Val: 1001, Left: nil, Right: nil,})
	for queue.Len() > 0 {
		curEle := queue.Front()
		cur := curEle.Value.(*TreeNode)
		queue.Remove(curEle)
		if cur.Val == 1001 {
			// fmt.Println(level)
			*levels = append(*levels, slices.Clone(level))
			level = level[:0]
			cur = nil // gc
			// Mark another level end
			// Except for the final node in bfs
			if queue.Len() > 0 {
				queue.PushBack(&TreeNode{Val: 1001, Left: nil, Right: nil,})
			}
			continue
		}
		level = append(level, cur.Val)
		if cur.Left != nil {
			queue.PushBack(cur.Left)
		}
		if cur.Right != nil {
			queue.PushBack(cur.Right)
		}
	}
}

func levelOrder(root *TreeNode) [][]int {
    if root == nil {
		return [][]int{}
	}
	if root.Left == nil && root.Right == nil {
		return [][]int{{root.Val}}
	}
	var levels [][]int
	bfsHelper(root, []int{} ,&levels)
	return levels
}
