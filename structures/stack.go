package structures

import "errors"

// ErrEmptyStack tolds the stack is empty
var (
	ErrEmptyStack = errors.New("empty stack")
)

// stack concrete implementation
type stack[T comparable] struct {
	size int64
	top  *node[T]
}

// NewStack retrieve an empty stack
func NewStack[T comparable]() Stacker[T] {
	return &stack[T]{}
}

// Pop retrieves the element at the top of the stack
// If stack is empty, retrieve an ErrEmptyStack
func (s *stack[T]) Pop() (T, error) {
	if s.size == 0 {
		var t T
		return t, ErrEmptyStack
	}

	topValue := s.top.value
	s.top = s.top.next
	s.size--

	return topValue, nil
}

// Push allows to insert an element at the top the stack
func (s *stack[T]) Push(value T) {
	s.top = &node[T]{value, s.top}
	s.size++
}

// Find check if an element exists on the stack
func (s *stack[T]) Find(value T) bool {
	if s.size == 0 {
		return false
	}

	node := s.top

	for node.value != value {
		node = node.next

		if node == nil {
			return false
		}
	}

	return true
}

// Top retrieves a copy of the element at the top of the stack
// If stack is empty, retrieve an ErrEmptyStack
func (s *stack[T]) Top() (T, error) {
	if s.size == 0 {
		var t T
		return t, ErrEmptyStack
	}

	return s.top.value, nil
}

// Bottom retrieves a copy of the element at the bottom of the stack
// If stack is empty, retrieve an ErrEmptyStack
func (s *stack[T]) Bottom() (T, error) {
	if s.size == 0 {
		var t T
		return t, ErrEmptyStack
	}

	node := s.top
	bottom := node

	for node != nil {
		bottom = node
		node = node.next
	}

	return bottom.value, nil
}

// Size retrieves the the quantity of elements on the stack
func (s *stack[T]) Size() int64 {
	return s.size
}
