func firstUniqChar(s string) int {
	f := make([]uint8, 26)
	for _, c := range s {
		f[c-'a']++
	}
	// fmt.Println(f)
	one := uint8(1)
	for i, c := range s {
		if f[c-'a'] == one {
			return i
		}
	}
	return -1
}
