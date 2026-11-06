func isPalindrome(s string) bool {
	i, j := 0, len(s)-1
	// Palindrome check
	for i <= j {
		if s[i] != s[j] {
			return false
		}
		i++
		j--
	}
	return true
}
func validPalindrome(s string) bool {
	if isPalindrome(s) {
		return true
	}
	i, j := 0, len(s)-1
	for i <= j {
		if s[i] != s[j] {
			break
		}
		i++
		j--
	}
    // Check the rest of the string skipping
    // either 'i' or 'j'
	// fmt.Println(s[i+1:j+1])
	// fmt.Println(s[i:j])
	return isPalindrome(s[i+1:j+1]) || isPalindrome(s[i:j])
}
