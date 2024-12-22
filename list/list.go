package list

type Element[T any] struct {
	Value T
	Prev  *Element[T]
	Next  *Element[T]
}

// Doubly linked list.
type List[T any] struct {
	head *Element[T]
	tail *Element[T]
}

// New creates a new empty List[T].
func New[T any]() *List[T] {
	var list List[T]
	return &list
}

// Front returns the reference of data at the first element of the list.
func (list *List[T]) Front() *Element[T] {
	return list.head
}

// Back returns the reference of data at the last element of the list.
func (list *List[T]) Back() *Element[T] {
	return list.tail
}

// PushBack adds data to the end of the list.
func (list *List[T]) PushBack(val T) {
	if list.tail != nil {
		list.tail.Next = &Element[T]{
			Value: val,
			Prev:  list.tail,
			Next:  nil,
		}
		list.tail = list.tail.Next
	} else {
		list.tail = &Element[T]{
			Value: val,
			Prev:  nil,
			Next:  nil,
		}
		list.head = list.tail
	}
}

// PopBack removes last element and returns the value of the element.
func (list *List[T]) PopBack() T {
	defer func() {
		list.tail = list.tail.Prev
		if list.tail != nil {
			list.tail.Next = nil
		} else {
			list.head = nil
		}
	}()
	return list.tail.Value
}

// PushFront adds data to the begin of the list.
func (list *List[T]) PushFront(val T) {
	if list.head != nil {
		list.head.Prev = &Element[T]{
			Value: val,
			Prev:  nil,
			Next:  list.head,
		}
		list.head = list.head.Prev
	} else {
		list.head = &Element[T]{
			Value: val,
			Prev:  nil,
			Next:  nil,
		}
		list.tail = list.head
	}
}

// PopFront removes the first element and returns the value of the element.
func (list *List[T]) PopFront() T {
	defer func() {
		list.head = list.head.Next
		if list.head != nil {
			list.head.Prev = nil
		} else {
			list.tail = nil
		}
	}()
	return list.head.Value
}
