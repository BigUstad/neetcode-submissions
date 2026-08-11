func combinationSum(nums []int, target int) [][]int {
    if len(nums) == 0 { return [][]int{} }
	if len(nums) == 1 {
		if target == nums[0] {
			return [][]int{{nums[0]}}
		}
		return [][]int{}
	}
	var backtrack func(need int, index int)
	var allResults [][]int
	var currentPath []int
	backtrack = func(need int, index int) {
		// fmt.Println(currentPath)
		if need == 0 {
			res := make([]int, len(currentPath))
			copy(res, currentPath)
			// fmt.Print("Selected - ")
			// fmt.Println(res)
			allResults = append(allResults, res)
			// currentPath = currentPath[:0]
		} else if need < 0 {
			// Has moved past, so need to retrack
			return
		}
		for i := index; i < len(nums); i++ {
			currentPath = append(currentPath, nums[i])
			backtrack(need - nums[i], i)
			currentPath = currentPath[0:(len(currentPath)-1)]
		}
	}
	backtrack(target, 0)
	return allResults
}
