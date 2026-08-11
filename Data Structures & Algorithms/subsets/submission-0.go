import ("slices")

func backtrack(nums []int, sRet *[][]int, start int, cur []int) {
	*sRet = append(*sRet, slices.Clone(cur))
	for i := start; i < len(nums); i++ {
		cur = append(cur, nums[i])
		// fmt.Println(cur)
		backtrack(nums, sRet, i+1, cur)
		cur = cur[:len(cur)-1] // We're done with i and we are going to the next one.
	}
}

func subsets(nums []int) [][]int {
	var sRet [][]int
	if len(nums) == 0 {
		return sRet
	}
	if len(nums) == 1 {
		return [][]int {
			{},
			{nums[0]},
		}
	}

	backtrack(nums, &sRet, 0, []int{})

	return sRet
}
