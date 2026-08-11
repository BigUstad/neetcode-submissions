/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func hasCycle(head *ListNode) bool {
    follower := head
    runner := head
    for follower != nil && runner != nil {
        if runner.Next != nil && runner.Next.Next != nil {
            runner = runner.Next.Next

        } else {
            return false
        }
        if follower.Next != nil {
            follower = follower.Next
        }
        if follower == runner {
            return true
        }
    }
    if follower == nil || runner == nil {
        return false
    }
    return true    
}
