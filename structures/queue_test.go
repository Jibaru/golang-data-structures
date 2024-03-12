package structures_test

import (
	"testing"

	"github.com/Jibaru/golang-data-structures/structures"
)

func TestNewQueue(t *testing.T) {
	q := structures.NewQueue[int]()

	if q.Size() != 0 {
		t.Fatalf("NewQueue: Expected 0, got %v", q.Size())
	}
}

func TestQueue_Push(t *testing.T) {
	q := structures.NewQueue[int]()

	q.Push(1)
	q.Push(2)
	q.Push(3)

	if q.Size() != 3 {
		t.Fatalf("Queue.Push: Expected 3, got %v", q.Size())
	}
}

func TestQueue_Pop(t *testing.T) {
	q := structures.NewQueue[int]()

	q.Push(1)
	q.Push(2)
	q.Push(3)

	v, err := q.Pop()
	if err != nil {
		t.Fatalf("Queue.Pop: Expected nil, got %v", err)
	}

	if v != 1 {
		t.Fatalf("Queue.Pop: Expected 1, got %v", v)
	}

	if q.Size() != 2 {
		t.Fatalf("Queue.Pop: Expected 2, got %v", q.Size())
	}
}

func TestQueue_Pop_Empty(t *testing.T) {
	q := structures.NewQueue[int]()

	_, err := q.Pop()
	if err == nil {
		t.Fatalf("Queue.Pop: Expected ErrEmptyQueue, got nil")
	}
}

func TestQueue_Top(t *testing.T) {
	q := structures.NewQueue[int]()

	q.Push(1)
	q.Push(2)
	q.Push(3)

	v, err := q.Top()
	if err != nil {
		t.Fatalf("Queue.Top: Expected nil, got %v", err)
	}

	if v != 1 {
		t.Fatalf("Queue.Top: Expected 1, got %v", v)
	}

	if q.Size() != 3 {
		t.Fatalf("Queue.Top: Expected 3, got %v", q.Size())
	}
}

func TestQueue_Top_Empty(t *testing.T) {
	q := structures.NewQueue[int]()

	_, err := q.Top()
	if err == nil {
		t.Fatalf("Queue.Top: Expected ErrEmptyQueue, got nil")
	}
}

func TestQueue_Bottom(t *testing.T) {
	q := structures.NewQueue[int]()

	q.Push(1)
	q.Push(2)
	q.Push(3)

	v, err := q.Bottom()
	if err != nil {
		t.Fatalf("Queue.Bottom: Expected nil, got %v", err)
	}

	if v != 3 {
		t.Fatalf("Queue.Bottom: Expected 3, got %v", v)
	}

	if q.Size() != 3 {
		t.Fatalf("Queue.Bottom: Expected 3, got %v", q.Size())
	}
}

func TestQueue_Bottom_Empty(t *testing.T) {
	q := structures.NewQueue[int]()

	_, err := q.Bottom()
	if err == nil {
		t.Fatalf("Queue.Bottom: Expected ErrEmptyQueue, got nil")
	}
}

func TestQueue_Find(t *testing.T) {
	q := structures.NewQueue[int]()

	q.Push(1)
	q.Push(2)
	q.Push(3)

	if !q.Find(2) {
		t.Fatalf("Queue.Find: Expected true, got false")
	}
}

func TestQueue_Find_Empty(t *testing.T) {
	q := structures.NewQueue[int]()

	if q.Find(2) {
		t.Fatalf("Queue.Find: Expected false, got true")
	}
}

func TestQueue_Find_NotFound(t *testing.T) {
	q := structures.NewQueue[int]()

	q.Push(1)
	q.Push(2)
	q.Push(3)

	if q.Find(4) {
		t.Fatalf("Queue.Find: Expected false, got true")
	}
}

func TestQueue_Size(t *testing.T) {
	q := structures.NewQueue[int]()

	q.Push(1)
	q.Push(2)
	q.Push(3)

	if q.Size() != 3 {
		t.Fatalf("Queue.Size: Expected 3, got %v", q.Size())
	}
}

func TestQueue_Size_Empty(t *testing.T) {
	q := structures.NewQueue[int]()

	if q.Size() != 0 {
		t.Fatalf("Queue.Size: Expected 0, got %v", q.Size())
	}
}
