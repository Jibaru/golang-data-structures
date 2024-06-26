package structures

import "errors"

var (
	ErrEmptyDoubleLinkedList             = errors.New("empty list")
	ErrIndexOutOfRangeInDoubleLinkedList = errors.New("index out of range")
)

// doubleLinkedList is a double linked list that stores the first node and the size of the list
type doubleLinkedList[T comparable] struct {
	first *doubleNode[T]
	size  int64
}

// NewDoubleLinkedList creates a new double linked list
func NewDoubleLinkedList[T comparable]() DoubleLinkedLister[T] {
	return &doubleLinkedList[T]{}
}

// Size returns the size of the list
func (s *doubleLinkedList[T]) Size() int64 {
	return s.size
}

// Find returns true if the value exists in the list
func (s *doubleLinkedList[T]) Find(value T) bool {
	for current := s.first; current != nil; current = current.next {
		if current.value == value {
			return true
		}
	}

	return false
}

// PushAt adds a new node with value at index
func (s *doubleLinkedList[T]) PushAt(index int, value T) error {
	if index < 0 || int(s.Size()) < index {
		return ErrIndexOutOfRangeInDoubleLinkedList
	}

	newNode := &doubleNode[T]{value: value}

	if s.first == nil || index == 0 {
		newNode.next = s.first
		if s.first != nil {
			s.first.prev = newNode
		}
		s.first = newNode
		s.size++
		return nil
	}

	current := s.first
	for i := 0; i < index-1; i++ {
		current = current.next
	}

	newNode.next = current.next
	newNode.prev = current
	if current.next != nil {
		current.next.prev = newNode
	}
	current.next = newNode
	s.size++
	return nil
}

// PopAt removes the node at index
func (s *doubleLinkedList[T]) PopAt(index int) (T, error) {
	if s.Size() == 0 {
		var t T
		return t, ErrEmptyDoubleLinkedList
	}

	if index < 0 || int(s.Size()) <= index {
		var t T
		return t, ErrIndexOutOfRangeInDoubleLinkedList
	}

	var value T

	if index == 0 {
		value = s.first.value
		s.first = s.first.next
		if s.first != nil {
			s.first.prev = nil
		}
		s.size--
		return value, nil
	}

	current := s.first
	for i := 0; i < index-1; i++ {
		current = current.next
	}

	value = current.next.value

	if current.next.next != nil {
		current.next.next.prev = current
	}
	current.next = current.next.next
	s.size--
	return value, nil
}

// First returns the value of the first node
func (s *doubleLinkedList[T]) First() (T, error) {
	if s.size == 0 {
		var t T
		return t, ErrEmptyDoubleLinkedList
	}

	return s.first.value, nil
}

// Last returns the value of the last node
func (s *doubleLinkedList[T]) Last() (T, error) {
	if s.size == 0 {
		var t T
		return t, ErrEmptyDoubleLinkedList
	}

	current := s.first
	for current.next != nil {
		current = current.next
	}

	return current.value, nil
}

// GetAt returns the value of the node at index
func (s *doubleLinkedList[T]) GetAt(index int) (T, error) {
	if s.Size() == 0 {
		var t T
		return t, ErrEmptyDoubleLinkedList
	}

	if index < 0 || int(s.Size()) <= index {
		var t T
		return t, ErrIndexOutOfRangeInDoubleLinkedList
	}

	current := s.first
	for i := 0; i < index; i++ {
		current = current.next
	}

	return current.value, nil
}
