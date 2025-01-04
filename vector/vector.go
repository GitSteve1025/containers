// Copyright (c) 2024 Tecy.
// This file is licensed under the MIT License.
// See the LICENSE file in the project root for more information.

package vector

type Vector[T any] []T

// New creates a new empty Vector[T].
// Also, vec := make(Vector[T], ...) is valid.
func New[T any]() *Vector[T] {
	return new(Vector[T])
}

// NewWithData creates a Vector[T] with data
// Data will be placed in order.
func NewWithData[T any](data ...T) *Vector[T] {
	vec := make(Vector[T], len(data))
	copy(vec, data)
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
// Front will return nil if it is empty.
func (vec *Vector[T]) Front() *T {
	if len(*vec) > 0 {
		return &(*vec)[0]
	}
	return nil
}

// Back returns the reference of data at the last element of the vector.
// Front will return nil if it is empty.
func (vec *Vector[T]) Back() *T {
	if len(*vec) > 0 {
		return &(*vec)[len(*vec)-1]
	}
	return nil
}

// Resize resizes the vector to the specified number of elements.
// Resize will allocate new space.
func (vec *Vector[T]) Resize(n int) {
	temp := make(Vector[T], n)
	copy(temp, *vec)
	*vec = temp
}

// Assign assigns a given value to a vector.
// Assign will allocate new space.
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
// When vec is empty, vec will not be modified.
// PopBack returns the default value of T when vec is empty.
func (vec *Vector[T]) PopBack() (value T) {
	if len(*vec) > 0 {
		temp := (*vec)[len(*vec)-1]
		*vec = (*vec)[:len(*vec)-1]
		return temp
	}
	return
}

// Insert inserts given value into vector before specified position.
// If pos < 0 or pos > len(*vec), vec will not be modified.
func (vec *Vector[T]) Insert(pos int, val T) {
	if 0 <= pos && pos <= len(*vec) {
		*vec = append(*vec, val)
		copy((*vec)[pos+1:], (*vec)[pos:])
		(*vec)[pos] = val
	}
}

// Erase removes element at given position and returns the value of the element.
// When pos is out of range, vec will not be modified and erase will return the default value of T.
func (vec *Vector[T]) Erase(pos int) (value T) {
	if 0 <= pos && pos < len(*vec) {
		temp := (*vec)[pos]
		copy((*vec)[pos:], (*vec)[pos+1:])
		*vec = (*vec)[:len(*vec)-1]
		return temp
	}
	return
}

// At returns a reference to the data at position pos.
// If pos is out of range, At will return nil.
func (vec *Vector[T]) At(pos int) *T {
	if 0 <= pos && pos < len(*vec) {
		return &(*vec)[pos]
	}
	return nil
}

// ShrinkToFit is to reduce Capacity() to Size().
// This function will create a new slice.
func (vec *Vector[T]) ShrinkToFit() {
	temp := make(Vector[T], len(*vec))
	copy(temp, *vec)
	*vec = temp
}

// Clear clears Vector[T]
func (vec *Vector[T]) Clear() {
	*vec = (*vec)[:0]
}
