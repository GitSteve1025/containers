package stack

type Stack[T any] struct {
	value []T
}

// New creates an empty stack.
func New[T any]() *Stack[T] {
	return new(Stack[T])
}

// Size returns the size of the stack.
func (stack *Stack[T]) Size() int {
	return len(stack.value)
}

// Empty returns true if the stack is empty; otherwise, it returns false.
func (stack *Stack[T]) Empty() bool {
	return len(stack.value) == 0
}

// Push puts value into the stack.
func (stack *Stack[T]) Push(value T) {
	stack.value = append(stack.value, value)
}

// Pop removes and returns the top of the stack, or the default value of T if the stack is empty.
func (stack *Stack[T]) Pop() (value T) {
	if len(stack.value) > 0 {
		temp := stack.value[len(stack.value)-1]
		stack.value = (stack.value)[:len(stack.value)-1]
		return temp
	}
	return
}

// Top returns a reference to the top element of the stack, or nil if the stack is empty.
func (stack *Stack[T]) Top() *T {
	if len(stack.value) > 0 {
		return &stack.value[len(stack.value)-1]
	}
	return nil
}
