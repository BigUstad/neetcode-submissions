func lengthOfLongestSubstring(s string) int {
	l := len(s)
	if l == 0 { return 0 }
	r := []rune(s)
	// First Sliding window length
	m := make(map[rune]int)
	max := 0
	// fmt.Println(string(r[0:i]))
	// Search for a new window if it exists
	start := 0
	for j := 0; j < l; j++ {
		if x, ok := m[r[j]]; ok {
			for start <= x {
				delete(m, r[start])
				start++
			}			
		}
		m[r[j]] = j
		if max < len(m) { max = len(m) }
		// fmt.Println(m)
		// fmt.Println(string(r[start:j+1]))
	}
	return max
}
