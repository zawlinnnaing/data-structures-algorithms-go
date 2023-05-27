package singly

func ReverseLinkedList(node *SinglyLinkedListNode) *SinglyLinkedListNode {
	current := node
	var prev *SinglyLinkedListNode
	var next *SinglyLinkedListNode
	for current != nil {
		next = current.Next
		current.Next = prev
		prev = current
		current = next
	}
	node = prev
	return node
}

func ReverseLinkedListExample() {
	head := SinglyLinkedListNode(SinglyLinkedListNode{Data: "head"})
	node1 := SinglyLinkedListNode(SinglyLinkedListNode{Data: "node1"})
	node2 := SinglyLinkedListNode(SinglyLinkedListNode{Data: "node2"})

	head.AddNode(&node1).AddNode(&node2)

	ReverseLinkedList(&head)
	PrintNode(&node2)
}
