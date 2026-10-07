func minWindow(s string, t string) string {
    if len(t) == 0 || len(t) > len(s) { return "" }
	tm := make(map[rune]int)
	have, need := 0, len(t)
	i, j := -1, -1
	sz := math.MaxInt32
	n := len(s)
	for _, c := range t {
		tm[rune(c)]++
	}
	// fmt.Println("tm: ", tm)
	tmc := make(map[rune]int)
	l, r := 0, 0
	// Iterating on right pointer
	// Building the t map copy and comparing with t map
	for r < n {
		sr := rune(s[r])
        tmc[sr]++
		if have < need && tm[sr] > 0 && tmc[sr] <= tm[sr] {
			have++
		}
		// fmt.Println(need, have, s[l:r+1])
		// fmt.Println("tmc: ", tmc)
		for need == have {
			cur := r - l + 1
			if cur < sz {
				i, j = l, r
				sz = cur
				// fmt.Println("sz: ", sz)
			}
			// Moving the left pointer to find a shorter window
			sl := rune(s[l])
            l++
            tmc[sl]--
            // fmt.Println("sl: ", sl, "tmc.sl: ", tmc[sl])
			if tm[sl] > 0 && tmc[sl] < tm[sl] {
				have--
			}
		}
		r++
	}
	if sz == math.MaxInt32 || j == -1 {
		return ""
	}

	return s[i:j+1]
}
