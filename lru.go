package lru

import (
	"errors"
	"github.com/notEpsilon/go-lru/list"
	"sync"
)

var (
	invalidCapacity = errors.New("capacity must be a positive integer")
	keyNotFound     = errors.New("key doesn't exist")
)

type entry[K comparable, V any] struct {
	key   K
	value V
}

type LRUCache[K comparable, V any] struct {
	lock    sync.RWMutex
	store map[K]*list.Element[*entry[K, V]]
	order *list.List[*entry[K, V]]
	cap   int
}

func New[K comparable, V any](capacity int) (*LRUCache[K, V], error) {
	if capacity <= 0 {
		return nil, invalidCapacity
	}

	return &LRUCache[K, V]{
		store: make(map[K]*list.Element[*entry[K, V]]),
		order: list.New[*entry[K, V]](),
		cap:   capacity,
	}, nil
}

func (c *LRUCache[K, V]) Get(key K) (V, error) {
	c.lock.RLock()
	if en, ok := c.store[key]; ok {
		c.order.MoveToFront(en)
		c.lock.RUnlock()
		return en.Value.value, nil
	}
	c.lock.RUnlock()
	return *new(V), keyNotFound
}

func (c *LRUCache[K, V]) Set(key K, value V) {
	c.lock.Lock()
	if en, ok := c.store[key]; ok {
		en.Value.value = value
		c.order.MoveToFront(en)
		c.lock.Unlock()
		return
	}

	if len(c.store) == c.cap {
		toEvict := c.order.Back()
		delete(c.store, toEvict.Value.key)
		c.order.Remove(toEvict)
	}

	c.store[key] = c.order.PushFront(&entry[K, V]{key: key, value: value})
	c.lock.Unlock()
}

func (c *LRUCache[K, V]) Remove(key K) {
	c.lock.Lock()
	if en, ok := c.store[key]; ok {
		delete(c.store, en.Value.key)
		c.order.Remove(en)
	}
	c.lock.Unlock()
}

func (c *LRUCache[K, V]) Contains(key K) bool {
	c.lock.RLock()
	_, ok := c.store[key]
	c.lock.RUnlock()
	return ok
}

func (c *LRUCache[K, V]) Peek(key K) (V, error) {
	c.lock.RLock()
	if en, ok := c.store[key]; ok {
		val := en.Value.value
		c.lock.RUnlock()
		return val, nil
	}
	c.lock.RUnlock()
	return *new(V), keyNotFound
}

func (c *LRUCache[K, V]) Size() int {
	c.lock.RLock()
	size := len(c.store)
	c.lock.RUnlock()
	return size
}

func (c *LRUCache[K, V]) Capacity() int {
	c.lock.RLock()
	capacity := c.cap
	c.lock.RUnlock()
	return capacity
}
