package singly

// Time complexity: O(n), Traversing every node
// Space complexity: O(1), Only three pointers needed
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

// Algorithm explanation: https://www.geeksforgeeks.org/reverse-a-linked-list/#:~:text=Reverse%20a%20linked%20list%20using%20Recursion%3A
// Diagram: https://media.geeksforgeeks.org/wp-content/cdn-uploads/2009/07/Linked-List-Rverse.gif
// Time complexity: O(n), Traversing every node
// Space complexity: O(n), Function call stack = Number of nodes
func ReverseLinkedListRecursive(head *SinglyLinkedListNode) *SinglyLinkedListNode {
	if head == nil || head.Next == nil {
		return head
	}
	rest := ReverseLinkedListRecursive(head.Next)
	head.Next.Next = head
	head.Next = nil
	return rest
}

func ReverseLinkedListExample() {
	nodes := PrepareExample()
	reversedNode := ReverseLinkedList(&nodes[0])
	PrintNode(reversedNode)
}

func ReverseLinkedListRecursiveExample() {
	nodes := PrepareExample()
	reversedNode := ReverseLinkedListRecursive(&nodes[0])
	PrintNode(reversedNode)
}
