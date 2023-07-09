package slice

import "fmt"

type Mutable[T any] struct {
	elements []T
}

func (this *Mutable[T]) Push(value T) {
	this.elements = append(this.elements, value)
}

func (this *Mutable[T]) Reverse() {
	reversed := []T{}
	for idx := range this.elements {
		reversed = append(reversed, this.elements[len(this.elements)-1-idx])
	}
	this.elements = reversed
}

func (this *Mutable[T]) Elements() []T {
	return this.elements
}

func (this *Mutable[T]) String() string {
	return fmt.Sprintf("%v", this.elements)
}

func NewMutable[T any](values ...T) Mutable[T] {
	return Mutable[T]{elements: values}
}
