/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func reverseList(head *ListNode) *ListNode {
    // 0, 1, 2 nodes
    if head == nil || head.Next == nil {
        return head
    }
    if head.Next.Next == nil {
        newHead := head.Next
        head.Next.Next = head
        head.Next = nil
        return newHead
    }
    // 3 nodes or more
    var cur, prev, next *ListNode
    cur = head
    prev = nil
    next = head.Next
    // | -> | -> |
    // | <- | <- |
    for cur != nil {
        cur.Next = prev
        prev = cur
        cur = next
        if next != nil {
            next = next.Next
        }
    }
    return prev
}
