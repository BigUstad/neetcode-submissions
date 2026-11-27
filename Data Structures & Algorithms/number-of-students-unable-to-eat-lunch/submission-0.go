import (
	"slices"
)
func countStudents(students []int, sandwiches []int) int {
	// NOTE: This solution "somewhat" works.
	// But the queue simulation in solutions are better and have better performance.
	// The total count of repeatitions is arbitrary.
	// There could be a better check for how many times after rotation the students are not leaving the queue
	// Also, after x rotations both arrays should have the same elements. 
	// Or after sorting?
	slices.Sort(slices.Clone(students))
	slices.Sort(slices.Clone(sandwiches))
	stStr := fmt.Sprint(students)
	saStr := fmt.Sprint(sandwiches)
	fmt.Println(saStr, stStr)
	if stStr == saStr {
		return 0
	}
	// Sandwich counts
	zSnCnt := 0
	oSnCnt := 0
	for _, s := range students {
		if s == 0 {
			zSnCnt++
		} else {
			oSnCnt++
		}
	}
	for _, s := range sandwiches {
		// fmt.Println(zSnCnt, oSnCnt)
		if s == 0 {
			if zSnCnt > 0 {
				zSnCnt--
			} else {
				break
			}
		} else {
			if oSnCnt > 0 {
				oSnCnt--
			} else {
				break
			}
		}
	}
	return zSnCnt + oSnCnt
}
