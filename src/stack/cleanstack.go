package stack

import "fmt"

type Node[T any] struct {
	value T
	next  *Node[T]
}

type CleanStack[T any] struct {
	head   *Node[T]
	length int
}

func (stack *CleanStack[T]) Push(value T) {
	node := &Node[T]{value: value}
	stack.length += 1
	if stack.head == nil {
		stack.head = node
		return
	}
	node.next = stack.head
	stack.head = node
}

func (stack *CleanStack[T]) Pop() (value T) {
	if stack.head == nil {
		return value
	}
	stack.length -= 1
	temp := stack.head
	stack.head = stack.head.next
	return temp.value
}

func (stack *CleanStack[T]) Peek() (value T) {
	if stack.head == nil {
		return value
	}
	return stack.head.value
}

func (stack *CleanStack[T]) Length() int {
	return stack.length
}

func (stack *CleanStack[T]) String() string {
	if stack.head == nil {
		return "<empty stack>"
	}
	str := ""
	temp := stack.head
	for i := 0; i < stack.length; i++ {
		str += fmt.Sprintf("%v\n", temp.value)
		temp = temp.next
	}
	return str
}

func NewStack[T any](initialValues ...T) *CleanStack[T] {
	stack := CleanStack[T]{length: 0}
	for _, value := range initialValues {
		stack.Push(value)
	}
	return &stack
}

func RunStackExample() {
	exampleStack := NewStack[int](2, 3, 4, 5, 6)
	stackLength := exampleStack.length
	fmt.Println("Stack after pushing \n", exampleStack)
	for i := 0; i < stackLength; i++ {
		fmt.Println("Popped value", exampleStack.Pop())
	}
	fmt.Println("Stack after popping:", exampleStack)
}
