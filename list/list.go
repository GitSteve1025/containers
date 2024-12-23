// Copyright (c) 2024 Tecy.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package list

type Element[T any] struct {
	// The value stored in this element.
	Value T
	// previous pointer of the element.
	prev *Element[T]
	// next pointer of the element.
	next *Element[T]
	// The list to which this element belongs.
	list *List[T]
}

func (e *Element[T]) Prev() *Element[T] {
	if p := e.prev; e.list != nil && p != &e.list.root {
		return p
	}
	return nil
}

func (e *Element[T]) Next() *Element[T] {
	if n := e.next; e.list != nil && n != &e.list.root {
		return n
	}
	return nil
}

// doubly-linked list.
type List[T any] struct {
	// sentinel list element, only &root, root.prev, and root.next are used
	root Element[T]
	// length of list.
	size int
}

// New creates a new empty List[T].
func New[T any]() *List[T] {
	list := new(List[T])
	list.init()
	return list
}

// init initializes or clears list.
func (list *List[T]) init() {
	list.root.prev = &list.root
	list.root.next = &list.root
	list.size = 0
}

// Front returns the reference of data at the first element of the list.
func (list *List[T]) Front() *Element[T] {
	if list.size != 0 {
		return list.root.next
	}
	return nil
}

// Back returns the reference of data at the last element of the list.
func (list *List[T]) Back() *Element[T] {
	if list.size != 0 {
		return list.root.prev
	}
	return nil
}

// insert inserts v before at, increments list.size and returns the element.
func (list *List[T]) insert(v *Element[T], at *Element[T]) *Element[T] {
	v.prev = at.prev
	v.next = at
	v.list = list
	at.prev.next = v
	at.prev = v
	list.size++
	return v
}

// insertValue is a convenience wrapper for insert(&Element[T]{Value: val}, at) and returns the new element.
func (list *List[T]) insertValue(val T, at *Element[T]) *Element[T] {
	return list.insert(&Element[T]{Value: val}, at)
}

// PushBack adds data to the end of the list and returns the new element.
func (list *List[T]) PushBack(val T) *Element[T] {
	return list.insertValue(val, &list.root)
}

// PushFront adds data to the begin of the list and returns the new element.
func (list *List[T]) PushFront(val T) *Element[T] {
	return list.insertValue(val, list.root.next)
}

// InsertBefore inserts val before at, and return the new element.
// If at is not an element of list, the list is not modified.
// The at must not be nil.
func (list *List[T]) InsertBefore(val T, at *Element[T]) *Element[T] {
	if at.list == list {
		return list.insertValue(val, at)
	}
	return nil
}

// InsertAfter inserts val after at, and return the new element.
// If at is not an element of list, the list is not modified.
// The at must not be nil.
func (list *List[T]) InsertAfter(val T, at *Element[T]) *Element[T] {
	if at.list == list {
		return list.insertValue(val, at.next)
	}
	return nil
}

// erase erases at element.
func (list *List[T]) erase(at *Element[T]) {
	at.prev.next = at.next
	at.next.prev = at.prev
	at.prev = nil // avoid memory leaks
	at.next = nil // avoid memory leaks
	at.list = nil
	list.size--
}

// PopBack removes last element and returns the value of the element.
// It will return default value of T when list is empty.
func (list *List[T]) PopBack() T {
	if list.size > 0 {
		defer list.erase(list.root.prev)
		return list.root.prev.Value
	}
	var dft T
	return dft
}

// PopFront removes the first element and returns the value of the element.
// It will return default value of T when list is empty.
func (list *List[T]) PopFront() T {
	if list.size > 0 {
		defer list.erase(list.root.next)
		return list.root.next.Value
	}
	var dft T
	return dft
}

// Erase erases at element and returns the value.
// If at is not an element of list, the list is not modified.
// The at must not be nil.
func (list *List[T]) Erase(at *Element[T]) T {
	if at.list == list {
		defer list.erase(at)
		return at.Value
	}
	var dft T
	return dft
}
