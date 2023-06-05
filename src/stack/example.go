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

func RunAllStackExamples() {
	fmt.Println("Running stack examples --------")
	fmt.Println("Running special stack examples")
	runSpecialStackExample()
	fmt.Println("Running space optimized stack examples")
	runSpaceOptimizedExample()
}
