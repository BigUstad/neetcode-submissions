func kthGrammar(n int, k int) int {
    var dfs func(rootVal int) int
    dfs = func(rootVal int) int {
        if n <= 1 { return rootVal }
        totalNodes := int(math.Pow(float64(2), float64(n-1)))
        nextRootVal := 1
        n--
        // So, k belongs in the right subtree
        if k > (totalNodes/2) {
            if rootVal == 1 {
                nextRootVal = 0
            }
            k -= totalNodes/2
        } else {
            // k belongs in the left subtree
            // if rootVal == 1, then nextRootVal = 1
            if rootVal == 0 {
                nextRootVal = 0
            }
        }
        return dfs(nextRootVal)
	}
	return dfs(0)
}
