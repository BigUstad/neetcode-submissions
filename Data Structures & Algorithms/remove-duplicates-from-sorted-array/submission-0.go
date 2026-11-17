func removeDuplicates(nums []int) int {
	i := 0
	j := 0
	ln := len(nums)
	if ln <= 1 {
		return ln
	}
	// Follower runner
	// Basically runner runs till nums[follower] != nums[runner]
	for j < ln {
		if nums[i] != nums[j] {
			i++
			nums[i] = nums[j]
		}
		j++
	}
	fmt.Println(nums)
	return i+1
}
