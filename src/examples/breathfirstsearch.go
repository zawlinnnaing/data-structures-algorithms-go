package examples

import (
	"fmt"

	"github.com/zawlinnnaing/data-structures-algorithms-go/src/tree"
)

func RunBreathFirstSearchExample() {
	binaryNode := tree.BinaryNode[int]{Value: 1}
	binaryNode.PushLeft(2)
	binaryNode.PushRight(3)
	binaryNode.Left.PushLeft(4)
	binaryNode.Left.PushRight(5)
	binaryNode.Right.PushLeft(6)
	binaryNode.Right.PushRight(7)

	result := tree.BreathFirstSearch[int](&binaryNode, 5)

	fmt.Println("Result is:", result)
}
