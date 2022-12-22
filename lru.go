package lru

import (
	"errors"
	"github.com/notEpsilon/go-lru/list"
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
	if en, ok := c.store[key]; ok {
		c.order.MoveToFront(en)
		return en.Value.value, nil
	}
	return *new(V), keyNotFound
}

func (c *LRUCache[K, V]) Set(key K, value V) {
	if en, ok := c.store[key]; ok {
		en.Value.value = value
		c.order.MoveToFront(en)
		return
	}

	if len(c.store) == c.cap {
		toEvict := c.order.Back()
		delete(c.store, toEvict.Value.key)
		c.order.Remove(toEvict)
	}

	c.store[key] = c.order.PushFront(&entry[K, V]{key: key, value: value})
}

func (c *LRUCache[K, V]) Contains(key K) bool {
	_, ok := c.store[key]
	return ok
}

func (c *LRUCache[K, V]) Peek(key K) (V, error) {
	if en, ok := c.store[key]; ok {
		return en.Value.value, nil
	}
	return *new(V), keyNotFound
}
