package lrucache

import (
	"container/list"
	"sync"
)

type LRUCache struct {
	capacity int
	cache    map[interface{}]*list.Element
	list     *list.List
	mu       sync.Mutex
}

type entry struct {
	key   interface{}
	value interface{}
}

func NewLRUCache(capacity int) *LRUCache {
	return &LRUCache{
		capacity: capacity,
		cache:    make(map[interface{}]*list.Element),
		list:     list.New(),
	}
}

func (c *LRUCache) Get(key interface{}) interface{} {
	c.mu.Lock()
	defer c.mu.Unlock()

	if elem, ok := c.cache[key]; ok {
		c.list.MoveToFront(elem)
		return elem.Value.(*entry).value
	}
	return nil
}

func (c *LRUCache) Put(key, value interface{}) {
	c.mu.Lock()
	defer c.mu.Unlock()

	if elem, ok := c.cache[key]; ok {
		c.list.MoveToFront(elem)
		elem.Value.(*entry).value = value
		return
	}

	if c.list.Len() >= c.capacity {
		// Remove the least recently used item
		lastElem := c.list.Back()
		if lastElem != nil {
			delete(c.cache, lastElem.Value.(*entry).key)
			c.list.Remove(lastElem)
		}
	}

	newElem := c.list.PushFront(&entry{key, value})
	c.cache[key] = newElem
}

func (c *LRUCache) Len() int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.list.Len()
}
