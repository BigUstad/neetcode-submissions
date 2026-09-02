func longestPalindrome(s string) string {
    n := len(s)
	dp := make([][]bool, n)
	for i := range dp {
		dp[i] = make([]bool, n)
	}
	resIdx, resLen := 0, 0
	for i := n - 1; i >= 0; i-- {
		for j := i; j < n; j++ {
			if s[i] == s[j] && ((j - i) <= 2 || dp[i+1][j-1]) {
				dp[i][j] = true
				// j - i + 1 being the substring length
				subStrLth := (j - i + 1)
				if resLen < subStrLth {
					resLen = subStrLth
					resIdx = i
				}
			}
		}
	}

	return s[resIdx:resIdx+resLen]
}
