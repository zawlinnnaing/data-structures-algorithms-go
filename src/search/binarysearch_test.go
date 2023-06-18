package search_test

import (
	"testing"

	"github.com/zawlinnnaing/data-structures-algorithms-go/src/search"
)

func TestBinarySearch(t *testing.T) {
	hayStack := []int{1, 3, 5, 7, 9, 11, 13}
	shouldFound := search.BinarySearch(hayStack, 5)
	if !shouldFound {
		t.Errorf("Should found value 5, got %v", shouldFound)
	}
	shouldNotFound := search.BinarySearch(hayStack, 6)
	if shouldNotFound {
		t.Errorf("Should not found value 6, got %v", shouldNotFound)
	}
}

func TestTwoCrystalBallsProblem(t *testing.T) {
	results := []bool{false, false, false, true, true, true, true}
	breakIndex := 3
	result := search.TwoCrystalBallsProblem(results)
	if result != breakIndex {
		t.Errorf("Break index should be %d; got %d", breakIndex, result)
	}
}
