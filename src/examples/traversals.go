package examples

import (
	"fmt"

	"github.com/zawlinnnaing/data-structures-algorithms-go/src/tree"
	"github.com/zawlinnnaing/data-structures-algorithms-go/src/tree/traversals"
)

func BinaryNodeTraversalsExample() {
	rootNode := tree.BinaryNode[int]{Value: 1}
	rootNode.PushLeft(2)
	rootNode.PushRight(3)
	rootNode.Left.PushLeft(4)
	rootNode.Left.PushRight(5)
	rootNode.Right.PushLeft(6)
	rootNode.Right.PushRight(7)
	fmt.Println("Pre order traversal:", traversals.PreOrder[int](rootNode))
	fmt.Println("In order traversal:", traversals.InOrder[int](rootNode))
	fmt.Println("Post order traversal:", traversals.PostOrder[int](rootNode))
}
