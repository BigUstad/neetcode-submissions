import ("slices")
func subsetsWithDup(nums []int) [][]int {
    n := len(nums)
    if n == 0 { return [][]int{} }
    if n == 1 {
        return [][]int{{}, {nums[0]}}
    }
    var backtrack func(int, []int)
    var allResults [][]int
    allResults = append(allResults, []int{})
    slices.Sort(nums)
    m := make(map[string]bool)
    backtrack = func(index int, cur []int) {
        if len(cur) > 0 {
            // fmt.Println(cur)
            key := fmt.Sprint(cur)
            if m[key] {
                // Don't pursue this subarray which is already there
                return
            }
            m[key] = true
            res := slices.Clone(cur)
            allResults = append(allResults, res)
        }
        if index >= n {
            return
        }
        for i := index; i < n; i++ {
            cur = append(cur, nums[i])
            backtrack(i+1, cur)
            cur = cur[:len(cur)-1]
        }
    }
    backtrack(0, []int{})
    return allResults
}
