/**
 * Definition of Interval:
 * type Interval struct {
 *    start int
 *    end   int
 * }
 */

import (
	"cmp"
	"slices"
	pq "github.com/emirpasic/gods/queues/priorityqueue"
)

// latest on top
func pqIntervalComparator(a, b interface{}) int {
	return cmp.Or(
		-cmp.Compare(a.(Interval).end, b.(Interval).end),
		-cmp.Compare(a.(Interval).start, b.(Interval).start),
	)
}

func minIntervalComparator(a, b Interval) int {
	return cmp.Or(
		cmp.Compare(a.start, b.start),
		cmp.Compare(a.end, b.end),
	)
}

func minMeetingRooms(intervals []Interval) int {
	l := len(intervals)
	if l <= 1 { return l }
	var sPQ []*pq.Queue
	slices.SortFunc(intervals, minIntervalComparator)
	// First!!
	sPQ = append(sPQ, pq.NewWith(pqIntervalComparator))
	for _, v := range intervals {
		// fmt.Println(v)
		foundPQ := false
		// Find the right pq
		for _, p := range sPQ {
			t, ok := p.Peek()
			if !ok ||
				(ok && v.start >= t.(Interval).end) {
				// if ok {
				// 	fmt.Println(v)
				// 	fmt.Println(t)
				// }
				foundPQ = true
				p.Enqueue(v)
				break
			}
		}
		if !foundPQ {
			newP := pq.NewWith(pqIntervalComparator)
			sPQ = append(sPQ, newP)
			newP.Enqueue(v)

		}
		// fmt.Println(len(sPQ))
		// fmt.Println("***")
	}
	return len(sPQ)
}
