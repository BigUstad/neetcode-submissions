import ("slices")
func rob(nums []int) int {
    n := len(nums)
    if n == 0 {
        return 0
    }
    if n == 1 {
        return nums[0]
    }
    // DP Store
    steal := slices.Repeat([]int {-1}, n + 1)
    steal[n] = 0
    steal[n - 1] = nums[n - 1]
    for i := n - 2; i >= 0; i-- {
        steal[i] = max (
            steal[i+1],
            steal[i+2] + nums[i],
        )
    }
    // At index 0 we have the answer
    return steal[0]
}
