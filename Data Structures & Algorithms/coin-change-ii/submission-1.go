func change(amount int, coins []int) int {
    n := len(coins)
    dp := make([][]uint32, n+1)
    // Fine with allocation value being 0
    // Refer to Back To Back SWE solution
    // Time & Space - O(amount*n)
    for i := 0; i <= n; i++ {
        // Column 0 is amount 0
        dp[i] = make([]uint32, amount+1)
    }
    var tabulate func() int
    tabulate = func() int {
        // Top row is all 0. We need it for the rest of the calculations
        for j := 0; j <= amount; j++ {
            // empty [] is 1 for amount 0
            if j == 0 {
                dp[0][j] = 1
                continue
            }
            // empty [] is 0 for amount 1 to value 'amount'
            dp[0][j] = 0
        }
        // i = 0 is covered above
        for i := 1; i <= n; i++ {
            for j := 0; j <= amount; j++ {
                var addendum uint32
                if (j-coins[i-1]) >= 0 {
                    addendum = dp[i][j-coins[i-1]]
                }
                dp[i][j] = dp[i-1][j] +
                            addendum
            }
        }
        return int(dp[n][amount])
    }
    return tabulate()
}
