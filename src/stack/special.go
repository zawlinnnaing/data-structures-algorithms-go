package stack

import (
	"errors"
)

type SpecialStack struct {
	Stack[int]
	minPointersStack Stack[int]
}

func (stack *SpecialStack) GetMinValue() (int, error) {
	if stack.IsEmpty() || stack.minPointersStack.IsEmpty() {
		return 0, errors.New("Stack is empty")
	}
	minPointer := stack.minPointersStack.Peak()
	return stack.elements[minPointer], nil
}

func (stack *SpecialStack) Push(value int) {
	stack.Stack.Push(value)
	if stack.minPointersStack.IsEmpty() {
		stack.minPointersStack.Push(stack.topPointer)
		return
	}
	minValue := stack.minPointersStack.Peak()
	if value < minValue {
		stack.minPointersStack.Push(stack.topPointer)
	}
}

func (stack *SpecialStack) Pop() (int, error) {

	if stack.topPointer == stack.minPointersStack.Peak() {
		stack.minPointersStack.Pop()
	}

	poppedValue, error := stack.Stack.Pop()

	return poppedValue, error
}

func CreateSpecialStack() SpecialStack {
	specialStack := SpecialStack{Stack: *CreateStack[int](nil), minPointersStack: *CreateStack[int](nil)}
	return specialStack
}
