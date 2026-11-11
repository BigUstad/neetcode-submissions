func scoreOfString(s string) int {
	i := 1
	last := len(s)
	sum := 0
	for i < last {
		diff := int(s[i]) - int(s[i-1])
		if diff < 0 {
			diff = -diff
		}
		// fmt.Println(s[i], s[i-1], diff)
		sum += diff
		i++
	}
	return sum
}
