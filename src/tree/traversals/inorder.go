package traversals

import (
	"github.com/zawlinnnaing/data-structures-algorithms-go/src/slice"
	"github.com/zawlinnnaing/data-structures-algorithms-go/src/tree"
)

func walkInOrder[T any](node *tree.BinaryNode[T], path *slice.Mutable[T]) {

	if node == nil {
		return
	}

	walkInOrder[T](node.Left, path)
	path.Push(node.Value)
	walkInOrder[T](node.Right, path)

}

func InOrder[T any](node tree.BinaryNode[T]) slice.Mutable[T] {
	path := slice.NewMutable[T]()
	walkInOrder[T](&node, &path)
	return path
}
