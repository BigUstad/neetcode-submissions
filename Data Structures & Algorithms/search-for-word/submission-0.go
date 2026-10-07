func exist(board [][]byte, word string) bool {
	if len(board) == 0 { return false }
	if len(word) == 0 { return true }
	rows, cols := len(board), len(board[0])
	var starts [][]int
	first := byte(word[0])
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if board[i][j] == first {
				starts = append(starts, []int{i,j})
			}
		}
	}
	path := make(map[[2]int]bool)
	var dfs func(r, c, i int) bool
	dfs = func(r, c, i int) bool {
		if i == len(word) {
			return true
		}
		coordinates := [2]int{r,c}
		// Out of bounds or already visited.
		// Return false
		if r < 0 || c < 0 || r >= rows || c >= cols ||
			board[r][c] != word[i] || path[coordinates] {
			return false
		}
		path[coordinates] = true
		res := dfs(r+1, c, i+1) ||
			   dfs(r-1, c, i+1) ||
			   dfs(r, c+1, i+1) ||
			   dfs(r, c-1, i+1)
		// This is the backtrack move
		delete(path, coordinates)
		return res
	}

	for _, s := range starts {
		if dfs(s[0], s[1], 0) {
			return true
		}
	}
	return false
}
