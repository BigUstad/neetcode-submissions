/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Children []*Node
 * }
 */

import ("slices")
func postorder(root *Node) []int {
    // Empty list if root is empty
    if root == nil { return []int{} }
    var stck []*Node
    var res []int
    // Stack for iterative traversal of the tree
    stck = append(stck, root)
    for len(stck) > 0 {
        // Pop from top of the stack
        cur := stck[len(stck)-1]
        stck = stck[:len(stck)-1]
        // Add cur val to result
        res = append(res, cur.Val)
        // fmt.Println(res)
        // Push all children of cur stack from start to end
        // Popped from end to start of children
        for _, c := range cur.Children {
            stck = append(stck, c)
        }
    }
    // Reverse the result list
    // In Post order traversal (iterative), if right child exists
    // traverse node from left child then move right
    slices.Reverse(res)
    return res
}
