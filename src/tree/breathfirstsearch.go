package tree

import (
	"github.com/zawlinnnaing/data-structures-algorithms-go/src/queue"
)

func BreathFirstSearch[T comparable](binaryNode *BinaryNode[T], needle T) bool {
	nodeQueue := queue.NewQueue[*BinaryNode[T]](binaryNode)

	for nodeQueue.Length() > 0 {
		node := nodeQueue.Pop()

		if node.Value == needle {
			return true
		}

		if node.Left != nil {
			nodeQueue.Push(node.Left)
		}

		if node.Right != nil {
			nodeQueue.Push(node.Right)
		}
	}
	return false
}
