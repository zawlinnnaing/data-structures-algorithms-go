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

func DeleteExample() {
	deleteHeadNodes := PrepareExample()
	remainingNodes := DeleteHead(&deleteHeadNodes[0])
	fmt.Println("Running delete head example:")
	PrintNode(remainingNodes)
	deleteMiddleNodes := PrepareExample()
	fmt.Println("Running delete middle example:")
	DeleteMiddle(&deleteMiddleNodes[0], 1)
	PrintNode(&deleteMiddleNodes[0])
}
