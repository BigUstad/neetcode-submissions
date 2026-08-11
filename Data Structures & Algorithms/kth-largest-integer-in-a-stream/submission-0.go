import (
  pq "github.com/emirpasic/gods/queues/priorityqueue"
  "github.com/emirpasic/gods/utils"
)

type KthLargest struct {
    queue *pq.Queue
    qK int
}


func Constructor(k int, nums []int) KthLargest {
    var kth KthLargest
    kth.queue = pq.NewWith(utils.IntComparator)
    kth.qK = k
    for _, n := range nums {
        kth.Add(n)
    }
    return kth
}


func (this *KthLargest) Add(val int) int {
    if this.queue.Size() == this.qK {
        topEle, _ := this.queue.Peek()
        top := topEle.(int)
        // Not in top k
        if val <= top { return top }
        // Pop the least value, automatic heapify after add
        _, _ = this.queue.Dequeue()
    }
    this.queue.Enqueue(val)

    kthEle, _ := this.queue.Peek()
    return kthEle.(int)
}

