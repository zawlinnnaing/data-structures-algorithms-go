package hashmap

import (
	"errors"

	"github.com/zawlinnnaing/data-structures-algorithms-go/src/linkedlists/doubly"
)

type LRU[K comparable, V comparable] struct {
	length        int
	lookup        map[K]*doubly.DoublyNode[V]
	reverseLookup map[*doubly.DoublyNode[V]]K
	head          *doubly.DoublyNode[V]
	tail          *doubly.DoublyNode[V]
	capacity      int
}

func (cache *LRU[K, V]) Get(key K) (value V, err error) {
	node := cache.lookup[key]
	if node == nil {
		return value, errors.New("Value not found")
	}
	cache.detach(node)
	cache.prepend(node)
	return node.Data, err
}

func (cache *LRU[K, V]) Update(key K, value V) {
	node := cache.lookup[key]
	if node == nil {
		node = doubly.NewNode[V](value)
		cache.prepend(node)
		cache.length += 1
		cache.lookup[key] = node
		cache.reverseLookup[node] = key
		cache.trimCache()
	} else {
		cache.detach(node)
		cache.prepend(node)
		node.Data = value
	}
}

func (cache *LRU[K, V]) detach(node *doubly.DoublyNode[V]) {
	if node.Prev != nil {
		node.Prev.Next = node.Next
	}
	if node.Next != nil {
		node.Next.Prev = node.Prev
	}
	if cache.head == node {
		cache.head = node.Next
	}

	if cache.tail == node {
		cache.tail = node.Prev
	}
	node.Next = nil
	node.Prev = nil
}

func (cache *LRU[K, V]) prepend(node *doubly.DoublyNode[V]) {
	if cache.head == nil {
		cache.head = node
		cache.tail = node
		return
	}
	node.Next = cache.head
	cache.head.Prev = node
}

func (cache *LRU[K, V]) trimCache() {
	if cache.length <= cache.capacity {
		return
	}
	key := cache.reverseLookup[cache.tail]
	delete(cache.lookup, key)
	delete(cache.reverseLookup, cache.tail)
	cache.detach(cache.tail)
}

func NewLRU[K comparable, V comparable](capacity int) *LRU[K, V] {
	return &LRU[K, V]{
		length:        0,
		capacity:      capacity,
		reverseLookup: make(map[*doubly.DoublyNode[V]]K),
		lookup:        make(map[K]*doubly.DoublyNode[V]),
	}
}
