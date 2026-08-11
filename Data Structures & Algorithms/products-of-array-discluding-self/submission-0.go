func productExceptSelf(nums []int) []int {
	// Prefix sum(product)
	n := len(nums)
	if n == 0 { return []int{}}
	prefixProds := make([]int, n)
	suffixProds := make([]int, n)
	prefixProds[0] = 1
	suffixProds[n - 1] = 1
	for i := 1 ; i < n ; i++ {
		prefixProds[i] = prefixProds[i - 1] * nums[i - 1]
	}
	for i := n - 2; i >= 0; i-- {
		suffixProds[i] = suffixProds[i+1] * nums[i+1]
	}
	for i := 0; i < n; i++ {
		prefixProds[i] *= suffixProds[i]
	}
	return prefixProds
}
