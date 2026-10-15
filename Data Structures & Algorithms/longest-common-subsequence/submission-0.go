func longestCommonSubsequence(text1 string, text2 string) int {
    // Fill 2D iteratively from the end
    // dp[i][j] represents the LCS length of substrings
    // text1[i:] & text2[j:]. 
    // Processed in reverse order.
    // dp[i][j] is computed based on dp[i+1][j+1], dp[i+1][j], dp[i][j+1]

    // Create 2D array dp. Size m+1 & n+1. By default it is initialized to 0
    res := 0
    m, n := len(text1), len(text2)
    dp := make([][]uint16, m+1)
    for i := 0; i <= m; i++ {
        dp[i] = make([]uint16, n+1)
    }
    // The right-most column is all 0s.
    // The bottom-most column is all 0s.
    // They are used for counting specifically for
    // cells matching string char borders (at len(text1)-1, len(text2)-1. Leftwards & Upwards)
    // Leetcode: DP Approach 3 2-D pic
    // Neetcode Bottom-up DP
    for i := m-1; i >= 0; i-- {
        for j := n-1; j >= 0; j-- {
            // If text1[i] == text2[j], increment count from diagonal match
            if text1[i] == text2[j] {
                dp[i][j] = 1 + dp[i+1][j+1]
            } else {
                // If no match, dp[i][j] takes up match count from cell directly below or right cell beside
                dp[i][j] = max(dp[i+1][j], dp[i][j+1])
            }
        }
    }

    res = int(dp[0][0])
    return res
}
