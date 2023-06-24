package queue_test

import (
	"testing"

	"github.com/zawlinnnaing/data-structures-algorithms-go/src/queue"
)

func TestQueuePush(t *testing.T) {
	exampleQueue := queue.NewQueue[string]("Hello", "World")
	exampleQueue.Push("Golang")
	if exampleQueue.Length() != 3 {
		t.Errorf("Expect Queue Length to be %v, Received %v", 3, exampleQueue.Length())
	}
	if exampleQueue.Peek() != "Hello" {
		t.Errorf("Expected %v, Received %v", "Hello", exampleQueue.Peek())
	}
}

func TestQueuePop(t *testing.T) {
	exampleQueue := queue.NewQueue[string]("hello", "world")
	result := exampleQueue.Pop()
	if result != "hello" {
		t.Errorf("Expected %v, Received %v", "hello", result)
	}
	if exampleQueue.Length() != 1 {
		t.Errorf("Expected 1, Received %v", exampleQueue.Length())
	}
}
