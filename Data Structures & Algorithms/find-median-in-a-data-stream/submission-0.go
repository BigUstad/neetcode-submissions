import (
	"github.com/emirpasic/gods/trees/binaryheap"
	"github.com/emirpasic/gods/utils"
)

// Difference between heap sizes should be 1
// The sizes can be same
// The smaller numbers are in maxHeap
// The bigger numbers are in minHeap
type MedianFinder struct {
    minHeap *binaryheap.Heap
    maxHeap *binaryheap.Heap
}

func intMaxCmp(a, b interface{}) int {
    return -utils.IntComparator(a, b)
}

func Constructor() MedianFinder {
    var m MedianFinder
    m.minHeap = binaryheap.NewWithIntComparator()
    m.maxHeap = binaryheap.NewWith(intMaxCmp)
    return m
}


func (this *MedianFinder) AddNum(num int)  {
    // Add to max heap
    this.maxHeap.Push(num)
    // Pop from max heap to balance
    maxi, emax := this.maxHeap.Pop()
    if emax { this.minHeap.Push(maxi.(int)) }
    // Rebalance if required
    if this.maxHeap.Size() < this.minHeap.Size() {
        mini, emin := this.minHeap.Pop()
        if emin { this.maxHeap.Push(mini.(int)) }
    }
}


func (this *MedianFinder) FindMedian() float64 {
    if this.minHeap.Empty() && this.maxHeap.Empty() {
        return float64(0)
    }
	mini, _ := this.minHeap.Peek()
	maxi, _ := this.maxHeap.Peek()
	if this.minHeap.Size() == this.maxHeap.Size() {
		return (float64(mini.(int) + maxi.(int)) / float64(2))
	}
	return float64(maxi.(int))
}


/**
 * Your MedianFinder object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddNum(num);
 * param_2 := obj.FindMedian();
 */