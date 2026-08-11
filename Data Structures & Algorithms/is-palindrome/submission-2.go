func isPalindrome(s string) bool {
	n := len(s)
	if n == 0 || n == 1 {
		return true
	}
	s = strings.ToLower(s)
	i := 0
	j := n - 1
	for i <= j {
		// Skip non letters
		// Only checking index i but moving both indices i & j
		if !unicode.IsLetter(rune(s[i])) && !unicode.IsNumber(rune(s[i])) {
			i++
			continue
		}
		if !unicode.IsLetter(rune(s[j])) && !unicode.IsNumber(rune(s[j])) {
			j--
			continue
		}
		if s[i] != s[j] {
			return false
		}
		i++
		j--
	}
	return i >= j
}
