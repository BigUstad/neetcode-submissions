func twoSum(nums []int, target int) []int {
    diffMap := make(map[int]int)

    for i, n := range nums {
        v, exists := diffMap[n]
        if exists {
			i, v = min(i, v), max(i, v)
            return []int{i, v}
        }
        diffMap[target - n] = i
    }
    return []int{}    
}
