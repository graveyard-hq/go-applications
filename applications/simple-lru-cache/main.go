package main

import (
	"container/list"
)

type LRUCache struct {
	capacity int
	cache    map[int]*list.Element
	ll       *list.List
}

type keyValuePair struct {
	key   int
	value int
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		capacity: capacity,
		cache:    make(map[int]*list.Element),
		ll:       list.New(),
	}
}

func (c *LRUCache) Get(key int) int {
	if elem, ok := c.cache[key]; ok {
		c.ll.MoveToFront(elem)
		return elem.Value.(*keyValuePair).value
	}
	return -1
}

func (c *LRUCache) Put(key int, value int) {
	if elem, ok := c.cache[key]; ok {
		c.ll.MoveToFront(elem)
		elem.Value.(*keyValuePair).value = value
		return
	}
	newElem := c.ll.PushFront(&keyValuePair{key: key, value: value})
	c.cache[key] = newElem
	if c.ll.Len() > c.capacity {
		c.removeOldest()
	}
}

func (c *LRUCache) removeOldest() {
	elem := c.ll.Back()
	if elem != nil {
		c.ll.Remove(elem)
		kv := elem.Value.(*keyValuePair)
		delete(c.cache, kv.key)
	}
}

