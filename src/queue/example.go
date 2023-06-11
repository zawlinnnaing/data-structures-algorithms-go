package queue

import (
	"fmt"
	"log"
)

func runLRUCacheExample() {
	cache, cacheCreationError := CreateLRUCache(4)
	if cacheCreationError != nil {
		log.Fatal("Cache creation failed:", cacheCreationError.Error())
	}
	itemsToPush := []int{2, 3, 4, 3, 6, 9, 3}
	for _, value := range itemsToPush {
		cache.Push(value)
		fmt.Printf("After pushing %d: %v \n", value, cache.data)
	}
}

func RunQueueExamples() {
	fmt.Println("Running queue examples:")
	runLRUCacheExample()
}
