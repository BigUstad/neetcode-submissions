// import ("container/heap")
type IntHeap []int

func (h IntHeap) Len() int { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] } // Min-heap: lowest top
func (h IntHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }
// Peek returns the top element without removing it from the heap
func (h IntHeap) Peek() int {
    if len(h) == 0 {
        panic("Cannot peek at an empty heap")
    }
    return h[0]
}

func (h *IntHeap) Push(x any) {
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0:n-1]
	return x
}

func findKthLargest(nums []int, k int) int {
	minHeap := make(IntHeap, 0, k)
	heap.Init(&minHeap)
	for i, n := range nums {
		if i >= k && minHeap.Peek() > n {
			// No need to push if n is less than top
			continue
		}
		if i >= k {
			heap.Pop(&minHeap)
		}
		heap.Push(&minHeap, n)
	}
	return minHeap.Peek()
}
