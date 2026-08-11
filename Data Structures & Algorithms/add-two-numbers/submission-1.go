/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var newHead, l3Cur *ListNode
    if l1 == nil && l2 == nil {
		return newHead
	}
	carryOver := 0
	for l1 != nil || l2 != nil || carryOver != 0 {
		l1Val := 0
		if l1 != nil { l1Val = l1.Val }
		l2Val := 0
		if l2 != nil { l2Val = l2.Val }
		l3Val := l1Val + l2Val + carryOver
		if l3Val >= 10 {
			carryOver = l3Val / 10
			l3Val = l3Val % 10
		} else { carryOver = 0 }
		if newHead == nil {
			newHead = &ListNode{ Val: l3Val, Next: nil,}
			l3Cur = newHead
		} else {
			l3Cur.Next = &ListNode{ Val: l3Val, Next: nil,}
			l3Cur = l3Cur.Next
		}
		if l1 != nil { l1 = l1.Next }
		if l2 != nil { l2 = l2.Next }
	}
	return newHead
}

