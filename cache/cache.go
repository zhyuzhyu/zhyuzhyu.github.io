package cache

import "container/list"

type LRUCache struct {
	capacity int
	cache    map[int]*list.Element
	list     *list.List
}

type Pair struct {
	key, value int
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		capacity: capacity,
		cache:    make(map[int]*list.Element),
		list:     list.New(),
	}
}

func (this *LRUCache) Put(key, value int) {
	if elem, ok := this.cache[key]; ok {
		this.list.MoveToFront(elem)
		elem.Value = Pair{key, value}
	} else {
		this.cache[key] = this.list.PushFront(Pair{key, value})
		if this.list.Len() > this.capacity {
			delete(this.cache, this.list.Remove(this.list.Back()).(Pair).key)
		}
	}
}

func (this *LRUCache) Get(key int) int {
	if elem, ok := this.cache[key]; ok {
		this.list.MoveToFront(elem)
		return elem.Value.(Pair).value
	}
	return -1
}
