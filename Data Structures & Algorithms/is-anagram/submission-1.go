func isAnagram(s string, t string) bool {
	setChar := make([]int, 26)
	for _, c := range s {
		setChar[int(c) - int('a')]++
	}
	for _, c := range t {
		setChar[int(c) - int('a')]--
	}
	for _, i := range setChar {
		if i > 0 || i < 0 {
			return false
		}
	}
	return true
}
