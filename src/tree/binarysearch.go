package tree

func BinaryTreeSearch(binaryNode *BinaryNode[int], value int) bool {
	if binaryNode == nil {
		return false
	}
	if binaryNode.Value == value {
		return true
	}
	if value > binaryNode.Value {
		return BinaryTreeSearch(binaryNode.Right, value)
	}

	return BinaryTreeSearch(binaryNode.Left, value)
}
