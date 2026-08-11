/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func findNthNode(head *ListNode, n int) (runner *ListNode) {
    runner = head
    for n > 0 {
        n--
        runner = runner.Next
    }
    return
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
    var prevN *ListNode
    follower := head
    runner := findNthNode(head, n)
    for runner != nil {
        prevN = follower
        follower = follower.Next
        runner = runner.Next
    }
    if prevN == nil { return head.Next }

    // reassigning Next essentially deletes the node
    prevN.Next = follower.Next
    // For gc
    follower = nil
    return head
}
