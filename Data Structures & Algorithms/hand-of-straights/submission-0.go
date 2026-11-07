import (
    "github.com/emirpasic/gods/trees/binaryheap"
)
func isNStraightHand(hand []int, groupSize int) bool {
    // min max map
    min := math.MaxInt
    max := math.MinInt
    // min heap but all values in hand.
    // Duplicates are allowed in binaryheap
    heap := binaryheap.NewWithIntComparator()
    // map keeps track of the count of duplicates so that I'm able use duplicate cards
    // but not use it after it has been grouped.
    m := make(map[int]int)
    for _, h := range hand {
        heap.Push(h)
        m[h]++
        if h < min {
            min = h
        }
        if h > max {
            max = h
        }
    }
    // fmt.Println(heap.Values())

    for !heap.Empty() {
        hi, _ := heap.Pop()
        h := hi.(int)
        if m[h] == 0 {
            continue
        }
        j := h
        count := groupSize
        // group of groupSize cards accumulation
        for count > 0  {
            if m[j] == 0 {
                return false
            }
            if m[j] > 0 {
                m[j]--
                count--
            }
            j++
        }
        if count > 0 {
            return false
        }
    }

    return true
}
