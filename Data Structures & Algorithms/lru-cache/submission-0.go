
type LRUEle struct {
	key int
	value int
}

type LRUCache struct {
	lrumap map[int]*list.Element
	lru *list.List
    capacity int
}

func Constructor(capacity int) LRUCache {
    return LRUCache{
		lrumap: make(map[int]*list.Element),
		lru: list.New(),
        capacity: capacity,
	}
}

func (this *LRUCache) Get(key int) int {
    listele, ok := this.lrumap[key]
	if !ok { return -1 }
	ele := this.lru.PushFront(LRUEle{key: key, value: listele.Value.(LRUEle).value,})
	this.lrumap[key] = ele
	this.lru.Remove(listele)
    // fmt.Print("Get: ")
    // fmt.Print(key)
    // fmt.Print(" - ")
    // fmt.Println(this.lru.Len())
    if this.lru.Len() > this.capacity {
        delete(this.lrumap, this.lru.Back().Value.(LRUEle).key)
        this.lru.Remove(this.lru.Back())
    }
	return listele.Value.(LRUEle).value
}

func (this *LRUCache) Put(key int, value int) {
	listele, ok := this.lrumap[key]
    delete(this.lrumap, key)
	ele := this.lru.PushFront(LRUEle{key: key, value: value,})
	if ok { this.lru.Remove(listele) }
	this.lrumap[key] = ele
    // fmt.Print("Put: ")
    // fmt.Print(key)
    // fmt.Print(" - ")
    // fmt.Println(this.lru.Len())
    if this.lru.Len() > this.capacity {
        delete(this.lrumap, this.lru.Back().Value.(LRUEle).key)
        this.lru.Remove(this.lru.Back())
    }
    // fmt.Println(this.lru.Len())
}
