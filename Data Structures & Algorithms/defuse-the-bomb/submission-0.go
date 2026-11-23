func decrypt(code []int, k int) []int {
	ln := len(code)
	res := make([]int, ln)
	if k == 0 {
		return res
	}
	start := 1
	end := k
	sum := 0
	if k < 0 {
		start = ln - (-k)
		end = ln - 1
	}
	// First window
	for i := start; i <= end; i++ {
		sum += code[i]
	}
	// Continue with window
	for i := 0; i < ln; i++ {
		res[i] = sum
		sum -= code[start % ln]
		end++
		sum += code[end % ln]
		start++
	}
	return res
}
