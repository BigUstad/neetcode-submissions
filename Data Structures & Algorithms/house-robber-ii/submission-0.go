func rob(nums []int) int {
    n := len(nums)
    if n == 0 { return 0 }
    if n == 1 { return nums[0]}

    helper := func(start, end int) int {
        prevMax := 0
        curMax := 0
        //numIndex := end
        for i := start; i <= end; i++ {
            atCur := prevMax+nums[i]
            prevMax = curMax
            curMax = max(curMax, atCur)
        }
        return curMax
    }
    // fmt.Println("first:")
    first := helper(0, n-2)
    // fmt.Println("Second:")
    second := helper(1, n-1)
    return max(first, second)
}
