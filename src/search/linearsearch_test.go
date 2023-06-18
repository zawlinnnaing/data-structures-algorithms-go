package search_test

import (
	"testing"

	"github.com/zawlinnnaing/data-structures-algorithms-go/src/search"
)

func TestLinearSearch(t *testing.T) {
	hackStack := []int{1, 2, 3, 5, 6}
	needle := 3
	shouldFound := search.LinearSearch(hackStack, needle)

	if !shouldFound {
		t.Errorf("Found must be true; got %v", shouldFound)
	}
	nonExistsNeedle := 4
	shouldNotFound := search.LinearSearch(hackStack, nonExistsNeedle)
	if shouldNotFound {
		t.Errorf("Should not found needle %d; got %v", nonExistsNeedle, shouldFound)
	}
}
