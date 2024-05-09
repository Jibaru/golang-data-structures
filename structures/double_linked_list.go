package structures

import "errors"

var (
	ErrEmptyDoubleLinkedList             = errors.New("empty list")
	ErrIndexOutOfRangeInDoubleLinkedList = errors.New("index out of range")
)

type doubleLinkedList[T comparable] struct {
	first *doubleNode[T]
	size  int64
}

func NewDoubleLinkedList[T comparable]() DoubleLinkedLister[T] {
	return &doubleLinkedList[T]{}
}

func (s *doubleLinkedList[T]) Size() int64 {
	return s.size
}

func (s *doubleLinkedList[T]) Find(value T) bool {
	for current := s.first; current != nil; current = current.next {
		if current.value == value {
			return true
		}
	}

	return false
}

func (s *doubleLinkedList[T]) PushAt(index int, value T) error {
	if index < 0 || int(s.Size()) < index {
		return ErrIndexOutOfRangeInDoubleLinkedList
	}

	if index == 0 {
		s.first = &doubleNode[T]{
			value: value,
			prev:  nil,
			next:  s.first,
		}
		s.size++
		return nil
	}

	current := s.first
	for i := 0; i < index-1; i++ {
		current = current.next
	}

	current.next = &doubleNode[T]{
		value: value,
		prev:  current,
		next:  current.next,
	}
	s.size++

	return nil
}

func (s *doubleLinkedList[T]) PopAt(index int) (T, error) {
	if s.Size() == 0 {
		var t T
		return t, ErrEmptyDoubleLinkedList
	}

	if index < 0 || int(s.Size()) <= index {
		var t T
		return t, ErrIndexOutOfRangeInDoubleLinkedList
	}

	if index == 0 {
		value := s.first.value
		s.first = s.first.next
		s.first.prev = nil
		s.size--
		return value, nil
	}

	current := s.first
	for i := 0; i < index-1; i++ {
		current = current.next
	}

	value := current.next.value
	current.next = current.next.next
	current.next.prev = current
	s.size--

	return value, nil
}

func (s *doubleLinkedList[T]) First() (T, error) {
	if s.size == 0 {
		var t T
		return t, ErrEmptyDoubleLinkedList
	}

	return s.first.value, nil
}

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
