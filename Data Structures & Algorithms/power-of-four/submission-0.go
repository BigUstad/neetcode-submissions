func isPowerOfFour(n int) bool {
	s := 1
	var determinePowFour func(s int) bool
	determinePowFour = func(s int) bool {
		if s == n {
			return true
		}
		if s > n {
			return false
		}
		ret := determinePowFour(s * 4)
		return ret
	}
	return determinePowFour(s)
}
