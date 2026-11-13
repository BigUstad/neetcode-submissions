type MRUQueue struct {
    q *list.List
}


func Constructor(n int) *MRUQueue {
    var m MRUQueue
    m.q = list.New()
    for i := 1; i <=n; i++ {
        m.q.PushBack(i)
    }
    return &m
}


func (this *MRUQueue) Fetch(k int) int {
    // Starting from the Back.
    // Assumes fetch will be repeated to check the speed
    cur := this.q.Back().Value.(int)
    if this.q.Len() == k {
        return cur
    }
    e := this.q.Front()
    for k > 0 {
        cur = e.Value.(int)
        // fmt.Print("  ", cur)
        k--
        if k > 0 { e = e.Next()}
    }
    // fmt.Println()
    this.q.MoveAfter(e, this.q.Back())
    return cur
}
