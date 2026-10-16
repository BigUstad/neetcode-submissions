func maxProfit(prices []int) int {
    n := len(prices)
	// take U forward
	// DP 39 solution
	dp := make([][2]int, n)
	for i := 0; i < n; i++ {
		dp[i] = [2]int{-1, -1}
	}
	var dfs func(int, int) int
	dfs = func(index, buy int) int {
		if index >= n {
			return 0
		}
		if dp[index][buy] >= 0 {
			return dp[index][buy]
		}
		if buy == 1 {
			// If you decide to buy, deduct prices[i].
			// Next day (index + 1), could be sell (buy 0)
			// If you don't decide to buy, price is 0.
			// with decision to buy (1) postponed to next day (index + 1)
			dp[index][buy] = max(
				(dfs(index + 1, 0) - prices[index]),
				(dfs(index + 1, 1) + 0),
				)
			return dp[index][buy]
		}
		// If "not" buy (sell or cooldown). Get the next day (index + 1)
		// price from the market.
		// Since buy is only allowed after cooldown (index + 2)
		// Or you could decide to not sell today (price is 0)
		// And only sell the next day (index + 1)
		dp[index][buy] = max(
			(dfs(index + 2, 1) + prices[index]),
			(dfs(index + 1, 0) + 0),
			)
		return dp[index][buy]
	}
	// buy == 1
	// NOTE: The first run had no dp.
	// Hence time limit exceeded.
	return dfs(0, 1)
}
