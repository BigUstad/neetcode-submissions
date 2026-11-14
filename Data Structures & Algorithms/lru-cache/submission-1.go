
// LRU Ele will be stored in LRU linked list
type LRUEle struct {
	key int
	value int
}

// The actual "cache" will be the lru list element
// key: key(Put input), value: list element where key/value is stored.
// Cache eviction: When LRU cache is "at" capacity, the node at the back of the list is evicted.
// Cache eviction: During eviction, the back of the elemtn gives the key and that element is deleted in map.
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

func (this *LRUCache) processEviction() {
	if this.lru.Len() <= this.capacity {
		return
	}
	delete(this.lrumap, this.lru.Back().Value.(LRUEle).key)
	this.lru.Remove(this.lru.Back())
}
// Cache eviction: When LRU cache is "at" capacity, the node at the back of the list is evicted.
// Cache eviction: During eviction, the back of the elemtn gives the key and that element is deleted in map.
// TODO: Carve out eviction in a separate function
func (this *LRUCache) Get(key int) int {
    listele, ok := this.lrumap[key]
	if !ok { return -1 }
	ele := this.lru.PushFront(LRUEle{key: key, value: listele.Value.(LRUEle).value,})
	this.lrumap[key] = ele
	this.lru.Remove(listele)

	this.processEviction()

	return listele.Value.(LRUEle).value
}

// Cache eviction: When LRU cache is "at" capacity, the node at the back of the list is evicted.
// Cache eviction: During eviction, the back of the elemtn gives the key and that element is deleted in map.
// TODO: Carve out eviction in a separate function
func (this *LRUCache) Put(key int, value int) {
	listele, ok := this.lrumap[key]
    delete(this.lrumap, key)
	ele := this.lru.PushFront(LRUEle{key: key, value: value,})
	if ok { this.lru.Remove(listele) }
	this.lrumap[key] = ele
	this.processEviction()
}
