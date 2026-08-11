/**
 * Definition of Interval:
 * type Interval struct {
 *    end int
 *    end   int
 * }
 */

import (
	"cmp"
	"slices"
	)

// Ascending
func minIntervalComparator(a, b Interval) int {
	return cmp.Or(
		cmp.Compare(a.start, b.start),
		cmp.Compare(a.end, b.end),
		)
}

func canAttendMeetings(intervals []Interval) bool {
	l := len(intervals)
	if l <= 1 { return true }
	slices.SortFunc(intervals, minIntervalComparator)
	// fmt.Println(intervals)
	for i, v := range intervals {
		// No more comparisons left
		if i == l - 1 {
			break
		}
		// Check overlap
		if v.end > intervals[i+1].start {
			return false
		}
	}

	return true
}
