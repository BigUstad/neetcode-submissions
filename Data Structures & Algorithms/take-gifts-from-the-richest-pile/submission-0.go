import (
    "github.com/emirpasic/gods/trees/binaryheap"
	"cmp"
)
func pickGifts(gifts []int, k int) int64 {
    inverseInt64Comparator := func(a, b interface{}) int {
		return -cmp.Compare(a.(int64), b.(int64))
	}
    heap := binaryheap.NewWith(inverseInt64Comparator)
    for _, g := range gifts {
        heap.Push(int64(g))
    }
    if heap.Empty() { return int64(0) }
    for k > 0 {
        t, _ := heap.Pop()
        fs := math.Floor(math.Sqrt(float64(t.(int64))))
        ts := int64(fs)
        heap.Push(ts)
        k--
    }
    var sum int64 = 0
    for !heap.Empty() {
        t, _ := heap.Pop()
        sum += t.(int64)
    }
    return sum
}
