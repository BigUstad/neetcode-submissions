import (
    tm "github.com/emirpasic/gods/maps/treemap"
)

func addToMap(m *tm.Map, ele int) {
    if cI, e := m.Get(ele); e {
        c := cI.(int)
        c++
        m.Put(ele, c)
        return
    }
    m.Put(ele, 1)
}

func removeFromMap(m *tm.Map, ele int) {
    cI, e := m.Get(ele)
    if !e {
        return
    }
    c := cI.(int)
    c--
    if c == 0 {
        m.Remove(ele)
        return
    }
    m.Put(ele, c)
}

func maxSlidingWindow(nums []int, k int) []int {
    // The treeset sliding window
    m := tm.NewWithIntComparator()
    var res []int
    for i, n := range nums {
        if i < k-1 {
            addToMap(m, n)
            continue
        } else if i >= k {
            removeFromMap(m, nums[i-k])
        }
        addToMap(m, n)
        // Check the end.
        // Skipping the bool check as tree being not empty is ensured
        k, _ := m.Max()
        res = append(res, k.(int))
        // fmt.Println(res)
    }
    return res
}