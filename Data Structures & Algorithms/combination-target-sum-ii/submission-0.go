import ("slices")

func combinationSum2(candidates []int, target int) [][]int {
    if len(candidates) == 0 { return [][]int{} }
    if len(candidates) == 1 {
        if target == candidates[0] {
            return [][]int{{candidates[0]}}
        }
        return [][]int{}
    }
    var backtrack func(int, int, []int)
    var allResults [][]int
    n := len(candidates)
    slices.Sort(candidates)
    backtrack = func(need, index int, cur []int) {
        // fmt.Println(cur)
        // target has been met.
        // this set can be added
        if need == 0 {
            res := slices.Clone(cur)
            // strRes := fmt.Sprint(res)
            // if m[strRes] {
            //     return
            // }
            // m[strRes] = true
            allResults = append(allResults, res)
            return
        } // need > 0 would need more iterations below.
        if need < 0 || index == n {
            // this set can be rejected
            // the accumulated elements' sum is greater than target
            // Or we've exhausted the candidates array
            return
        }
        // Two-branch recursion?
        // We make the decision - include candidates[i]
        cur = append(cur, candidates[index])
        backtrack(need - candidates[index], index+1, cur)
        // backtracking move - knock off the end element
        // so that the next index in the array can be appended for analysis
        // We make the decision skip candidates[i].
        // After backtracking skip the equal elements for next iteration of branching.
        cur = cur[:len(cur)-1]
        for index < n-1 && candidates[index] == candidates[index+1] {
            index++
        }
        // Another backtracking move. We start with a fresh candidate with original total
        backtrack(need, index+1, cur)
    }
    backtrack(target, 0, []int{})

    return allResults
}
