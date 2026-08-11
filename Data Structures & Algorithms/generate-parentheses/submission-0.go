func generateParenthesis(n int) []string {
	var res []string
	var validParanthesesHelper func([]rune, int, int)
	totalCount := 2 * n
	validParanthesesHelper = func(cur []rune, openp, closep int) {
		// fmt.Println(string(cur))
		if len(cur) == totalCount {
			res = append(res, string(cur))
			return
		}
		if openp < n {
			cur = append(cur, '(')
			validParanthesesHelper(cur, openp+1, closep)
			cur = cur[0:len(cur)-1]
		}
		if openp > closep {
			cur = append(cur, ')')
			validParanthesesHelper(cur, openp, closep+1)
			cur = cur[0:len(cur)-1]
		}

	}

	validParanthesesHelper([]rune(""), 0, 0)
	return res
}
