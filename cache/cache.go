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

func NewLRUCache(capacity int) *LRUCache {
	return &LRUCache{
		capacity: capacity,
		cache:    make(map[int]*list.Element),
		list:     list.New(),
	}
}

func (c *LRUCache) Put(key, value int) {
	if elem, ok := c.cache[key]; ok {
		c.list.MoveToFront(elem)
		elem.Value = Pair{key, value}
	} else {
		if c.list.Len() >= c.capacity {
			delete(c.cache, c.list.Back().Value.(Pair).key)
			c.list.Remove(c.list.Back())
		}
		c.list.PushFront(list.Element{Value: Pair{key, value}})
		c.cache[key] = c.list.Front()
	}
}

func (c *LRUCache) Get(key int) (int, bool) {
	if elem, ok := c.cache[key]; ok {
		c.list.MoveToFront(elem)
		return elem.Value.(Pair).value, true
	}
	return -1, false
}
