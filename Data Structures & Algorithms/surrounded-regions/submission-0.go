
func solve(board [][]byte) {
    rows := len(board)
    if rows == 0 { return }
    cols := len(board[0])
    keepOs := make([][]bool, rows)
    for i := 0; i < rows; i++ {
        keepOs[i] = make([]bool, cols)
    }
    // Trying time O(m * n) & space O(m * n)
    // Optimizing given space by only allocating bool
    // 1. Capture the unsorrounded regions. Mark those that are already O as true in keepOs 2D structure.
    //    This function to run only on the border cells
    var capture func(int, int)
    // 2. Capture the sorrounded regions. Do DFS to capture the "sorrounding" regions for those that are already O. Mark true as needed.
    var markSorrounded func()
    // 3. Last phase is convert O -> X
    var convert func()

    capture = func(r, c int) {
        if r < 0 || c < 0 || r == rows || c == cols {
            return
        }
        if board[r][c] != 'O' {
            return
        }
        if keepOs[r][c] {
            return
        }
        keepOs[r][c] = true
        // fmt.Println("Keeping O for ", r, c)
        adjacent := [4][2]int{{r+1, c}, {r-1, c}, {r, c+1}, {r, c-1}}
        for _, a := range adjacent {
            capture(a[0], a[1])
        }
    }

    markSorrounded = func() {
        lastRow := rows-1
        lastCol := cols-1
        for r, row := range board {
            for c, cell := range row {
                // fmt.Printf("%d %d %c %t\n", r, c, cell, (cell == 'O'))
                if cell == 'O' &&
                    (r == 0 || r == lastRow ||
                     c == 0 || c == lastCol) {
                    // fmt.Println("Capturing starting with ", r, c)
                    capture(r, c)
                }
            }
        }
    }

    convert = func() {
        for r, row := range board {
            for c, cell := range row {
                if cell == 'O' && !keepOs[r][c] {
                    board[r][c] = 'X'
                }
            }
        }
    }
    // Execute 2 steps in order
    markSorrounded()
    convert()
    // Well - done!! 
}
