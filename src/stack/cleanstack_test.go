package stack_test

import (
	"testing"

	"github.com/zawlinnnaing/data-structures-algorithms-go/src/stack"
)

func TestStackPop(t *testing.T) {
	exampleStack := stack.NewStack[string]("hello", "world", "ni hao")
	exampleStack.Pop()
	result := exampleStack.Peek()
	if result != "world" {
		t.Errorf("Expected world, Received %v", result)
	}
	if exampleStack.Length() != 2 {
		t.Errorf("Expected length 2, Received %v", exampleStack.Length())
	}
}

func TestStackPush(t *testing.T) {
	exampleStack := stack.NewStack[string]()
	exampleStack.Push("Hello")
	exampleStack.Push("World")
	result := exampleStack.Peek()
	if result != "World" {
		t.Errorf("Expected %v, Received %v", "world", result)
	}
	if exampleStack.Length() != 2 {
		t.Errorf("Expected %v, Received %v", 2, exampleStack.Length())
	}
}
