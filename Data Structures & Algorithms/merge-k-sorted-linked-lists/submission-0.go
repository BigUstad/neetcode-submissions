/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

import (
  pq "github.com/emirpasic/gods/queues/priorityqueue"
  "github.com/emirpasic/gods/utils"
)

// Comparator function (sort by element's priority value in descending order)
func byListNode(a, b interface{}) int {
    priorityA := a.(*ListNode).Val
    priorityB := b.(*ListNode).Val
    return -utils.IntComparator(priorityA, priorityB) // "-" descending order
}

func mergeKLists(lists []*ListNode) *ListNode {
    queue := pq.NewWith(byListNode)
	insertionDone := false
	// I need a way to maintain a maximum number of nodes in the pq
	// That is stop or pause after a certain number is reached.
	// But if the heap is popped and put in the resulting list, we'll need to rewind t.
	// If we popped too early, that'd be wrong
	for !insertionDone{
		insertionDone = true
		for i, l := range lists {
			if l != nil {
				insertionDone = false
			} else {
				continue
			}
			queue.Enqueue(l)
			lists[i] = l.Next
		}
		// fmt.Println("pq size = ", queue.Size())
	}
	var next *ListNode
	for !queue.Empty() {
		ni, e := queue.Dequeue()
		if !e { break; }
		n := ni.(*ListNode)
		n.Next = next
		next = n
	}
	// Return last next as front
	return next
}
