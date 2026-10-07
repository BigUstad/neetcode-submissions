func canJump(nums []int) bool {
	last := len(nums) - 1
    dp := make([]bool, last + 1)
	dp[last] = true
	// Compute dp array till
	// last - 1 index
	// index 0, if true, you can reach nums[last]
	for i := last - 1; i >= 0; i-- {
		// End index that can be reached from 'i'
		// Goal from i, if you will
		end := min(len(nums), 1 + i + nums[i])
		for j := i+1; j < end; j++ {
			if dp[j] {
				dp[i] = true
				break
			}
		}
	}
	// fmt.Println(dp)
	return dp[0]
}
