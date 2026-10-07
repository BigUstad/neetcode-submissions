import (
	"cmp"
	"slices"
)
func eraseOverlapIntervals(intervals [][]int) int {
    n := len(intervals)
    // length of resultant list
	res := 0
	// Sort by start of interval first & then last of interval
	slices.SortFunc(intervals, func(i, j []int) int {
		return cmp.Compare(i[1], j[1])
	})
	// fmt.Println(intervals)
	// My original approach was similar to dynamic programming top-down approach
	// Glitch in logic was to retain the smaller of end-interval in a completely overlapping interval
	// In [1,2][1,4][2,4] -> Retain [1,2][2,4]
	// Bottom up approach has "minimum removals = total intervals - maximum kept"
    // Top-down approach has 
    cur := []int{intervals[0][0], intervals[0][1]}
	for i := 1; i < n; i++ {
        start, end := intervals[i][0], intervals[i][1]
        // This case is no overlap
        // Just update the start & end value.
        // Move onto next case
        if start >= cur[1] {
            cur[0] = start
            cur[1] = end
        } else {
            // cur start is less than end of cur
            // we have an overlapping interval
            // let's keep the lesser end interval.
            // increment the overlapping count
            res++
            cur[1] = min(end, cur[1])
        }
	}
	// fmt.Println(dp)
	return res
}
