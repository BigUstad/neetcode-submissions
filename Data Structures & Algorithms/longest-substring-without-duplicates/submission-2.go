func lengthOfLongestSubstring(s string) int {
	l := len(s)
	if l == 0 { return 0 }
	r := []rune(s)
	// First Sliding window length
	m := make(map[rune]int)
	i, max := 0, 0
	for _, c := range s {
		if _, ok := m[c];ok {
			break
		}
		// fmt.Print(c)
		// fmt.Print(",")
		// fmt.Println(i)
		m[c] = i
		i++
	}
	max = i
	// fmt.Println(string(r[0:i]))
	// Search for a new window if it exists
	start := 1
	delete(m, r[0])
	j := i
	for ; j < l; j++ {
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
