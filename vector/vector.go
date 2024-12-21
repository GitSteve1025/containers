// Copyright (c) 2024 Tecy.  All rights reserved.
// license that can be found in the LICENSE file.

package vector

import "errors"

type Vector[T any] []T

// New creates a new empty Vector[T].
// Also, vec := make(Vector[T], ...) is valid.
func New[T any]() *Vector[T] {
	var vec Vector[T]
	return &vec
}

// Size returns the number of elements in the vector.
func (vec *Vector[T]) Size() int {
	return len(*vec)
}

// Capacity returns the total number of elements that the vector can hold before needing to allocate more memory.
func (vec *Vector[T]) Capacity() int {
	return cap(*vec)
}

// Empty returns true if the vector is empty.
func (vec *Vector[T]) Empty() bool {
	return len(*vec) == 0
}

// Front returns the reference of data at the first element of the vector.
func (vec *Vector[T]) Front() *T {
	return &(*vec)[0]
}

// Back returns the reference of data at the last element of the vector.
func (vec *Vector[T]) Back() *T {
	return &(*vec)[len(*vec)-1]
}

// Resize resizes the vector to the specified number of elements.
func (vec *Vector[T]) Resize(n int) {
	temp := make(Vector[T], n)
	copy(temp, *vec)
	*vec = temp
}

// Assign assigns a given value to a vector.
func (vec *Vector[T]) Assign(n int, val T) {
	*vec = make(Vector[T], n)
	for i := range *vec {
		(*vec)[i] = val
	}
}

// PushBack adds data to the end of the vector.
func (vec *Vector[T]) PushBack(val T) {
	*vec = append(*vec, val)
}

// PopBack removes last element and returns the value of the element.
func (vec *Vector[T]) PopBack() T {
	defer func() { *vec = (*vec)[:len(*vec)-1] }()
	return (*vec)[len(*vec)-1]
}

// Insert inserts given value into vector before specified position.
func (vec *Vector[T]) Insert(pos int, val T) {
	*vec = append(*vec, val)
	copy((*vec)[pos+1:], (*vec)[pos:])
	(*vec)[pos] = val
}

// Erase removes element at given position and returns the value of the element.
func (vec *Vector[T]) Erase(pos int) T {
	defer func() {
		copy((*vec)[pos:], (*vec)[pos+1:])
		*vec = (*vec)[:len(*vec)-1]
	}()
	return (*vec)[pos]
}

// At provides for safer data access.
// The parameter is first checked that it is in the range of the vector.
// The function throws index out of range if the check fails.
func (vec *Vector[T]) At(pos int) (T, error) {
	if pos < 0 || pos >= len(*vec) {
		var val T
		return val, errors.New("index is out of range")
	}
	return (*vec)[pos], nil
}

// ShrinkToFit is to reduce Capacity() to Size().
// This function will create a new slice.
func (vec *Vector[T]) ShrinkToFit() {
	temp := make(Vector[T], len(*vec))
	copy(temp, *vec)
	*vec = temp
}
