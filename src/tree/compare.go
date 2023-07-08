package tree

func CompareBinaryTrees[T comparable](treeA *BinaryNode[T], treeB *BinaryNode[T]) bool {
	if treeA == nil && treeB == nil {
		return true
	}
	if treeA == nil || treeB == nil {
		return false
	}

	if treeA.Value != treeB.Value {
		return false
	}

	return CompareBinaryTrees[T](treeA.Left, treeB.Left) && CompareBinaryTrees[T](treeA.Right, treeB.Right)
}
