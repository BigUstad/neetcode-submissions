type DisjointSet struct {
	parent []int
	// Keeps tree flat, balanced.
	rank []int 
}

func NewDisjointSet(n int) *DisjointSet {
	parent := make([]int, n)
	rank := make([]int, n)
	// Each node initialized as its own parent
	for i, _ := range parent {
		parent[i] = i
	}
	return &DisjointSet{
		parent: parent,
		rank: rank,
	}
}

// Find function is to find the root leader for given 'i'
func(d *DisjointSet) Find(i int) int {
	if d.parent[i] == i {
		return i
	}
	// Path Compression: Memorize the boss directly to skip steps next time
	d.parent[i] = d.Find(d.parent[i])
	return d.parent[i]
}

// Union merges groups containing i & j
func(d *DisjointSet) Union(i, j int) bool {
	rootI := d.Find(i)
	rootJ := d.Find(j)
	// Both i, j part of the same group.
	// And hence a cycle exists
	if rootI == rootJ {
		return false
	}
	// Keep the balance.
	// Hang the shorter tree under the taller tree
	if d.rank[rootI] < d.rank[rootJ] {
		d.parent[rootI] = rootJ
	} else if d.rank[rootI] > d.rank[rootJ] {
		d.parent[rootJ] = rootI
	} else {
		d.parent[rootJ] = rootI
		d.rank[rootI]++
	}
	return true
}

func findCircleNum(isConnected [][]int) int {
	n := len(isConnected)
	if n == 0 { return 0 }
	// 1. Initialize a NewDisjointSet(n) where n is the number of cities.
	ds := NewDisjointSet(n)
	// 2. Set a variable provinces = n.
	provinces := n
	//3. Loop through the matrix. Whenever you find an edge isConnected[i][j] == 1:
	for i, row := range isConnected {
		for j, _ := range row {
			if isConnected[i][j] == 1 {
			// • Call d.Union(i, j).
			// • If Union returns true,
			//	 it means two previously disconnected cities just joined families!
			//	 Decrement provinces--.
				if ds.Union(i, j) {
					provinces--
				}
			}
		}
	}
	// Return count of provinces you figured out.
	return provinces
}
