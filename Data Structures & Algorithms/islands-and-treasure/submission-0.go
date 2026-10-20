func islandsAndTreasure(grid [][]int) {
    // Define thresholds: rows & cols
    rows, cols := len(grid), len(grid[0])
    var addRoom func(int, int)
    // Visited map. This will make it capable of running BFS
    // from every empty room.
    // Start BFS once from all treasure (0 cells) at the same time
    // BFS spreads out in "waves" of distance 0, 1, 2
    // TODO: Study the intuitive BFS result in neetcode.
    // Intuitive & incremental
    visited := make([][]bool, rows)
    for r := 0; r < rows; r++ {
        visited[r] = make([]bool, cols)
    }
    q := make([][2]int, 0)

    addRoom = func(r, c int) {
        if r < 0 || r >= rows ||
            c < 0 || c >= cols ||
            grid[r][c] == -1 ||
            visited[r][c] {
            return
        }
        visited[r][c] = true
        q = append(q, [2]int{r, c})
    }
    // Starting the queue with all treasures
    // First time the wave reaches an empty cell, it must be
    // from the closest treasure. BFS guarantees the first visit
    // which would be the shortest distance in an unweighted grid
    for r := 0; r < rows; r++ {
        for c := 0; c < cols; c++ {
            if grid[r][c] == 0 {
                // Put all treasure cells into a queue
                q = append(q, [2]int{r, c})
                // Mark them as visited
                visited[r][c] = true
                // Walls (-1) are not added to queue
            }
        }
    }
    if len(q) == 0 { return }
    // Set dist to 0
    dist := 0
    // Iterate while queue is not empty
    for len(q) > 0 {
        // For each sell popped set grid[r][c] = dist (distance nearest to treasure)
        // 4 neighbors of current: mark visited & push to queue.
        // Only if within bounds & not a wall & not visited
        n := len(q)
        for i := 0; i < n; i++ {
            // current node
            r, c := q[0][0], q[0][1]
            q = q[1:]
            grid[r][c] = dist
            addRoom(r+1, c)
            addRoom(r-1, c)
            addRoom(r, c+1)
            addRoom(r, c-1)
        }
        // Increment distance by 1.
        // We've completed 1 layer of the room
        dist++
    }
}
