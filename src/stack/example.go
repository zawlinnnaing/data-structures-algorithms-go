package stack

import "fmt"

func runSpecialStackExample() {
	specialStack := CreateSpecialStack()
	specialStack.Push(0)
	specialStack.Push(-1)
	specialStack.Push(3)
	specialStack.Push(-2)
	specialStack.Pop()
	specialStack.PrintStack("\t")
	minValue, error := specialStack.GetMinValue()
	if error == nil {
		fmt.Printf("Minimum value: %d \n", minValue) // Should print -1
	} else {
		fmt.Print(error.Error())
	}
}

func runSpaceOptimizedExample() {
	spaceOptimizedStack := CreateSpaceOptimizedSpecialStack()
	spaceOptimizedStack.Push(0)
	spaceOptimizedStack.Push(-1)
	spaceOptimizedStack.Push(-3)
	spaceOptimizedStack.Push(5)
	spaceOptimizedStack.Push(-2)
	spaceOptimizedStack.Pop()
	spaceOptimizedStack.Pop()
	spaceOptimizedStack.Push(-6)
	spaceOptimizedStack.Pop()
	fmt.Printf("Min value for space optimized stack: %d \n", spaceOptimizedStack.GetMin())
}

func runMonotonicStackExample(mode MonotonicStackMode) {
	var initialElements []int
	if mode == Increasing {
		initialElements = []int{1, 3, 4, 6}
	} else {
		initialElements = []int{6, 4, 3, 1}
	}
	monotonicStack := CreateMonotonicStack(mode, initialElements)
	monotonicStack.innerStack.PrintStack("")
	elementsToPush := []int{2, 4, 3}
	for i := 0; i < len(elementsToPush); i++ {
		fmt.Printf("Monotonic stack after pushing %d: \n", elementsToPush[i])
		monotonicStack.Push(elementsToPush[i])
		monotonicStack.innerStack.PrintStack("")
	}
}

func RunAllStackExamples() {
	fmt.Println("Running stack examples --------")
	fmt.Println("Running special stack examples")
	runSpecialStackExample()
	fmt.Println("Running space optimized stack examples")
	runSpaceOptimizedExample()
	fmt.Println("Running monotonic increasing stack example")
	runMonotonicStackExample(Increasing)
	fmt.Println("Running monotonic decreasing stack example")
	runMonotonicStackExample(Decreasing)
}
