package array

import "fmt"

type RingBuffer[T any] struct {
	head   int
	tail   int
	Length int
	data   []T
}

func (ringBuffer *RingBuffer[T]) Shift(item T) {
	if ringBuffer.shouldGrow() {
		ringBuffer.Grow()
	}

	if ringBuffer.head == 0 {
		ringBuffer.head = len(ringBuffer.data) - 1
	} else {
		ringBuffer.head = ringBuffer.head - 1
	}
	ringBuffer.data[ringBuffer.head] = item
	ringBuffer.Length += 1
}

func (ringBuf *RingBuffer[T]) shouldGrow() bool {
	if len(ringBuf.data) == 0 {
		return true
	}
	return ringBuf.Length+1 > len(ringBuf.data)
}

func (ringBuf *RingBuffer[T]) UnShift() (value T) {

	if ringBuf.IsEmpty() {
		return
	}
	var zeroValue T
	value = ringBuf.data[ringBuf.head]
	ringBuf.data[ringBuf.head] = zeroValue
	ringBuf.Length -= 1
	ringBuf.head = (ringBuf.head + 1) % len(ringBuf.data)
	return
}

func (ringBuf *RingBuffer[T]) IsEmpty() bool {
	return ringBuf.Length == 0
}

func (ringBuf *RingBuffer[T]) Push(value T) {
	if ringBuf.shouldGrow() {
		ringBuf.Grow()
	}
	if ringBuf.IsEmpty() {
		ringBuf.head = ringBuf.tail
		ringBuf.data[ringBuf.tail] = value
	} else {
		ringBuf.tail = (ringBuf.tail + 1) % len(ringBuf.data)
		ringBuf.data[ringBuf.tail] = value
	}
	ringBuf.Length += 1
}

func (ringBuf *RingBuffer[T]) Pop() (value T) {
	if ringBuf.IsEmpty() {
		return value
	}
	var zeroValue T
	value = ringBuf.data[ringBuf.tail]
	ringBuf.data[ringBuf.tail] = zeroValue
	if ringBuf.tail == 0 {
		ringBuf.tail = len(ringBuf.data) - 1
	} else {
		ringBuf.tail -= 1
	}
	ringBuf.Length -= 1
	return
}

func (ringBuf *RingBuffer[T]) Grow() {
	oldDataLength := len(ringBuf.data)
	newArray := make([]T, len(ringBuf.data)+(len(ringBuf.data)/2)+1)
	for i := 0; i < oldDataLength; i++ {
		newArray[i] = ringBuf.Get(i)
	}
	ringBuf.data = newArray
	ringBuf.head = 0
	ringBuf.tail = oldDataLength - 1
}

func (ringBuf *RingBuffer[T]) Get(index int) (value T) {
	accessIndex := (ringBuf.head + index) % len(ringBuf.data)
	return ringBuf.data[accessIndex]
}

func (ringBuf *RingBuffer[T]) String() string {
	str := ""

	for i := 0; i < ringBuf.Length; i++ {
		str += fmt.Sprintf("%v ", ringBuf.Get(i))
	}
	return str
}

func NewRingBuffer[T any](items ...T) *RingBuffer[T] {
	tail := len(items) - 1
	if tail < 0 {
		tail = 0
	}

	return &RingBuffer[T]{
		head:   0,
		tail:   tail,
		data:   items,
		Length: len(items),
	}
}

func RunRingBufferExample() {
	ringBuf := NewRingBuffer[int]()

	examples := []int{5, 6, 7, 8}
	fmt.Println("Shifting ringBuf ==>")
	for _, value := range examples {
		ringBuf.Shift(value)
		fmt.Printf("After shifting %v: %v\n", value, ringBuf)
	}
	fmt.Println("After shifting length:", ringBuf.Length)
	fmt.Println("Unshifting ringBuf ==>")
	for ringBuf.Length > 0 {
		value := ringBuf.UnShift()
		fmt.Printf("After unshifting %v: %v\n", value, ringBuf)
	}
	fmt.Println("After unshifting length:", ringBuf.Length)
	fmt.Println("Pushing ring buffer ==>")
	for _, value := range examples {
		ringBuf.Push(value)
		fmt.Printf("After pushing %v: %v\n", value, ringBuf)
	}

	fmt.Println("Popping ring buffer ==>")
	for ringBuf.Length > 0 {
		popped := ringBuf.Pop()
		fmt.Printf("After popping %v: %v\n", popped, ringBuf)
	}

}
