/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func reverseList(head *ListNode) *ListNode {
    if head == nil || head.Next == nil { return head }
    if head.Next.Next == nil {
        temp := head.Next
        head.Next.Next = head
        head.Next = nil
        return temp
    }
    var prev, newHead *ListNode
    cur := head
    var reverseHelper func(*ListNode, *ListNode, *ListNode)
    reverseHelper = func(prev, cur, next *ListNode) {
        // End of recursion
        if next == nil {
            newHead = cur
            return
        }
        // fmt.Println(cur.Val)
        temp := next.Next
        next.Next = cur
        cur.Next = prev
        reverseHelper(cur, next, temp)
    }
    reverseHelper(prev, cur, cur.Next)
    return newHead
}
