package queue

import "fmt"

type Node[T any] struct {
	next  *Node[T]
	value T
}

type Queue[T any] struct {
	head   *Node[T]
	tail   *Node[T]
	length int
}

func (queue *Queue[T]) Push(value T) {
	node := &Node[T]{value: value}
	queue.length += 1
	// In case queue is empty
	if queue.tail == nil {
		queue.head, queue.tail = node, node
		return
	}
	queue.tail.next = node
	queue.tail = node

}

func (queue *Queue[T]) Pop() (value T) {
	if queue.head == nil {
		return value
	}
	queue.length -= 1
	firstNode := queue.head
	queue.head = queue.head.next
	firstNode.next = nil
	value = firstNode.value
	return
}

func (queue *Queue[T]) Peek() (value T) {
	if queue.head == nil {
		return value
	}
	return queue.head.value
}

func (queue Queue[T]) String() string {
	if queue.length == 0 {
		return "Queue is empty"
	}
	temp := queue.head
	str := ""
	for i := 0; i < queue.length; i++ {
		str += fmt.Sprintf("[%v]", temp.value)
		temp = temp.next
	}
	return str
}

func (queue Queue[T]) Length() int {
	return queue.length
}

func NewQueue[T any](initialValues ...T) *Queue[T] {
	queue := &Queue[T]{
		length: 0,
	}
	for _, value := range initialValues {
		queue.Push(value)
	}
	return queue
}

func RunQueueExample() {
	exampleQueue := NewQueue[int]()
	itemsToPush := []int{1, 3, 4, 5}
	for _, value := range itemsToPush {
		exampleQueue.Push(value)
	}
	fmt.Println("Queue after pushing", exampleQueue.String(), exampleQueue.length)
	queueLength := exampleQueue.Length()
	for i := 0; i < queueLength; i++ {
		fmt.Println("Popping queue:", exampleQueue.Pop())
	}
	fmt.Println("After popping queue:", exampleQueue.String())
}
