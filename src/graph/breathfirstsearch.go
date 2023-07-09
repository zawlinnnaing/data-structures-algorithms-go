package graph

import (
	"github.com/zawlinnnaing/data-structures-algorithms-go/src/queue"
	"github.com/zawlinnnaing/data-structures-algorithms-go/src/slice"
)

type AdjustMatrix = [][]int

func BreathFirstSearch(matrix AdjustMatrix, source int, needle int) []int {
	queue := queue.NewQueue[int]()
	seen := make([]bool, len(matrix))
	prev := make([]int, len(matrix))
	seen[source] = true
	queue.Push(source)

	for idx := range prev {
		prev[idx] = -1
	}

	for queue.Length() > 0 {
		current := queue.Pop()
		if current == needle {
			break
		}
		connections := matrix[current]
		for childIdx := range connections {
			if seen[childIdx] {
				continue
			}
			if connections[childIdx] == 0 {
				continue
			}
			queue.Push(childIdx)
			seen[childIdx] = true
			prev[childIdx] = current
		}
	}

	if prev[needle] == -1 {
		return []int{}
	}

	current := needle
	out := slice.NewMutable[int]()
	for prev[current] != -1 {
		out.Push(current)
		current = prev[current]
	}

	out.Reverse()

	return append([]int{source}, out.Elements()...)
}
