/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func reorderList(head *ListNode) {
	if head == nil || head.Next == nil {
		return
	}
    var reorderHelper func(*ListNode, *ListNode) *ListNode
	reorderHelper = func(root, cur *ListNode) *ListNode {
		if cur == nil {
			return root
		}
		root = reorderHelper(root, cur.Next)
		if root == nil {
			return nil
		}
		var temp *ListNode
		if root == cur || root.Next == cur {
			cur.Next = nil
		} else {
			temp = root.Next
			root.Next = cur
			cur.Next = temp
		}
		return temp
	}
	reorderHelper(head, head.Next)
}
