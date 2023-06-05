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

func RunAllStackExamples() {
	fmt.Println("Running stack examples --------")
	fmt.Println("Running special stack examples")
	runSpecialStackExample()
}
