package traversals

import (
	"github.com/zawlinnnaing/data-structures-algorithms-go/src/slice"
	"github.com/zawlinnnaing/data-structures-algorithms-go/src/tree"
)

func walkPostOrder[T any](node *tree.BinaryNode[T], path *slice.Mutable[T]) {

	if node == nil {
		return
	}

	walkPostOrder[T](node.Left, path)
	walkPostOrder[T](node.Right, path)
	path.Push(node.Value)

}

func PostOrder[T any](node tree.BinaryNode[T]) slice.Mutable[T] {
	path := slice.NewMutable[T]()
	walkPostOrder[T](&node, &path)
	return path
}
