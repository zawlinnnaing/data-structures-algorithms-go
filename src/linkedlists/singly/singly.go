package singly

import (
	"errors"
	"fmt"
)

type SinglyLinkedListNode struct {
	Data string
	Next *SinglyLinkedListNode
}

func (node *SinglyLinkedListNode) AddNode(nodeToAdd *SinglyLinkedListNode) *SinglyLinkedListNode {
	node.Next = nodeToAdd
	return nodeToAdd
}

func PrintNode(node *SinglyLinkedListNode) {
	tempNode := node
	for tempNode != nil {
		if tempNode.Next != nil {
			fmt.Print(tempNode.Data + " -> ")
		} else {
			fmt.Print(tempNode.Data)
		}
		tempNode = tempNode.Next
	}
	fmt.Print("\n")
}

func InsertHead(head *SinglyLinkedListNode, nodeToInert *SinglyLinkedListNode) {
	nodeToInert.Next = head
	head = nodeToInert
}

func InsertAfter(prevNode *SinglyLinkedListNode, nodeToInsert *SinglyLinkedListNode) (*SinglyLinkedListNode, error) {
	if prevNode == nil {
		return nil, errors.New("prevNode cannot be nil")
	}
	nodeToInsert.Next = prevNode.Next
	prevNode.Next = nodeToInsert
	return prevNode, nil
}

func InsertEnd(head *SinglyLinkedListNode, nodeToInsert *SinglyLinkedListNode) {
	if head == nil {
		head = nodeToInsert
		return
	}

	var last *SinglyLinkedListNode = head

	for last.Next != nil {
		last = last.Next
	}

	last.Next = nodeToInsert
}

func SimpleInsertExample() {
	head := SinglyLinkedListNode(SinglyLinkedListNode{Data: "head"})
	node1 := SinglyLinkedListNode(SinglyLinkedListNode{Data: "node1"})
	node2 := SinglyLinkedListNode(SinglyLinkedListNode{Data: "node2"})
	head.AddNode(&node1).AddNode(&node2)

	PrintNode(&head)
}

func InsertHeadExample() {
	head := SinglyLinkedListNode(SinglyLinkedListNode{Data: "head"})
	node1 := SinglyLinkedListNode(SinglyLinkedListNode{Data: "node1"})
	node2 := SinglyLinkedListNode(SinglyLinkedListNode{Data: "node2"})
	head.AddNode(&node1)

	InsertHead(&head, &node2)

	PrintNode(&node2)
}

func InsertAfterExample() {
	head := SinglyLinkedListNode(SinglyLinkedListNode{Data: "head"})
	node1 := SinglyLinkedListNode(SinglyLinkedListNode{Data: "node1"})
	node2 := SinglyLinkedListNode(SinglyLinkedListNode{Data: "node2"})

	nodeBetween1And2 := SinglyLinkedListNode(SinglyLinkedListNode{Data: "nodeBetween1And2"})

	head.AddNode(&node1).AddNode(&node2)
	InsertAfter(&node1, &nodeBetween1And2)

	PrintNode(&head)
}

func InsertEndExample() {
	head := SinglyLinkedListNode(SinglyLinkedListNode{Data: "head"})
	node1 := SinglyLinkedListNode(SinglyLinkedListNode{Data: "node1"})
	node2 := SinglyLinkedListNode(SinglyLinkedListNode{Data: "node2"})
	head.AddNode(&node1).AddNode(&node2)

	node3 := SinglyLinkedListNode(SinglyLinkedListNode{Data: "node3"})

	InsertEnd(&head, &node3)

	PrintNode(&head)
}
