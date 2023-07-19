package graph

import (
	"math"

	"github.com/zawlinnnaing/data-structures-algorithms-go/src/slice"
)

func hasUnvisitedNode(seen *[]bool) bool {
	for i := 0; i < len(*seen); i++ {
		if !(*seen)[i] {
			return true
		}
	}
	return false
}

func getLowestUnvisitedNode(seen *[]bool, distances *[]int) int {
	lowestDistance := math.MaxInt64
	lowestNode := -1
	for i, distance := range *distances {
		if (*seen)[i] {
			continue
		}
		if distance < lowestDistance {
			lowestDistance = distance
			lowestNode = i
		}
	}
	return lowestNode
}

func ShortestPath(inputGraph AdjacencyList, source int, sink int) []int {
	seen := make([]bool, len(inputGraph))
	prev := make([]int, len(inputGraph))
	distances := make([]int, len(inputGraph))

	for i := 0; i < len(inputGraph); i++ {
		seen[i] = false
		prev[i] = -1
		distances[i] = math.MaxInt64
	}

	seen[source] = true
	distances[source] = 0
	sourceEdges := inputGraph[source]
	for _, edge := range sourceEdges {
		distances[edge.To] = edge.Weight
		prev[edge.To] = source
	}

	for hasUnvisitedNode(&seen) {
		current := getLowestUnvisitedNode(&seen, &distances)
		if current == -1 {
			break
		}
		seen[current] = true
		edges := inputGraph[current]
		for _, edge := range edges {
			if seen[edge.To] {
				continue
			}
			dist := distances[current] + edge.Weight
			if dist < distances[edge.To] {
				distances[edge.To] = dist
				prev[edge.To] = current
			}
		}
	}

	out := slice.NewMutable[int]()
	current := sink
	for prev[current] != -1 {
		out.Push(current)
		current = prev[current]
	}
	out.Push(source)

	out.Reverse()

	return out.Elements()
}
