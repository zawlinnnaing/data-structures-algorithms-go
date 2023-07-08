package heap_test

import (
	"testing"

	"github.com/zawlinnnaing/data-structures-algorithms-go/src/heap"
	"github.com/zawlinnnaing/data-structures-algorithms-go/src/util"
)

func testHeap(expected []int, received []int, t *testing.T) {
	if !util.SliceEquals[int](received, expected) {
		t.Errorf("Expected %v, received %v", expected, received)
	}
}

func TestMinHeap(t *testing.T) {
	minHeap := heap.NewMinHeap()
	expected := []int{}
	minHeap.Insert(2)
	expected = append(expected, 2)
	testHeap(expected, minHeap.Data(), t)
	minHeap.Insert(5)
	expected = append(expected, 5)
	testHeap(expected, minHeap.Data(), t)
	minHeap.Insert(7)
	minHeap.Insert(3)
	expected = []int{2, 3, 7, 5}
	testHeap(expected, minHeap.Data(), t)
	out, _ := minHeap.Delete()
	if out != 2 {
		t.Errorf("Expected %v, received %v", 2, out)
	}
	expected = []int{3, 5, 7}
	testHeap(expected, minHeap.Data(), t)
}

func TestMinHeapEmpty(t *testing.T) {
	minHeap := heap.NewMinHeap()
	_, deleteError := minHeap.Delete()
	if deleteError == nil {
		t.Errorf("Expected error, received %v", deleteError)
	}
	minHeap.Insert(2)
	out, _ := minHeap.Delete()
	if out != 2 {
		t.Errorf("Expected 2, received %v", 2)
	}
	testHeap([]int{}, minHeap.Data(), t)
}
