func combine(n int, k int) [][]int {
    var combinations [][]int
    var backtrack func(int, []int)

    backtrack = func(start int, cur []int) {
        if len(cur) == k {
            combination := append([]int(nil), cur...)
            combinations = append(combinations, combination)
        }
        for i := start; i <= n; i++ {
            cur = append(cur, i)
            backtrack(i+1, cur)
            // Actual backtrack
            cur = cur[:len(cur)-1]
        }
    }
    backtrack(1, nil)
    return combinations
}
