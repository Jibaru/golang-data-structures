package structures

import (
	"errors"
)

// ErrEmptyQueue tolds the queue is empty
var (
	ErrEmptyQueue = errors.New("empty queue")
)

// stack concrete implementation
type queue[T comparable] struct {
	size   int64
	top    *node[T]
	bottom *node[T]
}

// NewQueue retrieve an empty queue
func NewQueue[T comparable]() Queuer[T] {
	return &queue[T]{}
}

// Pop retrieves the element at the top of the queue
// If queue is empty, retrieve an ErrEmptyQueue
func (s *queue[T]) Pop() (T, error) {
	if s.size == 0 {
		var t T
		return t, ErrEmptyQueue
	}

	topValue := s.top.value
	s.top = s.top.next

	if s.size == 1 {
		s.bottom = nil
	}

	s.size--

	return topValue, nil
}

// Push allows to insert an element at the bottom of the queue
func (s *queue[T]) Push(value T) {
	if s.size == 0 {
		s.bottom = &node[T]{value, nil}
		s.top = s.bottom
	} else {
		node := &node[T]{value, nil}
		s.bottom.next = node
		s.bottom = node
	}

	s.size++
}

// Find check if an element exists on the queue
func (s *queue[T]) Find(value T) bool {
	if s.size == 0 {
		return false
	}

	node := s.top

	for node != nil {
		if node.value == value {
			return true
		}

		node = node.next
	}

	return false
}

// Top retrieves a copy of the element at the top of the queue
// If queue is empty, retrieve an ErrEmptyQueue
func (s *queue[T]) Top() (T, error) {
	if s.size == 0 {
		var t T
		return t, ErrEmptyQueue
	}

	return s.top.value, nil
}

// Bottom retrieves a copy of the element at the bottom of the queue
// If queue is empty, retrieve an ErrEmptyQueue
func (s *queue[T]) Bottom() (T, error) {
	if s.size == 0 {
		var t T
		return t, ErrEmptyQueue
	}

	return s.bottom.value, nil
}

// Size retrieves the the quantity of elements on the queue
func (s *queue[T]) Size() int64 {
	return s.size
}
