func groupAnagrams(strs []string) [][]string {
	am := make(map[string][]int)
	for i, s := range strs {
		r := []rune(s)
		sort.Slice(r, func (i int, j int) bool {
			return r[i] <= r[j]
		} )
		sr := string(r)
		if amv, exists := am[sr]; exists {
			// append list of indices
			am[sr] = append(amv, i)
		} else {
			am[sr] = []int{i}
		}
	}
	res := make([][]string, len(am))
	i := 0
	for _, ame := range am {
		res[i] = make([]string, len(ame))
		for j, asi := range ame {
			res[i][j] = strs[asi]
		}
		i++
	}
	return res
}
