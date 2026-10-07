func numDecodings(s string) int {
    dp := make(map[int]int, len(s)+1)
    var dfs func(int) int
    dp[len(s)] = 1
    // Top-down approach for DP
    dfs = func(i int) int {
        // i has been already cached
        // i is the last case
        val, e := dp[i]
        if e {
            return val
        }
        // If not end of string
        // And string starts with 0
        // there are 0 more items
        if s[i] == '0' {
            return 0
        }
        // Sub problem becomes dfs(i+1)
        // also dfs(i+2)
        res := dfs(i+1)
        // 2 digit character. i+1 is inbounds
        // There is a second digit, hence there is a happy path.
        // If the digit is 1[0-9] or 2[0-6]
        if (i+1) < len(s) &&
            (s[i] == '1' ||
                (s[i] == '2' && s[i+1] <= '6')) {
            res += dfs(i+2)
        }
        dp[i] = res
        return res
    }
    // The Bottom-Up or the recursive solution is also simple
    // Iterate in reverse order. dp[i] = 0 for '0'. Else dp[i] = dp[i+1]. Add dp[i+2] for 2[0-6]
    return dfs(0)
}
