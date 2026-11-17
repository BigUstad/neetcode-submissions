func containsNearbyDuplicate(nums []int, k int) bool {
	sw := make(map[int]bool)
	i := 0
	ln := len(nums)
	// first sliding window length is k+1 because o-k
	for i < ln && i <= k {
		_, ok := sw[nums[i]]
		if ok { return true }
		sw[nums[i]] = true
		i++
	}
	// I guess I start from 0 but keep the sw length to k going forward
	for i < ln {
		delete(sw, nums[i-k-1])
		// fmt.Println(nums[i], sw)
		_, ok := sw[nums[i]]
		if ok { return true }
		sw[nums[i]] = true
		i++
	}
	return false
}

