package singly

import "fmt"

func DeleteHead(head *SinglyLinkedListNode) *SinglyLinkedListNode {
	temp := head.Next
	head.Next = nil
	return temp
}

func DeleteMiddle(head *SinglyLinkedListNode, position int) *SinglyLinkedListNode {
	prev := head
	current := head
	for i := 0; i <= position; i += 1 {
		if i == position && current != nil {
			prev.Next = current.Next
			current.Next = nil
			break
		}
		prev = current
		if prev == nil {
			break
		}
		current = current.Next

	}
	return head

}

func DeleteEnd(head *SinglyLinkedListNode) {
	prev := head
	current := head
	for prev.Next != nil {
		// Meaning current node is last node
		if current.Next == nil {
			prev.Next = nil
			break
		}
		prev = current
		current = current.Next
	}
}

func DeleteExample() {
	deleteHeadNodes := PrepareExample()
	remainingNodes := DeleteHead(&deleteHeadNodes[0])
	fmt.Println("Running delete head example:")
	PrintNode(remainingNodes)
	deleteMiddleNodes := PrepareExample()
	fmt.Println("Running delete middle example:")
	DeleteMiddle(&deleteMiddleNodes[0], 1)
	PrintNode(&deleteMiddleNodes[0])
	fmt.Println("Running delete end example:")
	deleteEndNodes := PrepareExample()
	DeleteEnd(&deleteEndNodes[0])
	PrintNode(&deleteEndNodes[0])
}
