/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Neighbors []*Node
 * }
 */

func cloneGraph(node *Node) *Node {
    if node == nil { return nil }
    // Serves as visited as well
	// No need for a marker to visit. Just return the clonenode record without creation
    oldToNew := make(map[*Node]*Node)
    var dfsHelper func(*Node)*Node
	// For BFS, I'd probably add the newly created node to a FIFO queue to add neighbors
    dfsHelper = func(node *Node) *Node {
        if node == nil { return nil }
        if n, exists := oldToNew[node]; exists {
            return n
        }
        cloneNode := &Node {
            Val: node.Val,
            Neighbors: nil,
        }
        oldToNew[node] = cloneNode
        for _, nei := range node.Neighbors {
            // NOTE: Recursive call is made here.
            // DFS & oldToNew as visited marker
            cloneNode.Neighbors = append(cloneNode.Neighbors, dfsHelper(nei))
        }
        return cloneNode
    }
    return dfsHelper(node) 
}
