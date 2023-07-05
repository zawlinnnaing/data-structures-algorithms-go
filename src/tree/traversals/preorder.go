package traversals

import (
	"github.com/zawlinnnaing/data-structures-algorithms-go/src/slice"
	"github.com/zawlinnnaing/data-structures-algorithms-go/src/tree"
)

func walkPreOrder[T any](node *tree.BinaryNode[T], path *slice.Mutable[T]) {

	if node == nil {
		return
	}

	path.Push(node.Value)
	walkPreOrder[T](node.Left, path)
	walkPreOrder[T](node.Right, path)
}

func PreOrder[T any](node tree.BinaryNode[T]) slice.Mutable[T] {
	path := slice.NewMutable[T]()
	walkPreOrder[T](&node, &path)
	return path
}
