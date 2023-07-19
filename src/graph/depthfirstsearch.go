package graph

import "github.com/zawlinnnaing/data-structures-algorithms-go/src/slice"

type GraphEdge struct {
	To     int
	Weight int
}

type AdjacencyList = [][]GraphEdge

func walk(graph *AdjacencyList, current int, needle int, seen []bool, path *slice.Mutable[int]) bool {

	if seen[current] {
		return false
	}

	path.Push(current)

	if current == needle {
		return true
	}

	seen[current] = true

	for _, edge := range (*graph)[current] {
		if walk(graph, edge.To, needle, seen, path) {
			return true
		}
	}

	path.Pop()

	return false
}

func DepthFirstSearch(graph AdjacencyList, source int, needle int) []int {
	seen := make([]bool, len(graph))
	path := slice.NewMutable[int]()

	walk(&graph, source, needle, seen, &path)
	return path.Elements()
}
