func findRedundantConnection(edges [][]int) []int {
    // Kahn's algorithm understanding.
	// The adjacency map. Or the neighbor map here.
	// We'd need 2 nodes with degree-2?
	// Remove all degree-1 nodes by adding them to a queue.
	// Removing those nodes makes the nodes they are connected to the newer leaf nodes
	// The newer leaf nodes are removed as well.
	// Once the queue is empty, the nodes that are left form the cycle nodes
	
	// The adjacency list & the indegree of each node
	var buildAdjacency func() (map[int]map[int]bool, map[int]int)
	var processLeaves func()
	var getResult func() []int
	maxNode := math.MinInt
	buildAdjacency = func() (map[int]map[int]bool, map[int]int) {
		adj := make(map[int]map[int]bool)
		indegree := make(map[int]int)
		for _, e := range edges {
			indegree[e[0]]++
			indegree[e[1]]++
			if e[0] > maxNode || e[1] > maxNode {
				maxNode = max(e[0], e[1])
			}
			_, e0 := adj[e[0]]
			_, e1 := adj[e[1]]
			if !e0 { adj[e[0]] = make(map[int]bool) }
			if !e1 { adj[e[1]] = make(map[int]bool) }
			adj[e[0]][e[1]] = true
			adj[e[1]][e[0]] = true
		}

		return adj, indegree
	}

	adj, indegree := buildAdjacency()
	processLeaves = func() {
		var q []int
		// IMP NOTE: Traverse nodes from the end.
		// Problem statement states that result needs to be "last such edge"
		// If the edge has 2 nodes & their indegree is 2, they are the cycle nodes
		for k := maxNode; k > 0; k--  {
			if indegree[k] == 1 {
				q = append(q, k)
			}
		}
		for len(q) > 0 {
			node := q[0]
			q = q[1:]
			for k, _ := range adj[node] {
				// Deduct 1 from indegree as the node has been removed
				indegree[k]--
				// If indegree count is 1, "this" k node has become a leaf
				// And can be added to queue
				if indegree[k] == 1 {
					q = append(q, k)
				}
			}
		}
	}
	getResult = func() []int {
		var res []int
		for i := len(edges) - 1; i >= 0; i-- {
			x, y := edges[i][0], edges[i][1]
			if indegree[x] == 2 && indegree[y] == 2 {
				return []int{x, y}
			}
		}
		return res
	}

	processLeaves()
	// fmt.Println(indegree)
	return getResult()
}
