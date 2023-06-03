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
	for current.Next != nil {
		prev = current
		current = current.Next
	}
	prev.Next = nil
}

func DeleteNodeRecursivelyByValue(head *SinglyLinkedListNode, prev *SinglyLinkedListNode, value string) {
	if head == nil {
		return
	}
	if head.Data == value {
		head = head.Next
		prev.Next = head
		return
	}
	prev = head
	DeleteNodeRecursivelyByValue(head.Next, prev, value)
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
	fmt.Println("Running delete node recursively:")
	nodesToDeleteRecursively := PrepareExample()
	DeleteNodeRecursivelyByValue(&nodesToDeleteRecursively[0], nil, "node1")
	PrintNode(&nodesToDeleteRecursively[0])
}
