func characterReplacement(s string, k int) int {
	l := len(s)
	if l == 0 { return 0 }
    // Sliding window map
    m := make([]int, 26)
    maxCount := 0
    maxFreq := 0
    start := 0
    for j := start; j < l; j++ {
        i := s[j] - 'A'
        m[i]++
        maxFreq = max(maxFreq, m[i])
        check := ((j + 1 - start - maxFreq) <= k)
        // fmt.Print(check)
        // fmt.Print(", ")
        // fmt.Println(j)
        // Advance start
        if !check {
            atStart := s[start] - 'A'
            m[atStart]--
            start++
        }
        // fmt.Println(s[start:j+1])
        maxCount = j + 1 - start
    }

    return maxCount
}
