func wordBreak(s string, wordDict []string) bool {
    // Dynamic programming
    // Neetcode called it bottom up approach
	// TODO: Learn trie solution. Or Trie 😄
    l := len(s)
    dp := make([]bool, l + 1)
    dp[l] = true
    for i := l - 1; i >= 0; i-- {
        // Every single word in word dictionary
        // len(s) need to match
        for _, w := range wordDict {
            wl := len(w)
            // Length of word in wordDict is within bounds for the total length of string s
            // And they are equal in comparison
            if i + wl <= l && s[i:i+wl] == w {
                dp[i] = dp[i+wl]
            }
            if dp[i] {
                // No need to check each word for index i
                // We only need it to be true once
                break
            }
        }
    }
    // As we keep filling from the 'last' to 0
    // dp[0] will have the eventual answer
    return dp[0]
}
