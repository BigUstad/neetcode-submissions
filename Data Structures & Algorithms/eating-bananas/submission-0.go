import (
	"slices"
)
func minEatingSpeed(piles []int, h int) int {
	left := 1
	right := slices.Max(piles)
	res := right
	for left <= right {
		k := left + ((right - left)/2)
		totalTime := 0
		for _, p := range piles {
			totalTime += int(math.Ceil(float64(p)/float64(k)))
		}
		// fmt.Println("t ", totalTime, ", k ", k)
		// fmt.Println("l ", left, " r ", right)
		// fmt.Println(".....")
		if totalTime <= h {
			if k < res {
				res = k
			}
			right = k - 1
		} else {
			left = k + 1
		}
	}
	return res
}
