type Solution struct{}
// trademark is used for special cases
const r_trademark rune = '\u00AE'
// degree is used for regular cases
const r_degree rune = '\u00B0'

func (s *Solution) Encode(strs []string) string {
	// Encode empty string list
	if len(strs) == 0 {
		return ""
	}
	// Encode single item string list
	if len(strs) == 1 {
		return strings.Join([]string{
			string(r_trademark),
			strs[0],
			string(r_trademark),
			}, "")
	}
	return strings.Join(strs, string(r_degree))
}

func (s *Solution) Decode(encoded string) []string {
	const empty_arr_encoded = string(r_trademark) + string(r_trademark)
	// Special cases
	if len(encoded) == 0 {
		return []string{}
	}
	if encoded == empty_arr_encoded {
		return []string{""}
	}
	// single item string
	// Special encoding with trademark
	if encoded[0:len(string(r_trademark))] == string(r_trademark) {
		return []string{strings.Trim(encoded, string(r_trademark))}
	}
	return strings.Split(encoded, string(r_degree))
}
