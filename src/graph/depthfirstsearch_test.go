package graph_test

import (
	"testing"

	"github.com/zawlinnnaing/data-structures-algorithms-go/src/graph"
	"github.com/zawlinnnaing/data-structures-algorithms-go/src/util"
)

func TestDepthFirstSearch(t *testing.T) {
	exampleGraph := graph.AdjacencyList([][]graph.GraphEdge{
		{
			graph.GraphEdge{To: 1},
		},
		{
			graph.GraphEdge{To: 2},
		},
		{},
	})
	path := graph.DepthFirstSearch(exampleGraph, 0, 2)
	expected := []int{0, 1, 2}
	if !util.SliceEquals[int](path, expected) {
		t.Errorf("Expected path %v, received path %v", expected, path)
	}
}
