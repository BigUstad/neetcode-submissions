func merge(nums1 []int, m int, nums2 []int, n int) {
	last := (m+n) - 1
	n1 := m - 1
	n2 := n - 1
	for n1 >= 0 && n2 >= 0 {
		if nums1[n1] > nums2[n2] {
			nums1[last] = nums1[n1]
			n1--
		} else {
			nums1[last] = nums2[n2]
			n2--
		}
		last--
		// fmt.Println("1. ", nums1)
	}
	for n2 >= 0 {
		nums1[last] = nums2[n2]
		last--
		n2--
		// fmt.Println("2. ", nums1)
	}
}
