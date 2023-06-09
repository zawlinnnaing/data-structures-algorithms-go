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

func (node *SinglyLinkedListNode) Length() int {
	nodeCount := 1
	tempNode := node
	for tempNode.Next != nil {
		tempNode = tempNode.Next
		nodeCount += 1
	}
	return nodeCount
}

func PrepareExample() [3]SinglyLinkedListNode {
	head := SinglyLinkedListNode(SinglyLinkedListNode{Data: "head"})
	node1 := SinglyLinkedListNode(SinglyLinkedListNode{Data: "node1"})
	node2 := SinglyLinkedListNode(SinglyLinkedListNode{Data: "node2"})

	head.AddNode(&node1).AddNode(&node2)
	return [3]SinglyLinkedListNode{head, node1, node2}
}

func RecursiveLength(node *SinglyLinkedListNode) int {
	if node == nil {
		return 0
	}
	return 1 + RecursiveLength(node.Next)
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
	nodes := PrepareExample()

	PrintNode(&nodes[0])
	fmt.Println("Length of linked list", nodes[0].Length())
	fmt.Println("Length of linked list in recursive way", RecursiveLength(&nodes[0]))
	fmt.Println()
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
	nodes := PrepareExample()
	node3 := SinglyLinkedListNode(SinglyLinkedListNode{Data: "node3"})

	InsertEnd(&nodes[0], &node3)

	PrintNode(&nodes[0])
}
