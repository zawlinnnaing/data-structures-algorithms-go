package doubly

import "fmt"

type DoublyNode[T comparable] struct {
	next *DoublyNode[T]
	prev *DoublyNode[T]
	data T
}

type DoublyLinkedList[T comparable] struct {
	length int
	head   *DoublyNode[T]
	tail   *DoublyNode[T]
}

func (linkedList *DoublyLinkedList[T]) Prepend(value T) {
	linkedList.length++
	node := DoublyNode[T]{data: value}

	if linkedList.head == nil {
		linkedList.head = &node
		linkedList.tail = &node
		return
	}

	node.next = linkedList.head
	linkedList.head.prev = &node
	linkedList.head = &node
}

func (this *DoublyLinkedList[T]) Append(value T) {
	this.length++
	node := DoublyNode[T]{data: value}
	if this.tail == nil {
		this.head = &node
		this.tail = &node
		return
	}

	this.tail.next = &node
	node.prev = this.tail
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
	if index > this.length {
		panic(fmt.Sprintf("Index must not be greater than length: %v.", this.length))
	}
	this.length++
	node := DoublyNode[T]{data: value}
	current := this.getAt(index)

	node.next = &current
	node.prev = current.prev

	current.prev = &node
	node.prev.next = &node
}

func (this *DoublyLinkedList[T]) RemoveAt(index int) (node DoublyNode[T]) {
	if this.head == nil || this.tail == nil {
		panic("List is empty")
	}

	if index > this.length {
		panic("Remove index is greater than list length")
	}
	this.length--
	node = this.getAt(index)

	this.removeNode(&node)

	return node
}

func (this *DoublyLinkedList[T]) Remove(value T) DoublyNode[T] {
	if this.length == 0 {
		panic("list is empty")
	}
	node := this.head
	for i := 0; i < this.length; i++ {
		if node.data == value {
			break
		}
		node = node.next
	}
	if node.data != value {
		panic(fmt.Sprintf("Cannot find value: %v in the list", value))
	}
	this.length--
	this.removeNode(node)

	return *node
}

func (this *DoublyLinkedList[T]) String() (result string) {
	curr := this.head
	for i := 0; i < this.length; i++ {
		result += fmt.Sprintf(" [%v] ", curr.data)
		if i != this.length-1 {
			result += "->"
		}
		curr = curr.next
	}
	return
}

func (this *DoublyLinkedList[T]) getAt(index int) DoublyNode[T] {
	current := this.head
	for i := 0; i < this.length; i++ {
		if i == index {
			break
		}
		current = current.next
	}
	return *current
}

func (this *DoublyLinkedList[T]) removeNode(node *DoublyNode[T]) {
	if node.next != nil {
		node.next.prev = node.prev
	}
	if node.prev != nil {
		node.prev.next = node.next
	}
	if node == this.head {
		this.head = node.next
	}
	if node == this.tail {
		this.tail = node.prev
	}
	if this.length == 0 {
		this.head = nil
		this.tail = nil
	}
}

func NewDoublyLinkedList[T comparable]() *DoublyLinkedList[T] {
	return &DoublyLinkedList[T]{
		length: 0,
	}
}
