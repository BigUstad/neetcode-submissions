type MyCircularQueue struct {
    capacity int
    queue    []int
    head     int
    tail     int
    count    int
}

// Size is q. 
func Constructor(k int) MyCircularQueue {
    return MyCircularQueue{
        capacity: k,
        queue: make([]int, k),
        head: 0,
        tail: 0,
        count: 0,
    }
}


func (this *MyCircularQueue) EnQueue(value int) bool {
    if this.count == this.capacity {
        return false
    }
    // Modulo can be more expensive than a comparison and subtraction
    // Especially more if the divisor is determined at runtime
    // next := (this.head + this.count) % this.capacity
    next := (this.head + this.count)
    if next >= this.capacity {
        next -= this.capacity
    }
    this.queue[next] = value
    this.tail = next
    this.count++
    return true
}


func (this *MyCircularQueue) DeQueue() bool {
    if this.count == 0 {
        return false
    }
    // Modulo can be more expensive than a comparison and subtraction
    // Especially more if the divisor is determined at runtime
    // this.head = (this.head + 1) % this.capacity
    this.head++
    if this.head >= this.capacity {
        this.head -= this.capacity
    }
    this.count--
    return true
}


func (this *MyCircularQueue) Front() int {
    if this.count == 0 {
        return -1
    }
    return this.queue[this.head]
}


func (this *MyCircularQueue) Rear() int {
    if this.count == 0 {
        return -1
    }
    return this.queue[this.tail]
}

func (this *MyCircularQueue) IsEmpty() bool {
    return this.count == 0
}

func (this *MyCircularQueue) IsFull() bool {
    return this.count == this.capacity
}


/**
 * Your MyCircularQueue object will be instantiated and called as such:
 * obj := Constructor(k);
 * param1 := obj.EnQueue(value);
 * param2 := obj.DeQueue();
 * param3 := obj.Front();
 * param4 := obj.Rear();
 * param5 := obj.IsEmpty();
 * param6 := obj.IsFull();
 */
 