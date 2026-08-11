/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

// import ("container/list")

func findMid(head *ListNode) (slow *ListNode) {
	slow = head
	fast := head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	return
}

func buildStack(mid *ListNode, stack *list.List) {
	for mid != nil {
		temp := mid.Next
		mid.Next = nil
		stack.PushFront(mid)
		mid = temp
	}
}

// 1. findMid of the list
// 2. From mid to last toss it into the stack
// 3. Take it from the stack and put it in the suggestested reorder
func reorderList(head *ListNode) {
    if head == nil || head.Next == nil ||
		head.Next.Next == nil {
		return
	}
	// 3 nodes or more, I guess
	mid := findMid(head)
	stack := list.New()
	buildStack(mid, stack)
	cur := head
	curStack := stack.Front()
	// curList := stack.Front().Value.(*ListNode)
	for curStack != nil && cur != nil {
		temp := cur.Next
		curNode := curStack.Value.(*ListNode)
		cur.Next = curNode
		curStack = curStack.Next()
		if curStack == nil {
			curNode.Next = nil
		} else {
			curNode.Next = temp
		}
		cur = temp
	}
	// Freeing for garbage collection I guess
	curStack = nil
	cur = nil
	mid = nil
}
