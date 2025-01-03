// Copyright (c) 2024 Tecy.
// This file is licensed under the MIT License.
// See the LICENSE file in the project root for more information.

// The minimum element in the tree is the root, at index 0.

package heap

type Heap[T any] struct {
	value      []T
	comparator func(left T, right T) bool
}

// New creates an empty heap using the provided comparator.
// Comparator will be used to build a min heap.
// Comparator must not be nil.
func New[T any](comparator func(left T, right T) bool) *Heap[T] {
	return &Heap[T]{
		comparator: comparator,
	}
}

// NewWithData creates a heap using the provided comparator and data, with a time complexity of O(n).
// Comparator will be used to build a min heap.
// Comparator must not be nil.
// Slices can be passed by expanding them, assuming data is of type []T, it can be passed using data... .
// If data is a slice, it may be modified.
func NewWithData[T any](comparator func(left T, right T) bool, data ...T) *Heap[T] {
	return makeHeap(comparator, data...)
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

// heapify is used to adjust a subtree to ensure it satisfies the heap property.
// // The complexity is O(log n)
func (heap *Heap[T]) heapify(parent int) {
	for {
		smallest := parent
		left := leftChild(parent)
		right := rightChild(parent)

		if left < len(heap.value) && heap.comparator(heap.value[left], heap.value[smallest]) {
			smallest = left
		}
		if right < len(heap.value) && heap.comparator(heap.value[right], heap.value[smallest]) {
			smallest = right
		}

		if smallest != parent {
			heap.value[parent], heap.value[smallest] = heap.value[smallest], heap.value[parent]
			parent = smallest
		} else {
			break
		}
	}
}

// Compare the newly inserted element with its parent.
// If the new element is smaller than the parent, swap their positions.
// Repeat this process until the new element's position satisfies the heap property or it becomes the root node.
// The complexity is O(log n)
func (heap *Heap[T]) upHeap(child int) {
	for {
		parent := (child - 1) / 2
		if parent == child || heap.comparator(heap.value[parent], heap.value[child]) {
			break
		}
		heap.value[parent], heap.value[child] = heap.value[child], heap.value[parent]
		child = parent
	}
}

// makeHeap builds a heap in O(n) time.
func makeHeap[T any](comparator func(left T, right T) bool, data ...T) *Heap[T] {
	heap := &Heap[T]{
		value:      data,
		comparator: comparator,
	}

	for i := len(heap.value)/2 - 1; i >= 0; i-- {
		heap.heapify(i)
	}
	return heap
}

// Push inserts value into the heap with a time complexity of O(log n).
func (heap *Heap[T]) Push(value T) {
	heap.value = append(heap.value, value)
	heap.upHeap(len(heap.value) - 1)
}

// Top returns the top element of the heap with a time complexity of O(1).
// If the heap is empty, Top will return the default value of T.
func (heap *Heap[T]) Top() (value T) {
	if len(heap.value) > 0 {
		return heap.value[0]
	}
	return
}

// Pop removes the top element of the heap with a time complexity of O(log n).
// If the heap is empty, Pop will return the default value of T.
func (heap *Heap[T]) Pop() (value T) {
	if len(heap.value) > 0 {
		n := len(heap.value) - 1
		temp := heap.value[0]
		heap.value[0] = heap.value[n]
		heap.value = heap.value[:n]
		heap.heapify(0)
		return temp
	}
	return
}
