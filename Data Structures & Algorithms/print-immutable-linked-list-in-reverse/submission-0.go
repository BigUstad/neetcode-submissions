/**
 * // This is the ImmutableListNode's API interface.
 * // You should not implement it, or speculate about its implementation.
 * type ImmutableListNode interface {
 *     PrintValue()
 *     GetNext() ImmutableListNode
 * }
 */

func printLinkedListInReverse(head ImmutableListNode) {
    var printHelper func(ImmutableListNode)
    printHelper = func(head ImmutableListNode) {
        // End condition
        if head != nil {
            printHelper(head.GetNext())
            head.PrintValue()
        }
    }
    printHelper(head)
}