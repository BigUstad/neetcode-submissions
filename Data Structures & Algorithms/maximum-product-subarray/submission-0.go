func maxProduct(nums []int) int {
	// There is no dp or memo. curMax & curMin the dp parameters
	// The DP solution rests on calculating the most optimal next move
	// based on current step
	// Time complexity O(n), Space O(1)
    res := nums[0]
	// Why is min needed?
	// min is needed to keep track of a negative multiplier that could flip & be 0?
	curMin, curMax := 1, 1
	// 1 is chosen by Neetcode as a neutral value
	// Reason being if 0 is used, it is like introducing a poison pill
	// which makes min or max become 0
	for _, n := range nums {
		// Neetcode says if condition is unnecessary but I kept it
		// Then it gave the wrong answer for testcase [-3, 0, -2]
		// I guess the curMax strives to be positive with tmpMin being included, so it is ok to remove if condition
		// if n == 0 {
		// 	curMin, curMax = 1, 1
		// 	continue
		// }
		tmpMax := curMax * n
		tmpMin := curMin * n
		// tmpMin is added in cases where curMax was previously negative & tmpMin becomes positive
		// curMin & curMax could both be negative, approaching 0. Hence adding n could give the correct option
		// in taking current max
		// In Neetcode's solution curMax = max(tmpMax, max(tmpMin, n)) & curMin = min(tmpMax, min(tmpMin, num))

		curMax = max(tmpMax, tmpMin, n)
		curMin = min(tmpMax, tmpMin, n)
		// fmt.Println(i, ") ", n, ", ", curMax, ", ", curMin)
		// Keeping track of resultant max
		res = max(res, curMax)
	}
	return res
}
