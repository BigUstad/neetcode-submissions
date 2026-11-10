
type ListEle struct {
    k int
    kv int
}

type MyHashMap struct {
    Buckets []*list.List
    bucketRange int
}


func Constructor() MyHashMap {
    return MyHashMap{
        bucketRange: 769,
        Buckets: make([]*list.List, 769),
    }
}


func (this *MyHashMap) Put(key int, value int)  {
    bucketIndex := key % this.bucketRange
    if this.Buckets[bucketIndex] == nil {
        this.Buckets[bucketIndex] = list.New()
    }
    for e := this.Buckets[bucketIndex].Front(); e != nil; e = e.Next() {
        // fmt.Println("search ", e.Value.(ListEle).k)
        if e.Value.(ListEle).k == key {
            // fmt.Println("found ", e.Value.(ListEle).k)
            if ele, ok := e.Value.(ListEle); ok {
                // fmt.Println("Put Replace key = ", key, " old value = ", e.Value.(ListEle).kv ," value = ", value, " bucketIndex = ", bucketIndex)
                ele.kv = value
				// IMPORTANT: ele is a copy. e is the iterator. Use the iterator to modify the original.
				// TODO: Using pointers is better, it avoids the replacement effort in Put.
                e.Value = ele
            }
            // fmt.Println("Put check key = ", e.Value.(ListEle).k, " value = ", e.Value.(ListEle).kv, " bucketIndex = ", bucketIndex)
            return
        }
    }
	// TODO: Using pointers is better, it avoids the replacement effort in Put on 33.
    /*e := */this.Buckets[bucketIndex].PushBack(ListEle{k: key, kv: value})
    // fmt.Println("Put key = ", e.Value.(ListEle).k, " value = ", e.Value.(ListEle).kv, " bucketIndex = ", bucketIndex)
}


func (this *MyHashMap) Get(key int) int {
    bucketIndex := key % this.bucketRange
    if this.Buckets[bucketIndex] == nil ||
        (this.Buckets[bucketIndex] != nil && this.Buckets[bucketIndex].Len() <= 0) {
        return -1
    }
    for e := this.Buckets[bucketIndex].Front(); e != nil; e = e.Next() {
        if e.Value.(ListEle).k == key {
            // fmt.Println("Get key = ", e.Value.(ListEle).k, " value = ", e.Value.(ListEle).kv, " bucketIndex = ", bucketIndex)
            return e.Value.(ListEle).kv
        }
    }
    return -1
}


func (this *MyHashMap) Remove(key int)  {
    bucketIndex := key % this.bucketRange
    if this.Buckets[bucketIndex] == nil ||
        (this.Buckets[bucketIndex] != nil && this.Buckets[bucketIndex].Len() <= 0) {
        return
    }
    for e := this.Buckets[bucketIndex].Front(); e != nil; e = e.Next() {
        if e.Value.(ListEle).k == key {
            // fmt.Println("Remove key = ", e.Value.(ListEle).k, " value = ", e.Value.(ListEle).kv, " bucketIndex = ", bucketIndex)
            this.Buckets[bucketIndex].Remove(e)
            return
        }
    }
}


/**
 * Your MyHashMap object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Put(key,value);
 * param_2 := obj.Get(key);
 * obj.Remove(key);
 */