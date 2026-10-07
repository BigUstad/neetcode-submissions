func pacificAtlantic(heights [][]int) [][]int {
	// This is the DFS solution.
	// TODO: Study the BFS solution. It uses queue as you'd expect
    rows, cols := len(heights), len(heights[0])
    // Stand in for hashset.
    // This here is an optimization
    // instead of having map[[2]int]bool -> row/col coordinates already visited.
    // I'll have a 2D array set. Set it true if already visited.
    // pac for pacific & atl for atlantic
    pac := make([][]bool, rows)
    atl := make([][]bool, rows)
    for i := 0; i < rows; i++ {
        pac[i] = make([]bool, cols)
        atl[i] = make([]bool, cols)
    }
    var dfs func(int, int, int, [][]bool)
    // Same dfs for atlantic & pacific
    dfs = func(r, c, prevHeight int, visit [][]bool) {
        // Cant continue for:
        // Out of bounds
        // already visited
        // Current height is less than previous height
        // NOTE: We start from top & bottom. Left & right.
        //       As we move inwards, height has to increase or stay the same for the water to flow
        if r < 0 || c < 0 ||
            r == rows || c == cols ||
            visit[r][c] || heights[r][c] < prevHeight {
            return
        }
        //fmt.Println(r, ",", c, " h ", heights[r][c], " p ", prevHeight)
        visit[r][c] = true
        dfs(r+1, c, heights[r][c], visit)
        dfs(r-1, c, heights[r][c], visit)
        dfs(r, c+1, heights[r][c], visit)
        dfs(r, c-1, heights[r][c], visit)
    }
    // Let's go through every column in first row.
    // In each element of the row. Let's run dfs
    // First row is always pacific. Let's cover it.
    // Water goes from equal or "greater" value. cur < previous
    // In the same iteration, let's check last row for atlantic
    for j := 0; j < cols; j++ {
        dfs(0, j, heights[0][j], pac)
        dfs(rows-1, j, heights[rows-1][j], atl)
    }
    // left column is pacific & right column is atlantic
    for i := 0; i < rows; i++ {
        dfs(i, 0, heights[i][0], pac)
        dfs(i, cols-1, heights[i][cols-1], atl)
    }
    // Last step is brute forcing the discovery
    var res [][]int
    for i := 0; i < rows; i++ {
        for j := 0; j < cols; j++ {
            if pac[i][j] && atl[i][j] {
                res = append(res, []int{i, j})
            }
        }
    }
    return res
}
