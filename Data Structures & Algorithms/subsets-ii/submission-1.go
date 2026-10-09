import ("slices")
func subsetsWithDup(nums []int) [][]int {
	var allResults [][]int
	slices.Sort(nums)
	n := len(nums)
	var backtrack func(int, []int)
	backtrack = func(index int, cur []int) {
		// Split cur into individual elements
		// fmt.Println(append([]int{}, cur...))
		allResults = append(allResults, append([]int{}, cur...))
		for i := index; i < n; i++ {
			if i > index && nums[i] == nums[i-1] {
				continue
			}
			cur = append(cur, nums[i])
			backtrack(i+1, cur)
			cur = cur[:len(cur)-1]
		}
	}
	backtrack(0, []int{})
	return allResults
}
