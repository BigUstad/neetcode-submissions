type MinminH []int

func (h MinminH) Len() int            { return len(h) }
func (h MinminH) Less(i, j int) bool  { return h[i] < h[j] }
func (h MinminH) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *MinminH) Push(x interface{}) { *h = append(*h, x.(int)) }
func (h *MinminH) Pop() interface{} {
    old := *h
    n := len(old)
    x := old[n-1]
    *h = old[:n-1]
    return x
}

func isNStraightHand(hand []int, groupSize int) bool {
    // min max map
    min := math.MaxInt
    max := math.MinInt
    // min minH but all values in hand.
    // Duplicates are allowed in binaryminH
    minH := &MinminH{}
    // map keeps track of the count of duplicates so that I'm able use duplicate cards
    // but not use it after it has been grouped.
    m := make(map[int]int)
    for _, h := range hand {
        *minH = append(*minH, h)
        m[h]++
        if h < min {
            min = h
        }
        if h > max {
            max = h
        }
    }
	heap.Init(minH)
    // fmt.Println(minH.Values())

    for minH.Len() > 0 {
        h := (*minH)[0]
		heap.Pop(minH)
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
