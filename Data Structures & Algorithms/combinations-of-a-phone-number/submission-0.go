func letterCombinations(digits string) []string {
    charSet := []string{
        "", "",
        "abc", "def", "ghi", "jkl",
        "mno", "pqrs", "tuv", "wxyz",
    }

    n := len(digits)
    if n == 0 {
        return []string{}
    }

    res := make([]string, 0)
    cur := make([]byte, n)

    var dfs func(int)

    dfs = func(index int) {
        if index == n {
            res = append(res, string(cur))
            return
        }

        d := digits[index] - '0'

        for i := 0; i < len(charSet[d]); i++ {
            cur[index] = charSet[d][i]
            dfs(index + 1)
        }
    }

    dfs(0)
    return res
}
