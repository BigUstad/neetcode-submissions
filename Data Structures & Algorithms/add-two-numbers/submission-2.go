/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	// Sentinal node
	sentinal := &ListNode{Val: 111, Next: nil}

    var addHelper func(*ListNode, *ListNode, *ListNode, int)
	addHelper = func(prev, l1, l2 *ListNode, carry int) {
		// End case
		if l1 == nil && l2 == nil {
			if carry != 0 {
				cur := &ListNode{Val: carry, Next: nil}
				prev.Next = cur
			}
			return
		}
		if l1 == nil || l2 == nil {
            l := l1
            if l2 != nil { l = l2 }
            sum := l.Val + carry
            l.Val = sum % 10
            prev.Next = l
            addHelper(l, l.Next, nil, (sum / 10))
			return
		}
		sum := l1.Val + l2.Val + carry
		l2.Val = sum % 10
		prev.Next = l2
		addHelper(l2, l1.Next, l2.Next, (sum / 10))
	}
	addHelper(sentinal, l1, l2, 0)
	return sentinal.Next
}
