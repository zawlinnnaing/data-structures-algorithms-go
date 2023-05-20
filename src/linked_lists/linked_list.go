package main

import "fmt"

type SinglyLinkedListNode struct {
	Data string
	Next *SinglyLinkedListNode
}

func (node *SinglyLinkedListNode) AddNode(nodeToAdd *SinglyLinkedListNode) {
	node.Next = nodeToAdd
}

func PrintNode(node *SinglyLinkedListNode) {
	tempNode := node
	for tempNode != nil {
		fmt.Print(tempNode.Data + " -> ")
		tempNode = tempNode.Next
	}
}

func SinglyLinkedListExample() {
	head := SinglyLinkedListNode(SinglyLinkedListNode{Data: "head"})
	node1 := SinglyLinkedListNode(SinglyLinkedListNode{Data: "node1"})
	node2 := SinglyLinkedListNode(SinglyLinkedListNode{Data: "node2"})
	head.AddNode(&node1)
	node1.AddNode(&node2)

	PrintNode(&head)
}

func main() {
	SinglyLinkedListExample()
}
