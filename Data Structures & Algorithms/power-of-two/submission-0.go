func isPowerOfTwo(n int) bool {
	s := 1
	var determinePowTwo func(s int) bool
	determinePowTwo = func(s int) bool {
		if s == n {
			return true
		}
		if s > n {
			return false
		}
		ret := determinePowTwo(s * 2)
		return ret
	}
	return determinePowTwo(s)
}
