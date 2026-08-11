/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
    var dummyHead, prev, cur, cur1, cur2 *ListNode
	dummyHead = &ListNode {
		Val: math.MinInt,
		Next: nil,
	}
	prev, cur = dummyHead, dummyHead
	cur1 = list1
	cur2 = list2
	for cur1 != nil && cur2 != nil {
		if cur1.Val <= cur2.Val {
			cur = cur1
			cur1 = cur1.Next
		} else {
			cur = cur2
			cur2 = cur2.Next
		}
		prev.Next = cur
		prev = cur
	}
	if cur1 != nil {
		prev.Next = cur1
	} else if cur2 != nil {
		prev.Next = cur2
	}
	cur = dummyHead.Next
	dummyHead = nil
	return cur
}
