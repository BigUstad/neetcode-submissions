func countComponents(n int, edges [][]int) int {
	// Still need to learn disjoint set union or union find
    if n == 0 { return 0 }
    var buildAdjacencyMap func([][]int) map[int][]int
    var dfs func(int, int)
    // Build the adjacency or neighbor map
    // We'll use this to deduce where to go in DFS
    buildAdjacencyMap = func(edges [][]int) map[int][]int {
        adj := make(map[int][]int)
        for _, e := range edges {
            vs1 := adj[e[0]]
            vs2 := adj[e[1]]
            adj[e[0]] = append(vs1, e[1])
            adj[e[1]] = append(vs2, e[0])
        }
        return adj
    }
    // Mark a node as visited after visit.
    // The aim is to cover all nodes
    visited := make([]bool, n+1)
    adj := buildAdjacencyMap(edges)
    dfs = func(prev, cur int) {
        if visited[cur] {
            return
        }
        visited[cur] = true
        for _, j := range adj[cur] {
            // Don't revisit prev
            if j == prev {
                continue
            }
            dfs(cur, j)
        }
    }
    // resCount - return as resultant component count
    // The assumption is the lower number is always encountered first?
    resCount := 0
    for i := 0; i < n; i++ {
        if !visited[i] {
            resCount++
            dfs(-1, i)
        }
    }
    return resCount    
}
