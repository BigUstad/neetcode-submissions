func orangesRotting(grid [][]int) int {
    curtime := 0
	if len(grid) == 0 { return curtime }
	// Adding cell number to queue
	// Cell number of currently rotting items
	var q [][2]int
	rows, cols := len(grid), len(grid[0])
    freshCount := 0
	for i, _ := range grid {
		for j, _ := range grid[i] {
			if grid[i][j] == 2 {
				q = append(q, [2]int{i, j})
			}
            if grid[i][j] == 1 {
				freshCount++
			}
		}
	}
    // fmt.Println("start curfresh = ", freshCount)
	if len(q) == 0 {
        if freshCount > 0 { return -1 }
        return curtime
    }
	// I'd like to start with minute 1 for curtime calculation
	for len(q) > 0 {
		n := len(q)
		for i := 0; i < n; i++ {
			r, c := q[0][0], q[0][1]
            // fmt.Println(r, c)
            // TODO: Queue slicing is not needed.
            // Assign front(q). Then assign back
            // Process front to back. Assign back index value to n
			q = q[1:]
            // fmt.Println(len(q))
			// ^^ remove q front
			// Set neighbors to be marked as rotten
            // If they are within bounds & if they are not empty.
            // Adding them if they are within bounds & have a fresh orange
            neighbors := [4][2]int {{r+1, c}, {r-1, c}, {r, c+1}, {r, c-1}}
            for _, n := range neighbors {
                if n[0] >= 0 && n[1] >= 0 &&
                    n[0] < rows && n[1] < cols &&
                    grid[n[0]][n[1]] == 1 {
                    grid[n[0]][n[1]] = 2
                    freshCount--
                    q = append(q, n)
                }
            }
            // fmt.Println(len(q))
		}
		if n > 0 {
            curtime++
		}
        // fmt.Println("cur fresh = ", freshCount)
		// fmt.Println(q, n, curtime)
	}
    // fmt.Println(curtime)
    // Shouldn't be less than 0
    if freshCount > 0 {
        return -1
    }
    // First level of BFS is all rotten.
    // Hence "- 1"
	return curtime-1
}
