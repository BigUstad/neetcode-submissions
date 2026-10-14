func uniquePaths(m int, n int) int {
    // Bottom-most row.
    row := make([]int, n)
    for i := 0; i < n; i++ {
        row[i] = 1
    }
    // Each cell's value depends on cell to the right & cell below (2 moves that are allowed)
    // Space optimized solution is to keep only one row at a time for computation.
    // And a 'prev' row for previous row computation.
    // Other rows except for the last one
    for i := 0; i < m-1; i++ {
        newRow := make([]int, n)
        // Rightmost column is always one (same as bottom most)
        // Update from right to left
        // Current cell = next-right cell + below cell. Gives the count of options
        newRow[n-1] = 1
        // From right to left.
        // All columns except for right-most column
        for j := n-2; j >= 0; j-- {
            newRow[j] = newRow[j+1] + row[j]
        }
        row = newRow
    }
    // Now we have gotton to the first row
    return row[0]    
}
