package sort_test

import (
	"testing"

	"github.com/zawlinnnaing/data-structures-algorithms-go/src/sort"
	"github.com/zawlinnnaing/data-structures-algorithms-go/src/util"
)

func TestQuick(t *testing.T) {
	input := []int{6, 5, 4, 3, 2, 1}

	sort.Quick(input)

	expected := []int{1, 2, 3, 4, 5, 6}

	if !util.SliceEquals[int](input, expected) {
		t.Errorf("Expected %v, received %v", expected, input)
	}
}
