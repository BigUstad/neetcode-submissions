import ("slices")
func intersection(nums1 []int, nums2 []int) []int {
    var res []int
	small := nums1
	big := nums2
	if len(nums2) > len(nums1) {
		small = nums2
		big = nums1
	}
	slices.Sort(big)
	m := make(map[int]bool)
	for _, n := range small {
		_, found := slices.BinarySearch(big, n)
		if found && !m[n] {
			res = append(res, n)
			m[n] = true
		}
	}
	return res
}
