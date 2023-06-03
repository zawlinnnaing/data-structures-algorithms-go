package singly

import "fmt"

func GetNodeFromEnd(head *SinglyLinkedListNode, index int) *SinglyLinkedListNode {
	listLen := head.Length()
	listIndexFromStart := listLen - index - 1
	current := head
	for i := 1; i <= listIndexFromStart; i++ {
		current = current.Next
	}
	return current
}

func GetNodeFromEndExample() {
	head := SinglyLinkedListNode(SinglyLinkedListNode{Data: "head"})
	node1 := SinglyLinkedListNode(SinglyLinkedListNode{Data: "node1"})
	node2 := SinglyLinkedListNode(SinglyLinkedListNode{Data: "node2"})
	node3 := SinglyLinkedListNode(SinglyLinkedListNode{Data: "node3"})
	head.AddNode(&node1).AddNode(&node2).AddNode(&node3)
	node := GetNodeFromEnd(&head, 2) // Should return: node1
	fmt.Printf("Node from end: %s \n", node.Data)
}
