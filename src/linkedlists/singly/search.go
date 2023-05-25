package singly

import "fmt"

func SearchNode(value string, node SinglyLinkedListNode) SinglyLinkedListNode {
	var foundNode SinglyLinkedListNode
	tempNode := &node
	for tempNode.Next != nil {
		if tempNode.Data == value {
			foundNode = *tempNode
			break
		}
		tempNode = tempNode.Next
	}
	return foundNode
}

func SearchNodeExample() {
	head := SinglyLinkedListNode(SinglyLinkedListNode{Data: "head"})
	node1 := SinglyLinkedListNode(SinglyLinkedListNode{Data: "node1"})
	node2 := SinglyLinkedListNode(SinglyLinkedListNode{Data: "node2"})

	head.AddNode(&node1).AddNode(&node2)

	foundNode := SearchNode("node1", head)
	notFoundNode := SearchNode("bluh", head)

	fmt.Println(`Found node`, foundNode)
	fmt.Println("Not found node", notFoundNode)
}
