import "slices"
func dailyTemperatures(temperatures []int) []int {
	n := len(temperatures)
	if n == 0 {
		return []int{}
	}
	if n == 1 {
		return []int{temperatures[0]}
	}
	if n == 2 {
		if temperatures[0] > temperatures[1] {
			return []int{temperatures[0]}
		}
		return []int{temperatures[1]}
	}
	warmer := make([]int, n)
	stack := make([]int, 0)
	for i, t := range temperatures {
		var sTop int
		sLen := len(stack)
        if sLen > 0 { sTop = stack[sLen - 1] }
		for sLen > 0 && t > temperatures[sTop] {
			// Next warmer day is x days away
			warmer[sTop] = (i - sTop)
			stack = slices.Delete(stack, sLen - 1, sLen)
			sLen = len(stack)
			if sLen > 0 { sTop = stack[sLen - 1] }
		}
		stack = append(stack, i)
	}
    return warmer
}
