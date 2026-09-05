func countSubstrings(s string) int {
    l := len(s)
	total := 0
	var expand = func (i, j int) int {
		count := 0
		// fmt.Fprint(os.Stdout, i, ",", j)
		for i >= 0 && j < l &&
				s[i] == s[j] {
			count++
			i--
			j++
		}
		// fmt.Fprintln(os.Stdout,",", count)
		return count
	}

	for i, _ := range s {
		// Count odd length palindrome
		total += expand(i, i)
		total += expand(i, i+1)
	}
	return total
}
