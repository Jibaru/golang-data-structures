package structures

import "errors"

var (
	ErrEmptySimpleLinkedList             = errors.New("empty list")
	ErrIndexOutOfRangeInSimpleLinkedList = errors.New("index out of range")
)

// simpleLinkedList is a simple linked list that stores the first node and the size of the list
type simpleLinkedList[T comparable] struct {
	first *node[T]
	size  int64
}

// NewSimpleLinkedList creates a new simple linked list
func NewSimpleLinkedList[T comparable]() SimpleLinkedLister[T] {
	return &simpleLinkedList[T]{}
}

// Size returns the size of the list
func (s *simpleLinkedList[T]) Size() int64 {
	return s.size
}

// Push adds a new node with value val to the end of the list
func (s *simpleLinkedList[T]) Find(value T) bool {
	for current := s.first; current != nil; current = current.next {
		if current.value == value {
			return true
		}
	}

	return false
}

// PushAt adds a new node with value at index
func (s *simpleLinkedList[T]) PushAt(index int, value T) error {
	if index < 0 || int(s.Size()) < index {
		return ErrIndexOutOfRangeInSimpleLinkedList
	}

	if index == 0 {
		s.first = &node[T]{value, s.first}
		s.size++
		return nil
	}

	current := s.first
	for i := 0; i < index-1; i++ {
		current = current.next
	}

	current.next = &node[T]{value, current.next}
	s.size++

	return nil
}

// PopAt removes the node at index
func (s *simpleLinkedList[T]) PopAt(index int) (T, error) {
	if s.Size() == 0 {
		var t T
		return t, ErrEmptySimpleLinkedList
	}

	if index < 0 || int(s.Size()) <= index {
		var t T
		return t, ErrIndexOutOfRangeInSimpleLinkedList
	}

	if index == 0 {
		value := s.first.value
		s.first = s.first.next
		s.size--
		return value, nil
	}

	current := s.first
	for i := 0; i < index-1; i++ {
		current = current.next
	}

	value := current.next.value
	current.next = current.next.next
	s.size--

	return value, nil
}

// First returns the value of the first node
func (s *simpleLinkedList[T]) First() (T, error) {
	if s.size == 0 {
		var t T
		return t, ErrEmptySimpleLinkedList
	}

	return s.first.value, nil
}

// Last returns the value of the last node
func (s *simpleLinkedList[T]) Last() (T, error) {
	if s.size == 0 {
		var t T
		return t, ErrEmptySimpleLinkedList
	}

	current := s.first
	for current.next != nil {
		current = current.next
	}

	return current.value, nil
}

// GetAt returns the value of the node at index
func (s *simpleLinkedList[T]) GetAt(index int) (T, error) {
	if s.Size() == 0 {
		var t T
		return t, ErrEmptySimpleLinkedList
	}

	if index < 0 || int(s.Size()) <= index {
		var t T
		return t, ErrIndexOutOfRangeInSimpleLinkedList
	}

	current := s.first
	for i := 0; i < index; i++ {
		current = current.next
	}

	return current.value, nil
}
