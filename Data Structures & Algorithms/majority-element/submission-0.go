func majorityElement(nums []int) int {
    var count uint16
	var candidate int
	for _, n := range nums {
		if count == 0 {
			candidate = n
		}
		if candidate == n {
			count++
		} else {
			count--
		}
	}
	return candidate
}
