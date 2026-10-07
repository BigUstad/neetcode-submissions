func validTree(n int, edges [][]int) bool {
    if n == 0 { return true }
    adj := make(map[int][]int)
    visited := make(map[int]bool)
    // Build adjacency map first.
    // As-we-go build is not easy for deductions
    // Adding adjacency map entries for both vertices
    for _, e := range edges {
        vs1 := adj[e[0]]
        vs2 := adj[e[1]]
        adj[e[0]] = append(vs1, e[1])
        adj[e[1]] = append(vs2, e[0])
    }
 
    // Need prev to avoid visiting 'from' node in dfs
    var dfs func(int, int) bool
    dfs = func(prev, cur int) bool {
        if visited[cur] {
            return false
        }
        visited[cur] = true
        for _, j := range adj[cur] {
            if j != prev && !dfs(cur, j) {
                return false
            }
        }
        return true
    }
    // DFS starts from 0 as 0 is guaranteed to be in the map
    // the length of visited node has to be n. Else it could be disconnected.
    // Running into same node during dfs implies loop & hence not a tree
    return dfs(-1, 0) && n == len(visited)
}