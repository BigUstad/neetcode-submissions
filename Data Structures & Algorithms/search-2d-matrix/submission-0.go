func searchMatrix(matrix [][]int, target int) bool {
	// Edge case: Empty matrix
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return false
	}

	numRows := len(matrix)
	numCols := len(matrix[0])

	// Treat the entire 2D space as a single 1D space
	low := 0
	high := (numRows * numCols) - 1

	for low <= high {
		mid := low + (high-low)/2

		// Map 1D virtual index back to physical 2D matrix coordinates
		currentRow := mid / numCols
		currentCol := mid % numCols
		midValue := matrix[currentRow][currentCol]

		if midValue == target {
			return true
		}

		if midValue < target {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}

	return false
}
