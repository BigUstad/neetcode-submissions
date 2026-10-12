func maxAreaOfIsland(grid [][]int) int {
    maxArea := 0
	rows := len(grid)
	if rows == 0 {
		return maxArea
	}
	cols := len(grid[0])
	// visited map tracks the cells that have already been 'visited' & checked
	// First allocate a 2-D array with number of rows
	visited := make([][]bool, rows)
	// For each row, allocate a full row of 'cols' count of columns
	for x := 0; x < rows; x++ {
		visited[x] = make([]bool, cols)
	}
    // fmt.Println(visited)
	// DFS recursive function
	// arguments starting with coordinates i, j

	// NOTE: I don't want to rewrite the input grid to make it 'visited'
	// instead of visited[i][j] = true, we can indeed do grid[i][j] = 0 to mark it visited
	var dfs func(int, int) int
	dfs = func(i, j int) int {
		// Ending of dfs cases.
		// (i) Out of bounds
		// (ii) If the value is 0, we won't pursue
		// (iii) If already visited
		if i < 0 || j < 0 || i == rows || j == cols ||
			grid[i][j] == 0 || visited[i][j] {
			return 0
		}
		// mark the cell as visited
		visited[i][j] = true
		cur := 1 + dfs(i + 1, j) +
					dfs(i - 1, j) +
					dfs(i, j + 1) +
					dfs(i, j - 1)
		return cur
	}
	// Iterate over the grid
	// DFS implies, each point in grid will be visited only once.
	// If it is already visited, it will just return 0
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			maxArea = max(maxArea, dfs(i, j))
		}
	}

	return maxArea
}
