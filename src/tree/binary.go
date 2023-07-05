package tree

import "fmt"

type BinaryNode[T any] struct {
	Value T
	Left  *BinaryNode[T]
	Right *BinaryNode[T]
}

func (this *BinaryNode[T]) PushLeft(value T) *BinaryNode[T] {
	node := BinaryNode[T]{Value: value}
	this.Left = &node
	return this
}

func (this *BinaryNode[T]) PushRight(value T) *BinaryNode[T] {
	node := BinaryNode[T]{Value: value}
	this.Right = &node
	return this
}

func NewBinaryNodeFromSlice[T any](values [][3]T) (rootNode *BinaryNode[T]) {
	if len(values) == 0 {
		return nil
	}
	shouldLeft := false
	var prevNode *BinaryNode[T]
	for index, tuples := range values {
		node1, node2, node3 := BinaryNode[T]{Value: tuples[0]}, BinaryNode[T]{Value: tuples[1]}, BinaryNode[T]{Value: tuples[2]}
		node1.Left = &node2
		node1.Right = &node3
		if index == 0 {
			rootNode = &node1
		}
		if prevNode != nil {
			if shouldLeft {
				prevNode.Left = &node1
			} else {
				prevNode.Right = &node1
			}
		}
		prevNode = &node1
		shouldLeft = !shouldLeft
	}
	fmt.Println("root node", rootNode.Left.Value)
	return
}
