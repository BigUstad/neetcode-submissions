func myPow(x float64, n int) float64 {
    var helper func(float64, int) float64
	helper = func(x float64, n int) float64 {
		if x == 0 { return 0 }
		if n == 0 { return 1 }
		res := helper(x*x, n/2)
		if n%2 != 0 {
			return x * res
		}
		return res
	}
	res := helper(x, int(math.Abs(float64(n))))
	if n < 0 { return 1/res}
	return res
}
