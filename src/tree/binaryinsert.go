package tree

func BinaryTreeInsert(binaryNode *BinaryNode[int], value int) {
	if binaryNode.Value < value {
		if binaryNode.Right == nil {
			insertNode := BinaryNode[int]{Value: value}
			binaryNode.Right = &insertNode
			return
		}
		BinaryTreeInsert(binaryNode.Right, value)
	} else {
		if binaryNode.Left == nil {
			insertNode := BinaryNode[int]{Value: value}
			binaryNode.Left = &insertNode
			return
		}
		BinaryTreeInsert(binaryNode.Left, value)
	}
}
