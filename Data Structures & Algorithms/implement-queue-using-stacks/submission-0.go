type MyQueue struct {
	s1 []int
	s2 []int
    front int
}

// I'm going to maintain s1 empty
// Push to s1 leads to push all from s2 to s1.
// Then push all from s1 to s2
func Constructor() MyQueue {
    return MyQueue{}
}


func (this *MyQueue) Push(x int)  {
    // fmt.Println("1a. ", this.s1, this.s2)
    if len(this.s2) == 0 && len(this.s1) == 0 {
        this.front = x
    }
    for len(this.s1) > 0 {
        last := len(this.s1) - 1
        ele := this.s1[last]
        // THis is pop of s1, LIFO
        this.s1 = this.s1[:last]
        this.s2 = append(this.s2, ele)
    }
    this.s2 = append(this.s2, x)

    // fmt.Println("1b. ", this.s1, this.s2)
}


func (this *MyQueue) Pop() int {
    // fmt.Println("2a. ", this.s1, this.s2)
    for len(this.s2) > 0 {
        last := len(this.s2) - 1
        ele := this.s2[last]
        // This is pop of s2 LIFO
        this.s2 = this.s2[:last]
        this.s1 = append(this.s1, ele)
    }
    last := len(this.s1) - 1
    ret := this.s1[last]
    this.s1 = this.s1[:last]
    last = len(this.s1) - 1
    if last >= 0 {
        this.front = this.s1[last]
    }
    // fmt.Println("2b. ", this.s1, this.s2)
    return ret
}


func (this *MyQueue) Peek() int {
    return this.front
}


func (this *MyQueue) Empty() bool {
    return len(this.s1) == 0 && len(this.s2) == 0
}

/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param2 := obj.Pop();
 * param3 := obj.Peek();
 * param4 := obj.Empty();
 */
