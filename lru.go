package lru

import (
	"container/list"
)

type Cache struct {
	MaxEntires int
	ll         *list.List
	cache      map[interface{}]*list.Element
}
type Key interface{}

type entry struct {
	key   Key
	value interface{}
}

func NewCache(maxEnties int) *Cache {
	return &Cache{
		MaxEntires: maxEnties,
		ll:         list.New(),
		cache:      make(map[interface{}]*list.Element),
	}
}

func (c *Cache) Add(key Key, value interface{}) {
	if c.cache == nil {
		c.cache = make(map[interface{}]*list.Element)
		c.ll = list.New()
	}

	if ele, ok := c.cache[key]; ok {
		c.ll.MoveToFront(ele)
		ele.Value.(*entry).value = value
		return
	}

	c.ll.PushFront(value)
	if c.MaxEntires != 0 && c.Len() > c.MaxEntires {
		c.RemoveOldest()
	}
}

func (c *Cache) Get(key Key) (value interface{}, ok bool) {
	if c.cache == nil {
		return
	}

	if ele, hit := c.cache[key]; hit {
		c.ll.MoveToFront(ele)
		return ele.Value.(*entry).value, true
	}
	return
}

func (c *Cache) RemoveOldest() {
	if c.cache == nil {
		return
	}
	ele := c.ll.Back()
	if ele != nil {
		c.removeElement(ele)
	}

}

func (c *Cache) removeElement(e *list.Element) {
	c.ll.Remove(e)
	delete(c.cache, e)

}

func (c *Cache) Len() int {
	return c.ll.Len()
}
