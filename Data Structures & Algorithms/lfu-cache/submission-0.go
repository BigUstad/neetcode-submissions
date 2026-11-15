type LFUEle struct {
    key int
    value int
    curFreq int
}

type LFUCache struct {
    // Frequency count map. key: freq value: linkedlist of LFUEle
    freqMap map[int]*list.List
    // cache map. key: key, value: linkedlist LFU element in freqMap.
    cache map[int]*list.Element
    leastFreq int
    capacity int
}


func Constructor(capacity int) LFUCache {
    return LFUCache{
        freqMap: make(map[int]*list.List),
        cache: make(map[int]*list.Element),
        capacity: capacity,
    }
}

// cache eviction is called before the new key is inserted.
// cache eviction gets rid of the least recently used item in the list.
// THe list is the list of least frequency count - hence the least frequently used list.
func (this *LFUCache) cacheEviction() {
    // fmt.Println("cache eviction", this.count, this.capacity)
    if len(this.cache) < this.capacity {
        return
    }
    v, ok := this.freqMap[this.leastFreq]
    // This would be an error
    if !ok || v.Len() == 0 {
        return
    }
    // fmt.Println("cache eviction", v.Front().Value.(LFUEle))
    delete(this.cache, v.Front().Value.(LFUEle).key)
    v.Remove(v.Front())
    if this.freqMap[this.leastFreq].Len() == 0 {
        delete(this.freqMap, this.leastFreq)
        // this.leastFreq++
    }
}

func (this *LFUCache) updateLeastFreq() {
    least := 200000
    for k, _ := range this.freqMap {
        if k < least {
            least = k
        }
    }
    this.leastFreq = least
}

func (this *LFUCache) Get(key int) int {
    v, ok := this.cache[key]
    if !ok {
        return -1
    }
    curFreq := v.Value.(LFUEle).curFreq
    _, ok = this.freqMap[curFreq]
    if ok {
        oldList := this.freqMap[curFreq]
        oldList.Remove(v)

        if oldList.Len() == 0 {
            delete(this.freqMap, curFreq)
            // if this.leastFreq == curFreq {
            //     this.leastFreq++
            // }
        }
    }
    var e LFUEle
    e.key = key
    e.value = v.Value.(LFUEle).value
    e.curFreq = curFreq+1
    _, ok = this.freqMap[curFreq+1]
    if !ok {
        this.freqMap[curFreq+1] = list.New()
    }
    ve := this.freqMap[curFreq+1].PushBack(e)
    this.cache[key] = ve
    this.updateLeastFreq()
    return ve.Value.(LFUEle).value
}


func (this *LFUCache) Put(key int, value int)  {
    if this.capacity == 0 {
        return
    }
    v, ok := this.cache[key]
    if !ok {
        this.cacheEviction()
        this.leastFreq = 1
        var e LFUEle
        e.key = key
        e.value = value
        e.curFreq = 1
        _, ok = this.freqMap[this.leastFreq]
        if !ok {
            this.freqMap[this.leastFreq] = list.New()
        }
        ve := this.freqMap[this.leastFreq].PushBack(e)
        this.cache[key] = ve
    } else {
        curFreq := v.Value.(LFUEle).curFreq
        this.freqMap[curFreq].Remove(v)
        if this.freqMap[curFreq].Len() == 0 {
            delete(this.freqMap, curFreq)
        }
        var v LFUEle
        v.key = key
        v.value = value
        v.curFreq = curFreq+1
        _, ok = this.freqMap[curFreq+1]
        if !ok {
            this.freqMap[curFreq+1] = list.New()
        }
        ve := this.freqMap[curFreq+1].PushBack(v)
        this.cache[key] = ve
    }
    this.updateLeastFreq()
}


/**
 * Your LFUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */