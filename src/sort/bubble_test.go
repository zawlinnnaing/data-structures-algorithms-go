package sort_test

import (
	"testing"

	"github.com/zawlinnnaing/data-structures-algorithms-go/src/sort"
	"github.com/zawlinnnaing/data-structures-algorithms-go/src/util"
)

func TestBubbleSort(t *testing.T) {
	input := []int{3, 4, 2, 5, 8, 6}
	result := sort.BubbleSort(&input)
	sorted := []int{2, 3, 4, 5, 6, 8}
	if !util.SliceEquals[int](*result, sorted) {
		t.Errorf("Expected %v, Received %v", sorted, *result)
	}
}
