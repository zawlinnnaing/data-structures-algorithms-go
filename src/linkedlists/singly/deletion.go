package singly

import "fmt"

func DeleteHead(head *SinglyLinkedListNode) *SinglyLinkedListNode {
	temp := head.Next
	head.Next = nil
	return temp
}

func DeleteMiddle(head *SinglyLinkedListNode) {}

func DeleteExample() {
	deleteHeadNodes := PrepareExample()
	remainingNodes := DeleteHead(&deleteHeadNodes[0])
	fmt.Println("Running delete head example:")
	PrintNode(remainingNodes)
}
