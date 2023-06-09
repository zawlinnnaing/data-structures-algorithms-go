package stack

import (
	"errors"
	"fmt"
)

type stackEntry interface {
	string | int
}

// TODO: use generic linked list to store elements
type Stack[T stackEntry] struct {
	topPointer int
	elements   []T
}

func (stack *Stack[T]) Pop() (T, error) {
	var value T
	if stack.IsEmpty() {
		return value, errors.New("Stack is Empty")
	}

	value = stack.elements[stack.topPointer]
	stack.elements = stack.elements[:len(stack.elements)-1]
	stack.topPointer -= 1
	return value, nil
}

func (stack *Stack[T]) Push(value T) {
	stack.topPointer += 1
	stack.elements = append(stack.elements, value)
}

func (stack *Stack[T]) Peak() T {
	var value T
	if stack.IsEmpty() {
		return value
	}
	return stack.elements[stack.topPointer]
}

func (stack *Stack[T]) Size() int {
	return stack.topPointer + 1
}

func (stack *Stack[T]) IsEmpty() bool {
	return stack.topPointer < 0
}

func (stack *Stack[T]) PrintStack(prefix string) {
	for i := 0; i < stack.Size(); i++ {
		fmt.Println(prefix + fmt.Sprintf("%v", stack.elements[i]))
	}
}

func CreateStack[T stackEntry](initialElements []T) *Stack[T] {
	stack := Stack[T]{topPointer: len(initialElements) - 1, elements: initialElements}
	return &stack
}
