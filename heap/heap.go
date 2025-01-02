// Copyright (c) 2024 Tecy.
// This file is licensed under the MIT License.
// See the LICENSE file in the project root for more information.

package heap

type Heap[T any] struct {
	value      []T
	comparator func(left T, right T) bool
}

// New creates an empty heap using the provided comparator.
func New[T any](comparator func(left T, right T) bool) *Heap[T] {
	return &Heap[T]{
		comparator: comparator,
	}
}

// NewWithData creates a heap using the provided comparator and data, with a time complexity of O(n).
// Slices can be passed by expanding them, assuming data is of type []T, it can be passed using data... .
func NewWithData[T any](comparator func(left T, right T) bool, data ...T) *Heap[T] {
	return MakeHeap(comparator, data...)
}

// Size returns the size of the heap.
func (heap *Heap[T]) Size() int {
	return len(heap.value)
}

// Empty returns true if the heap is empty; otherwise, it returns false.
func (heap *Heap[T]) Empty() bool {
	return len(heap.value) == 0
}

// leftChild returns the index of the left child of the parent.
func leftChild(parent int) int {
	return parent*2 + 1
}

// rightChild returns the index of the right child of the parent.
func rightChild(parent int) int {
	return parent*2 + 2
}

func (heap *Heap[T]) heapify(parent int) {
	largest := parent
	left := leftChild(parent)
	right := rightChild(parent)

	if left < len(heap.value) && heap.comparator(heap.value[largest], heap.value[left]) {
		largest = left
	}
	if right < len(heap.value) && heap.comparator(heap.value[largest], heap.value[right]) {
		largest = right
	}

	if largest != parent {
		heap.value[parent], heap.value[largest] = heap.value[largest], heap.value[parent]
		heap.heapify(largest)
	}
}

// MakeHeap builds a heap in O(n) time.
// Slices can be passed by expanding them, assuming data is of type []T, it can be passed using data... .
func MakeHeap[T any](comparator func(left T, right T) bool, data ...T) *Heap[T] {
	heap := &Heap[T]{
		value:      data,
		comparator: comparator,
	}

	for i := len(heap.value)/2 - 1; i >= 0; i-- {
		heap.heapify(i)
	}
	return heap
}

func (heap *Heap[T]) Push(value T) {
	// to do
}

func (heap *Heap[T]) Top() (value T) {
	if len(heap.value) > 0 {
		return heap.value[0]
	}
	return
}

func (heap *Heap[T]) Pop() (value T) {
	if len(heap.value) > 0 {
		// to do
	}
	return
}
