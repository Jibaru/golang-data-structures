package structures_test

import (
	"testing"

	"github.com/Jibaru/golang-data-structures/structures"
)

func TestNewStack(t *testing.T) {
	s := structures.NewStack[int]()

	if s.Size() != 0 {
		t.Fatalf("NewStack: Expected 0, got %v", s.Size())
	}
}

func TestStack_Push(t *testing.T) {
	s := structures.NewStack[int]()

	s.Push(1)
	s.Push(2)
	s.Push(3)

	if s.Size() != 3 {
		t.Fatalf("Stack.Push: Expected 3, got %v", s.Size())
	}
}

func TestStack_Pop(t *testing.T) {
	s := structures.NewStack[int]()

	s.Push(1)
	s.Push(2)
	s.Push(3)

	v, err := s.Pop()
	if err != nil {
		t.Fatalf("Stack.Pop: Expected nil, got %v", err)
	}

	if v != 3 {
		t.Fatalf("Stack.Pop: Expected 3, got %v", v)
	}

	if s.Size() != 2 {
		t.Fatalf("Stack.Pop: Expected 2, got %v", s.Size())
	}
}

func TestStack_Pop_Empty(t *testing.T) {
	s := structures.NewStack[int]()

	_, err := s.Pop()
	if err == nil {
		t.Fatalf("Stack.Pop: Expected ErrEmptyStack, got nil")
	}
}

func TestStack_Top(t *testing.T) {
	s := structures.NewStack[int]()

	s.Push(1)
	s.Push(2)
	s.Push(3)

	v, err := s.Top()
	if err != nil {
		t.Fatalf("Stack.Top: Expected nil, got %v", err)
	}

	if v != 3 {
		t.Fatalf("Stack.Top: Expected 3, got %v", v)
	}
}

func TestStack_Bottom(t *testing.T) {
	s := structures.NewStack[int]()

	s.Push(1)
	s.Push(2)
	s.Push(3)

	v, err := s.Bottom()
	if err != nil {
		t.Fatalf("Stack.Bottom: Expected nil, got %v", err)
	}

	if v != 1 {
		t.Fatalf("Stack.Bottom: Expected 1, got %v", v)
	}
}

func TestStack_Top_Empty(t *testing.T) {
	s := structures.NewStack[int]()

	_, err := s.Top()
	if err == nil {
		t.Fatalf("Stack.Top: Expected ErrEmptyStack, got nil")
	}
}

func TestStack_Bottom_Empty(t *testing.T) {
	s := structures.NewStack[int]()

	_, err := s.Bottom()
	if err == nil {
		t.Fatalf("Stack.Bottom: Expected ErrEmptyStack, got nil")
	}
}

func TestStack_Find(t *testing.T) {
	s := structures.NewStack[int]()

	s.Push(1)
	s.Push(2)
	s.Push(3)

	if !s.Find(2) {
		t.Fatalf("Stack.Find: Expected true, got false")
	}
}

func TestStack_Find_Empty(t *testing.T) {
	s := structures.NewStack[int]()

	if s.Find(2) {
		t.Fatalf("Stack.Find: Expected false, got true")
	}
}

func TestStack_Size(t *testing.T) {
	s := structures.NewStack[int]()

	s.Push(1)
	s.Push(2)
	s.Push(3)

	if s.Size() != 3 {
		t.Fatalf("Stack.Size: Expected 3, got %v", s.Size())
	}
}

func TestStack_Size_Empty(t *testing.T) {
	s := structures.NewStack[int]()

	if s.Size() != 0 {
		t.Fatalf("Stack.Size: Expected 0, got %v", s.Size())
	}
}
