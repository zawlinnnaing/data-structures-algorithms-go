package stack

type MonotonicStackMode int

const (
	Increasing MonotonicStackMode = iota
	Decreasing
)

type MonotonicStack struct {
	innerStack *Stack[int]
	Mode       MonotonicStackMode
}

func (stack *MonotonicStack) Push(element int) {
	for !stack.innerStack.IsEmpty() {
		elementFromStack := stack.innerStack.Peak()
		if stack.shouldPopElement(elementFromStack, element) {
			stack.innerStack.Pop()
		} else {

			break
		}
	}
	stack.innerStack.Push(element)
}

func (stack *MonotonicStack) shouldPopElement(elementFromStack int, incomingElement int) bool {
	if stack.Mode == Increasing {
		return elementFromStack >= incomingElement
	}
	return elementFromStack <= incomingElement
}

func CreateMonotonicStack(mode MonotonicStackMode) MonotonicStack {
	return MonotonicStack{
		innerStack: CreateStack[int](nil),
		Mode:       mode,
	}
}
