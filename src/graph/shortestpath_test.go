package graph_test

import (
	"testing"

	"github.com/zawlinnnaing/data-structures-algorithms-go/src/graph"
	"github.com/zawlinnnaing/data-structures-algorithms-go/src/util"
)

func TestShortestPath(t *testing.T) {
	inputGraph := graph.AdjacencyList{
		[]graph.GraphEdge{
			{
				To:     1,
				Weight: 1,
			},
			{
				To:     2,
				Weight: 5,
			},
		},
		[]graph.GraphEdge{
			{
				To:     2,
				Weight: 7,
			},
			{
				To:     3,
				Weight: 6,
			},
		},
		[]graph.GraphEdge{
			{
				To:     4,
				Weight: 1,
			},
		},
		{
			graph.GraphEdge{
				To:     2,
				Weight: 1,
			},
		},
		{},
	}
	expected := []int{0, 2, 4}
	result := graph.ShortestPath(inputGraph, 0, 4)
	if !util.SliceEquals[int](expected, result) {
		t.Fatalf("Expected %v, Received %v", expected, result)
	}
}
