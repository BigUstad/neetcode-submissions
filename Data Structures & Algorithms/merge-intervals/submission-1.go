import (
	"slices"
	"github.com/emirpasic/gods/maps/treemap"
)
func merge(intervals [][]int) [][]int {
    var res [][]int
	m := treemap.NewWithIntComparator()
	for _, i := range intervals {
		k1, k2 := i[0], i[1]
		v1, v2 := 1, -1
		if mvi, e := m.Get(k1); e {
			mv := mvi.(int)
			mv++
			v1 = mv
		}
		m.Put(k1, v1)
		if mvi, e := m.Get(k2); e {
			mv := mvi.(int)
			mv--
			v2 = mv 
		}
		m.Put(k2, v2)
	}
	mItr := m.Iterator()
	// fmt.Println(m)
	var interval []int
	have := 0
	for mItr.Next() {
		k, v := mItr.Key().(int), mItr.Value().(int)
		if len(interval) == 0 {
			interval = append(interval, k)
		}
		have += v
		if len(interval) > 0 && have == 0 {
			interval = append(interval, k)
			if len(interval) == 1 {
				// Interval with same value [k,k]
				interval = append(interval, interval[0])
			}
			res = append(res, slices.Clone(interval))
			interval = interval[:0]
		}
	}
	return res
}
