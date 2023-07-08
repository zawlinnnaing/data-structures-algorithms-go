package tree_test

import (
	"testing"

	"github.com/zawlinnnaing/data-structures-algorithms-go/src/tree"
)

func TestBinarySearch(t *testing.T) {
	binaryNode := getBinaryTree()
	result := tree.BinaryTreeSearch(&binaryNode, 25)

	if !result {
		t.Errorf("Expected true: received %v", result)
	}

	result = tree.BinaryTreeSearch(&binaryNode, 23)
	if result {
		t.Errorf("Expected false, received %v", result)
	}
}

func TestBinaryInsert(t *testing.T) {
	binaryNode := getBinaryTree()
	tree.BinaryTreeInsert(&binaryNode, 23)
	if binaryNode.Right.Left.Right == nil {
		t.Errorf("Failed to insert binary node")
	}
}

func getBinaryTree() tree.BinaryNode[int] {
	binaryNode := tree.BinaryNode[int]{
		Value: 18,
	}
	binaryNode.PushLeft(15)
	binaryNode.Left.PushLeft(13)
	binaryNode.Left.PushRight(16)
	binaryNode.PushRight(25)
	binaryNode.Right.PushLeft(20)
	binaryNode.Right.PushRight(30)
	return binaryNode
}
