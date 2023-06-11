package queue

import (
	"errors"
)

type LRUCache struct {
	data             []int
	maximumCacheSize int
}

func (cache *LRUCache) Push(item int) {
	itemIndex := cache.findIndex(item)
	if itemIndex < 0 {
		if len(cache.data) >= cache.maximumCacheSize {
			cache.Pop()
		}
	} else {
		cache.data = append(cache.data[:itemIndex], cache.data[itemIndex+1:]...)
	}

	cache.data = append(cache.data, item)
}

func (cache *LRUCache) Pop() (int, error) {
	if len(cache.data) <= 0 {
		return 0, errors.New("cache is empty")
	}
	poppedValue := cache.data[0]
	cache.data = cache.data[1:len(cache.data)]
	return poppedValue, nil
}

func (cache *LRUCache) findIndex(item int) int {
	for index, value := range cache.data {
		if value == item {
			return index
		}
	}
	return -1
}
func CreateLRUCache(cacheSize int) (*LRUCache, error) {
	if cacheSize <= 0 {
		return nil, errors.New("cacheSize must be greater than 0")
	}
	lruCache := LRUCache{maximumCacheSize: cacheSize, data: nil}
	return &lruCache, nil
}
