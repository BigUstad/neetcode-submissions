func partition(s string) [][]string {
	var res [][]string
	var dfs func(int)
	var cur []string
	var isPalindrome func(string, int, int) bool
	dfs = func(index int) {
		// fmt.Println(s[index:])
		if index >= len(s) {
			res = append(res, append([]string{}, cur...))
			return
		}
		for j := index; j < len(s); j++ {
			if isPalindrome(s, index, j) {
				cur = append(cur, s[index:j+1])
				dfs(j + 1)
				// Backtracking move
				cur = cur[:len(cur)-1]
			}
		}
	}

	isPalindrome = func(s string, i, j int) bool {
		for i <= j {
			if s[i] != s[j] {
				return false
			}
			i++
			j--
		}
		return true
	}
	dfs(0)
	return res
}
