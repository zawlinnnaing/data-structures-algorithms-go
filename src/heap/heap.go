package heap

import "errors"

type MinHeap struct {
	data   []int
	length int
}

func (heap *MinHeap) Insert(value int) {
	heap.data = append(heap.data, value)
	heap.heapifyUp(heap.length)
	heap.length += 1
}

func (heap *MinHeap) Delete() (int, error) {
	if heap.length == 0 {
		return 0, errors.New("empty heap")
	}
	out := heap.data[0]
	heap.length -= 1
	heap.data[0] = heap.data[heap.length]
	heap.heapifyDown(0)
	return out, nil
}

func (heap *MinHeap) Data() []int {
	return heap.data[:heap.length]
}

func (heap *MinHeap) heapifyDown(idx int) {
	leftChildIdx := heap.leftChildIdx(idx)
	rightChildIdx := heap.rightChildIdx(idx)
	if idx >= heap.length {
		return
	}

	// Check whether node at `idx` has children
	if leftChildIdx >= heap.length {
		return
	}

	value := heap.data[idx]

	leftChild := heap.data[leftChildIdx]
	rightChild := heap.data[rightChildIdx]

	if leftChild < rightChild && leftChild < value {
		heap.swap(idx, leftChildIdx)
		heap.heapifyDown(leftChildIdx)
		return
	}

	if rightChild < leftChild && rightChild < value {
		heap.swap(idx, rightChildIdx)
		heap.heapifyDown(rightChildIdx)
	}
}

func (heap *MinHeap) heapifyUp(idx int) {
	if idx == 0 {
		return
	}
	parentIndex := heap.parentIndex(idx)
	value := heap.data[idx]
	parent := heap.data[heap.parentIndex(idx)]
	if value > parent {
		return
	}
	heap.swap(idx, parentIndex)
	heap.heapifyUp(parentIndex)
}

func (heap *MinHeap) parentIndex(idx int) int {
	return (idx - 1) / 2
}

func (heap *MinHeap) leftChildIdx(idx int) int {
	return (2 * idx) + 1
}

func (heap *MinHeap) rightChildIdx(idx int) int {
	return (2 * idx) + 2
}

func (heap *MinHeap) swap(idx1 int, idx2 int) {
	temp := heap.data[idx1]
	heap.data[idx1] = heap.data[idx2]
	heap.data[idx2] = temp
}

func NewMinHeap() *MinHeap {
	return &MinHeap{data: []int{}}
}
