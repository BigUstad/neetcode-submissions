import (
  pq "github.com/emirpasic/gods/queues/priorityqueue"
  // "github.com/emirpasic/gods/utils"
  "cmp"
)

func byWeight(a, b interface{}) int {
    ele1 := a.([2]int)
    ele2 := b.([2]int)
    return cmp.Compare(ele1[0], ele2[0])
}

func minCostConnectPoints(points [][]int) int {
    // Minimum Spanning Tree(MST) problem
    // Prims Algorithm. O(n^2 log n). MinHeap
    // Requires adjacency list
    // NOTE: Study Kruskal's algorithm, Boruvka's algorithm
    // NOTE: There is a greedy MST solution that is only O(n^2).
    //       Instead of building all edges (which would be too many), we compute distances on the fly and update the best known connection cost for each node.

    // 1. Build a graph where each pair of points has an edge with weight = Manhattan distance.
    n := len(points)
    if n == 0 { return 0 }
    // adj[i] = list of distances to all other points from i
    // The indices previous to i is ignored. Undirected edges, goes both ways.
    // - Why ignore it? Assumption is we add nodes 0 to n-1 in order, hence we'll add them to visited in that order(mostly).
    //   In the example1(leetcode), Each node will add the cost of the far-node to heap.
    //   At the end we will end up picking up the least costing value to reach that node
    // adj key - node index (0 to n). (points[0], points[1]) give the coordinates.
    // adj value - list of (cost, node index)
    // NOTE: This for loop can be run while running Prim's algorithm
    //       Separation is for understanding.
	// adj is just 1000 in length. It doesn't have to be a map
    // adj := make(map[int][][2]int)
	adj := make([][][2]int, 1000)
	for i, _ := range adj {
		adj[i] = make([][2]int, 0)
	}
    for i := 0; i < n; i++ {
        x1, y1 := points[i][0], points[i][1]
        // Going further from i. No need to backtrack
        for j := i+1; j < n; j++ {
            x2, y2 := points[j][0], points[j][1]
            // Manhattan distance calculation
            xdist := x2 - x1
            if xdist < 0 { xdist = -xdist }
            ydist := y2 - y1
            if ydist < 0 { ydist = -ydist }
            dist := xdist + ydist
            adj[i] = append(adj[i], [2]int{dist, j})
            adj[j] = append(adj[j], [2]int{dist, i})
        }
    }
    // res accumulates the sum of minimum cost from cur node
    res := 0
    // 2. Start MST from node 0.
    // 3. Push (0, 0) into a min-heap (cost to add starting node is 0).
    minHeap := pq.NewWith(byWeight)
    minHeap.Enqueue([2]int{0, 0})
    // 4. Maintain a visited set for nodes already in the MST.
    visited := make(map[int]bool)
    // 5. While visited size < N:
    for len(visited) < n {
        topi, e := minHeap.Dequeue()
        // Unreachable neighbor perhaps
        if !e {
            break 
        }
        top := topi.([2]int)
        // 	a. Pop the smallest (cost, node) from heap.
        cost, i := top[0], top[1]
        // 	b. If node is already visited, skip it.
        if visited[i] {
            continue
        }
        visited[i] = true
        res += cost
        // 	c. Otherwise
        //      Add cost to answer. Mark node as visited
        //      Push all edges (edgeCost, neighbor) to heap for unvisited neighbors.
        for _, nei := range adj[i] {
            // nei is [2]int cost, i
            if visited[nei[1]] {
                continue
            }
            minHeap.Enqueue(nei)            
        }
    }
    // 5. The accumulated cost is the minimum cost to connect all points.
    return res
}