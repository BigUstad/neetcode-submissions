func longestConsecutive(nums []int) int {
	n := len(nums)
	if n == 0 { return 0 }
	m := make(map[int]bool)
	res := 0
	min := math.MaxInt
	for _, n := range nums {
		m[n] = true
		if n < min { min = n }
	}
	for k := range m {
		curCount := 0
		cur := k
		if _, found := m[k-1]; !found {
			for m[cur] {
				curCount++
				cur++
			}
			if curCount > res { res = curCount }
		}
	}
	return res
}
