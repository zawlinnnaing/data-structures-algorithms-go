package doubly

import "fmt"

type DoublyNode[T comparable] struct {
	Next *DoublyNode[T]
	Prev *DoublyNode[T]
	Data T
}

type DoublyLinkedList[T comparable] struct {
	length int
	head   *DoublyNode[T]
	tail   *DoublyNode[T]
}

func (linkedList *DoublyLinkedList[T]) Prepend(value T) {
	linkedList.length++
	node := DoublyNode[T]{Data: value}

	if linkedList.head == nil {
		linkedList.head = &node
		linkedList.tail = &node
		return
	}

	node.Next = linkedList.head
	linkedList.head.Prev = &node
	linkedList.head = &node
}

func (this *DoublyLinkedList[T]) Append(value T) {
	this.length++
	node := DoublyNode[T]{Data: value}
	if this.tail == nil {
		this.head = &node
		this.tail = &node
		return
	}

	this.tail.Next = &node
	node.Prev = this.tail
	this.tail = &node
}

func (this *DoublyLinkedList[T]) InsertAt(value T, index int) {
	if index == 0 {
		this.Prepend(value)
		return
	}
	if index == this.length {
		this.Append(value)
		return
	}
	if index >= this.length {
		panic(fmt.Sprintf("Index must not be greater than length: %v.", this.length))
	}
	this.length++
	node := DoublyNode[T]{Data: value}
	current := this.getAt(index)

	node.Next = current
	node.Prev = current.Prev

	current.Prev = &node
	node.Prev.Next = &node
}

func (this *DoublyLinkedList[T]) RemoveAt(index int) DoublyNode[T] {
	if this.head == nil || this.tail == nil {
		panic("List is empty")
	}

	if index >= this.length {
		panic("Remove index is greater than list length")
	}
	node := this.getAt(index)

	this.removeNode(node)

	return *node
}

func (this *DoublyLinkedList[T]) Remove(value T) DoublyNode[T] {
	if this.length == 0 {
		panic("list is empty")
	}
	node := this.head
	for i := 0; i < this.length; i++ {
		if node.Data == value {
			break
		}
		node = node.Next
	}
	if node == nil {
		panic(fmt.Sprintf("Cannot find value: %v in the list", value))
	}
	this.removeNode(node)

	return *node
}

func (this *DoublyLinkedList[T]) String() (result string) {
	curr := this.head
	for i := 0; i < this.length; i++ {
		// fmt.Println("curr", curr)
		result += fmt.Sprintf(" [%v] ", curr.Data)
		if i != this.length-1 {
			result += "->"
		}
		curr = curr.Next
	}
	return
}

func (this *DoublyLinkedList[T]) Get(index int) (value T) {
	node := this.getAt(index)
	if node == nil {
		return
	}
	return node.Data
}

func (this *DoublyLinkedList[T]) getAt(index int) *DoublyNode[T] {
	current := this.head
	for i := 0; i < index; i++ {
		current = current.Next
	}
	return current
}

func (this *DoublyLinkedList[T]) removeNode(node *DoublyNode[T]) {
	this.length--
	if node == this.head {
		this.head = node.Next
	}
	if node == this.tail {
		this.tail = node.Prev
	}
	if node.Next != nil {
		node.Next.Prev = node.Prev
	}
	if node.Prev != nil {
		node.Prev.Next = node.Next
	}
	if this.length == 0 {
		this.head = nil
		this.tail = nil
	}
	node.Next = nil
	node.Prev = nil
}

func NewDoublyLinkedList[T comparable]() *DoublyLinkedList[T] {
	return &DoublyLinkedList[T]{
		length: 0,
	}
}

func NewNode[V comparable](value V) *DoublyNode[V] {
	return &DoublyNode[V]{
		Data: value,
	}
}
