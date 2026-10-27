import (
  pq "github.com/emirpasic/gods/queues/priorityqueue"
  "cmp"
  // "github.com/emirpasic/gods/utils"
)

type MinNode struct {
    node int
    weight int
}

func byWeight(a, b interface{}) int {
    return cmp.Compare(a.(MinNode).weight, b.(MinNode).weight)
}

func networkDelayTime(times [][]int, n int, k int) int {
    // edges := make(map[int]map[int]int)
    // TODO: Study - Floyd Warshall Algorithm, Bellman Ford Algorithm, Shortest Path Faster Algorithm
    // 1. Build an adjacency list from times.
    // NOTE: 2D array could have also sufficed with high default values
    // Trying Neetcode with 2D array
	edges := make([][]int, n+1)
	for i := 0; i <=n; i++ {
		edges[i] = make([]int, n+1)
		for j := 0; j <= n; j++ {
			edges[i][j] = math.MinInt
		}
	}
    for _, t := range times {
        u, v, w := t[0], t[1], t[2]
        edges[u][v] = w
    }
    // 2. Use a min-heap initialized with (0, k) (time, node).
    minHeap := pq.NewWith(byWeight)
    minHeap.Enqueue(MinNode{
        node: k,
        weight: 0,
    })
    // 3. Maintain a visited set to avoid reprocessing nodes.
    // visited := make(map[int]bool)
	visited := make([]bool, n+1)
	// Setting visited[0] to true as it is not needed here.
	visited[0] = true
    t := 0
    // 4. Iterate While the heap is not empty.
    for !minHeap.Empty() {
        // Pop the node with the smallest time.
        // fmt.Println(minHeap)
        // fmt.Println(visited)
        i, _ := minHeap.Dequeue()
        curNode := i.(MinNode)
        // If already visited, skip.
        if visited[curNode.node] {
            continue
        }
        // Mark it visited and update the current time. We don't want to keep visiting a node's neighbors
        visited[curNode.node] = true
        // Push all unvisited neighbors with updated times into the heap.
        t = max(t, curNode.weight)
        // BFS portion of it. Visiting all the neighbors.
        for k, v := range edges[curNode.node] {
			// Exclude if weight is not provided
			if v == math.MinInt {
				continue
			}
			// fmt.Println("k=", k, " v=", v)
            if !visited[k] {
                minHeap.Enqueue(MinNode{
                    node: k,
                    weight: (curNode.weight + v),
                })
            }
        }
    }
    // 5. If all n nodes are visited, return the maximum time used.
    // 6. Otherwise, return -1 (some nodes are unreachable).
	for _, v := range visited {
		if !v { return -1 }
	}
    return t
}