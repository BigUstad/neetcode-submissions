func canPartition(nums []int) bool {
    // First total sum even check.
    // If it isn't even, it cannot be halved to an integer
    totalSum := 0
    for _, n := range nums {
        totalSum += n
    }

    if totalSum%2 != 0 {
        return false
    }

    target := totalSum / 2

    // Change Bottom-Up DP solution
    // 0/1 subset sum DP
    // but bool array
    dp := make([]bool, target+1)
    dp[0] = true

    // O(nT) time. O(T) space
    for _, num := range nums {
        // Iterate backwards
        for sum := target; sum >= num; sum-- {
            // Having dp[sum] only makes sense if there is another subarray sum
            // without the num
            if dp[sum-num] {
                dp[sum] = true
            }
        }

        if dp[target] {
            return true
        }
    }

    return dp[target]
}
