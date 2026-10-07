/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
    dummyHead := &ListNode{ Val: 111, Next: nil}
	var mergeHelper func(*ListNode, *ListNode, *ListNode)

	mergeHelper = func(head, list1, list2 *ListNode) {
		if list1 == nil {
			head.Next = list2
			return
		}
		if list2 == nil {
			head.Next = list1
			return
		}
		if list1.Val <= list2.Val {
			// fmt.Println(list1.Val)
			head.Next = list1
			mergeHelper(head.Next, list1.Next, list2)
		} else {
			// fmt.Println(list2.Val)
			head.Next = list2
			mergeHelper(head.Next, list1, list2.Next)
		}
	}

	mergeHelper(dummyHead, list1, list2)
	return dummyHead.Next
}
