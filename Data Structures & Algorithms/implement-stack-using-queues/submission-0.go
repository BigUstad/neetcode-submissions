type MyStack struct {
    q1 []int
    q2 []int    
}

func Constructor() MyStack {
    return MyStack{}    
}

// I'm going to maintain q1 empty
// Push to q1 leads to push all from q2 to q1.
// Then push all elements from q1 to q2
// This makes it a LIFO
func (this *MyStack) Push(x int)  {
    this.q1 = append(this.q1, x)
    // fmt.Println(this.q1, this.q2)
    for len(this.q2) > 0 {
        ele := this.q2[0]
        this.q2 = this.q2[1:]
        this.q1 = append(this.q1, ele)
    }
    this.q2 = nil
    for len(this.q1) > 0 {
        ele := this.q1[0]
        this.q1 = this.q1[1:]
        this.q2 = append(this.q2, ele)
    }
    this.q1 = nil
    // fmt.Println(this.q1, this.q2)
}


func (this *MyStack) Pop() int {
    ret := this.q2[0]
    this.q2 = this.q2[1:]
    return ret
}


func (this *MyStack) Top() int {
    ret := this.q2[0]
    return ret    
}


func (this *MyStack) Empty() bool {
    return len(this.q2) == 0
}


/**
 * Your MyStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param2 := obj.Pop();
 * param3 := obj.Top();
 * param4 := obj.Empty();
 */
