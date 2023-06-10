package stack

/**
* Algorithm reference: https://www.geeksforgeeks.org/design-and-implement-special-stack-data-structure/#:~:text=Further%20optimized%20O(1)%20time%20complexity%20and%20O(1)%20space%20complexity%20solution%20%3A
**/

const dummyValue = 9999999

type SpaceOptimizedSpecialStack struct {
	minValue     int
	wrapperStack Stack[int]
}

func getPushedValue(value int, minValue int) int {
	return (value * dummyValue) + minValue
}

func (stack *SpaceOptimizedSpecialStack) Push(value int) {
	if stack.wrapperStack.IsEmpty() || value < stack.minValue {
		stack.minValue = value
	}
	stack.wrapperStack.Push(getPushedValue(value, stack.minValue))
}

func (stack *SpaceOptimizedSpecialStack) Pop() (int, error) {
	pushedValue, error := stack.wrapperStack.Pop()
	if error != nil {
		return 0, error
	}
	if stack.wrapperStack.IsEmpty() {
		stack.minValue = -1
	} else {
		secondTopValue := stack.wrapperStack.Peak()
		stack.minValue = secondTopValue % dummyValue
	}
	return pushedValue / dummyValue, nil
}

func (stack *SpaceOptimizedSpecialStack) GetMin() int {
	if stack.wrapperStack.IsEmpty() {
		return -1
	}
	return stack.minValue
}

func CreateSpaceOptimizedSpecialStack() SpaceOptimizedSpecialStack {
	return SpaceOptimizedSpecialStack{wrapperStack: *CreateStack[int](nil)}
}
