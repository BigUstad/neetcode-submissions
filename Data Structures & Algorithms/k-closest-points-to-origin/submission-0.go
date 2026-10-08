type Element struct {
    a int
    b int
    d int
}

type XYHeap []Element
func (h XYHeap) Len() int { return len(h) }
func (h XYHeap) Less(i, j int) bool { return h[i].d > h[j].d }
func (h XYHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }
func (h *XYHeap) Push(x any) {
	*h = append(*h, x.(Element))
}

func (h *XYHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

func kClosest(points [][]int, k int) [][]int {
	var xe []Element
	h := XYHeap(xe)
	heap.Init(&h)
	for _, p := range points {
		e := Element{
			a: p[0],
			b: p[1],
			d: ((p[0]*p[0])+(p[1]*p[1])),
		}
		if h.Len() < k {
			heap.Push(&h, e)
			continue
		}
		top := heap.Pop(&h).(Element)
		if e.d < top.d {
			heap.Push(&h, e)
		} else {
			heap.Push(&h, top)
		}
	}
	res := make([][]int, k)
	for i := 0; i < k; i++ {
		e := heap.Pop(&h).(Element)
		res[i] = []int{e.a, e.b}
	}
	return res
}
