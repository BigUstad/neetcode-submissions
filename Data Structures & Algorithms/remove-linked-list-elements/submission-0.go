/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func removeElements(head *ListNode, val int) *ListNode {
    newHead := &ListNode { Val: 51, Next: nil }
    var newListBuilder func(*ListNode, *ListNode)
    newListBuilder = func(prev, cur *ListNode) {
        if cur == nil {
            // fmt.Println()
            prev.Next = nil
            return
        }
        if cur.Val != val {
            // fmt.Print("\t", prev.Val, "->", cur.Val)
            prev.Next = cur
            newListBuilder(cur, cur.Next)
        } else {
            newListBuilder(prev, cur.Next)
        }
        
    }

    newListBuilder(newHead, head)
    return newHead.Next
}
