type MyHashSet struct {
    myset []bool
}

func Constructor() MyHashSet {
    return MyHashSet {
        myset: make([]bool, 1000001),
    }
}

func (this *MyHashSet) Add(key int) {
    this.myset[key] = true
}

func (this *MyHashSet) Remove(key int) {
    this.myset[key] = false
}

func (this *MyHashSet) Contains(key int) bool {
    return this.myset[key]
}

/**
 * Your MyHashSet object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Add(key);
 * obj.Remove(key);
 * param_3 := obj.Contains(key);
 */
 