package main

import (
	"fmt"
)

type Cache interface {
	get(key string) interface{}
	set(key string, value interface{})
	clear()
}

type LRU_Value = interface{}

type Node struct {
	value LRU_Value
	prev  *Node
	next  *Node
	key   string
}

type LRU_Map map[string]*Node

type LRU_Cache struct {
	Head     *Node
	Last     *Node
	lru_map  LRU_Map
	capacity int
}

func createLRU_Cache(capacity int) LRU_Cache {
	return LRU_Cache{
		Head:     nil,
		lru_map:  make(LRU_Map, capacity),
		capacity: capacity,
	}
}

func (cache *LRU_Cache) get(key string) LRU_Value {
	value, ok := cache.lru_map[key]
	if !ok {
		return nil
	}

	cache.reEvaluate(key, false)
	return value.value
}

func (cache *LRU_Cache) set(key string, value LRU_Value) {
	_, ok := cache.lru_map[key]

	// element donest exist
	if !ok {
		var node = Node{prev: cache.Last, key: key, value: value}
		cache.lru_map[key] = &node

		// non existing previous element
		if node.prev != nil {
			node.prev.next = &node
		}

		// empty cache
		if cache.Head == nil {
			cache.Head = &node
		}

		cache.reEvaluate(key, true)
		// element exists
	} else {
		cache.lru_map[key].value = value
		cache.reEvaluate(key, false)
	}
}

func (cache *LRU_Cache) reEvaluate(key string, isNew bool) {
	// empty or one element
	if cache.Head == cache.Last {
		return
	}

	if isNew {
		// capacity full - delewte first
		if cache.capacity == len(cache.lru_map)+1 {
			cache.Head = cache.Head.next
			cache.Head.prev = nil
		}

		cache.Last = cache.lru_map[key]
	} else {
		var current = cache.lru_map[key]

		// last element
		if cache.Last == current {
			return
		}
		// not head
		var prev = cache.lru_map[key].prev
		var next = cache.lru_map[key].next

		Iterate(cache)

		if cache.Head != current {
			prev.next = next
			if next != nil {
				next.prev = prev
			}
		} else {
			cache.Head = next
		}

		cache.Last.next = current
		current.prev = cache.Last
		cache.Last = current
	}
}

func Iterate(cache *LRU_Cache) {
	for temp := cache.Head; temp != nil; {
		fmt.Println(temp.value)
		temp = temp.next
	}
}
