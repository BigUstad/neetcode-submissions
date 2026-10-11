func jump(nums []int) int {
	// The number of jumps required
    res := 0
	// reachable range left & right
	var left, right int
	last := len(nums) - 1
	// Iterate while trying to stretch right boundary till last
	for right < last {
		// fmt.Println("s ", left, ", ", right)
		farthest := 0
		for i := left; i <= right; i++ {
			farthest = max(farthest, i + nums[i])
		}
		// Start the search after current step
		left = right + 1
		right = farthest
		res++
		// fmt.Println("e ", left, ", ", right)
	}
	return res
}
