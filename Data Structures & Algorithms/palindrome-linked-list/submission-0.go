/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func isPalindrome(head *ListNode) bool {
	cur := head
	var rTraverse func(*ListNode) bool
	rTraverse = func(node *ListNode) bool {
		if node == nil {
			// Reached the end of the list
			// Should be same
			return true
		}
		if cur == nil {
			// Reached end of cur before node
			return false
		}
		if !rTraverse(node.Next) {
			return false
		}
		if cur.Val != node.Val {
			return false
		}
		cur = cur.Next
		// All Checks have passed
		return true
	}
	return rTraverse(head)
}
