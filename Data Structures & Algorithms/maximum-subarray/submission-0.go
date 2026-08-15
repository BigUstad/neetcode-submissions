func maxSubArray(nums []int) int {
    n := len(nums)
    if n == 0 { return 0 }
    if n == 1 { return nums[0] }
    cur := nums[0]
    max := nums[0]
    for i := 1; i < n ; i ++ {
        if cur < 0 {
            // Not worth it
            cur = nums[i]
        } else {
            cur += nums[i]
        }
        if cur > max {
            max = cur
        }
    }
    return max   
}
