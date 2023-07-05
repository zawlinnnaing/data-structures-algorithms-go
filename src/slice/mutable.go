package slice

import "fmt"

type Mutable[T any] struct {
	elements []T
}

func (this *Mutable[T]) Push(value T) {
	this.elements = append(this.elements, value)
}

func (this *Mutable[T]) String() string {
	return fmt.Sprintf("%v", this.elements)
}

func NewMutable[T any](values ...T) Mutable[T] {
	return Mutable[T]{elements: values}
}
