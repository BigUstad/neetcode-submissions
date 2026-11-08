/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func getNodeCount(head *ListNode) int {
    cnt := 0
    for head != nil {
        cnt++
        head = head.Next
    }
    return cnt
}

func reverseKGroup(head *ListNode, k int) *ListNode {
    if head == nil || head.Next == nil { return head }
    var newHead, curLast *ListNode
    // input - first.
    // returns - last, last.Next
    // the last.Next becomes the "next" first.
    // The 
    var getNextLast func(*ListNode) (*ListNode, bool)
    var reverseHelper func(*ListNode, *ListNode, *ListNode, *ListNode)
    var q []*ListNode
    getNextLast = func(first *ListNode) (*ListNode, bool) {
        if first == nil {
            return nil, false
        }
        // You already got the first
        curK := k-1
        last := first
        for curK > 0 && last != nil && last.Next != nil {
            last = last.Next
            curK--
        }
        // if first != nil { fmt.Println("1. first = ", first.Val) }
        // if last != nil { fmt.Println("1. last = ", last.Val) }
        if curK == 0 {
            return last, true
        }
        return last, false
    }

    reverseHelper = func(prev, cur, next, stop *ListNode) {
        // End of recursion
        if next == nil || cur == stop {
            return
        }
        temp := next.Next
        next.Next = cur
        cur.Next = prev
        reverseHelper(cur, next, temp, stop)
    }
    // var printHelper = func(first, last *ListNode) {
    //     for first != nil {
    //         fmt.Print(" - cur = ", first.Val)
    //         if first == last {
    //             break
    //         }
    //         first = first.Next
    //     }
    //     fmt.Println()
    // }

    first, last := head, head
    reverse := false
    if listCount := getNodeCount(head); listCount < k {
        return head
    } else if listCount == k {
        last, reverse = getNextLast(first)
        reverseHelper(nil, first, first.Next, last)
        return last        
    }
    // If listCount is <= k.
    // Just reverse the existing list & return the newHead (last)
    // if listCount <= k {
    //     last, reverse = getNextLast(first)
    //     reverseHelper(nil, first, first.Next, last)
    //     return last
    // }

    prevFirst := first
    for first != nil && last != nil && last.Next != nil {
        // prevFirst (first of last batch), curLast (last of current batch)
        // need to be linked to each other for the previous list to be preserved
        if curLast != nil && prevFirst != nil {
            // fmt.Println("2. prevFirst = ", prevFirst.Val, " curLast = ", curLast.Val)
            prevFirst.Next = curLast
        }
        last, reverse = getNextLast(first)
        temp := last.Next
        // if first != nil { fmt.Println("2. first = ", first.Val) }
        // if last != nil { fmt.Println("2. last = ", last.Val) }
        // fmt.Println()
        if reverse {
            q = append(q, first)
            q = append(q, last)
            reverseHelper(nil, first, first.Next, last)
        } else {
            q = append(q, last)
            q = append(q, first)
        }
        // printHelper(last, first)
        first = temp
    }
    if prevFirst != nil {
        prevFirst.Next = nil
    }
    // for _, qe := range q {
    //     fmt.Print("\t", qe.Val)
    // }
    // fmt.Println()
    var prevf *ListNode
    for len(q) > 0 {
        f := q[0]
        q = q[1:]
        l := q[0]
        q = q[1:]
        if newHead == nil {
            newHead = l
            prevf = f
            continue
        }
        if prevf != nil {
            prevf.Next = l
        }
        prevf = f
    }
    if prevf != nil {
        // fmt.Print("\t prevf ", prevf.Val)
        prevf.Next = nil
    }
    // fmt.Println()

    return newHead
}
