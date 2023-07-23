package hashmap_test

import (
	"testing"

	"github.com/zawlinnnaing/data-structures-algorithms-go/src/hashmap"
)

func TestLRU(t *testing.T) {
	cache := hashmap.NewLRU[string, int](5)
	result, err := cache.Get("test")
	if err == nil {
		t.Errorf("Expected error, received %v", result)
	}
	cache.Update("test", 42)
	result, _ = cache.Get("test")
	if result != 42 {
		t.Errorf("Expected %v, Received %v", 42, result)
	}
	cache.Update("test", 44)
	result, _ = cache.Get("test")
	if result != 44 {
		t.Errorf("Expected %v, Received %v", 44, result)
	}
	cache.Update("test2", 340)
	cache.Update("test3", 340)
	cache.Update("test4", 340)
	cache.Update("test5", 340)
	cache.Update("test6", 340)
	result, err = cache.Get("test")
	if err == nil {
		t.Errorf("Expected error, received %v", result)
	}
}
