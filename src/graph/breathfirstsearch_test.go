package graph_test

import (
	"testing"

	"github.com/zawlinnnaing/data-structures-algorithms-go/src/graph"
	"github.com/zawlinnnaing/data-structures-algorithms-go/src/util"
)

func TestBreathFirstSearch(t *testing.T) {
	matrix := graph.AdjacencyMatrix{
		[]int{0, 20, 0},
		[]int{0, 0, 40},
		[]int{0, 0, 0},
	}
	out := graph.BreathFirstSearch(matrix, 0, 2)
	expected := []int{0, 1, 2}
	if !util.SliceEquals[int](out, expected) {
		t.Errorf("Expected path to be %v, received %v", expected, out)
	}
}
