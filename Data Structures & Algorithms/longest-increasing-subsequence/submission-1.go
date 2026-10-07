func lengthOfLIS(nums []int) int {
    // Dynamic Programming Bottom Up solution is O(n^2)
    // Need to refer book notes or Neetcode video
    // Simplistic solution: nums = [1, 2, 4, 3]
    // STart at the end. LIS[3] = 1 Only [3] .Compute backwards
    // 1 + x. Because each element is a sub array [1], [2], [4], [3].
    // So, it is current element plus subsequence length calculated from the end of array
    // LIS[0] = max(1, 1+LIS[3], 1+LIS[2], 1+LIS)

    // TODO: For a better solution O(n log n)
    // Lookup DP + Binary search (brief below). Or Segment tree
    // i. traverse original array.
    // ii. Construct dp array where you replace an existing element with current element in the right place
    // iii. Or append current element to end of dp array
    // - I think min can only be replaced if len of dp is 1. Else discard cur

    l := len(nums) - 1
    dp := make([]int, len(nums))
    dp[l] = 1
	// Let's actually keep track
	maxLISLn := 0
	if l == 0 {
		maxLISLn = 1
	}
    for i := l-1; i >= 0; i-- {
        // Start with 1
        dp[i] = 1
        for j := i+1; j <= l; j++ {
            if nums[i] < nums[j] {
                dp[i] = max(dp[i], 1 + dp[j])
            }
        }
		if dp[i] > maxLISLn {
			maxLISLn = dp[i]
		}
        // fmt.Println(dp[i:])        
    }
    return maxLISLn  
}
