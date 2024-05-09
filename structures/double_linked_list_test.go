package structures_test

import (
	"testing"

	"github.com/Jibaru/golang-data-structures/structures"
)

func TestNewDoubleLinkedList(t *testing.T) {
	l := structures.NewDoubleLinkedList[int]()

	if l.Size() != 0 {
		t.Fatalf("NewDoubleLinkedList: Expected 0, got %v", l.Size())
	}
}

func TestDoubleLinkedList_PushAt(t *testing.T) {
	l := structures.NewDoubleLinkedList[int]()

	err := l.PushAt(0, 1)
	if err != nil {
		t.Fatalf("doubleLinkedList.PushAt: Expected nil, got %v", err)
	}

	err = l.PushAt(1, 2)
	if err != nil {
		t.Fatalf("doubleLinkedList.PushAt: Expected nil, got %v", err)
	}

	err = l.PushAt(2, 3)
	if err != nil {
		t.Fatalf("doubleLinkedList.PushAt: Expected nil, got %v", err)
	}

	if l.Size() != 3 {
		t.Fatalf("doubleLinkedList.PushAt: Expected 3, got %v", l.Size())
	}

	got, err := l.GetAt(0)
	if err != nil {
		t.Fatalf("doubleLinkedList.GetAt: Expected nil, got %v", err)
	}

	if got != 1 {
		t.Fatalf("doubleLinkedList.GetAt: Expected 1, got %v", got)
	}

	got, err = l.GetAt(1)
	if err != nil {
		t.Fatalf("doubleLinkedList.GetAt: Expected nil, got %v", err)
	}

	if got != 2 {
		t.Fatalf("doubleLinkedList.GetAt: Expected 2, got %v", got)
	}

	got, err = l.GetAt(2)
	if err != nil {
		t.Fatalf("doubleLinkedList.GetAt: Expected nil, got %v", err)
	}

	if got != 3 {
		t.Fatalf("doubleLinkedList.GetAt: Expected 3, got %v", got)
	}
}

func TestDoubleLinkedList_PopAt(t *testing.T) {
	l := structures.NewDoubleLinkedList[int]()

	err := l.PushAt(0, 1)
	if err != nil {
		t.Fatalf("doubleLinkedList.PushAt: Expected nil, got %v", err)
	}

	err = l.PushAt(1, 2)
	if err != nil {
		t.Fatalf("doubleLinkedList.PushAt: Expected nil, got %v", err)
	}

	err = l.PushAt(2, 3)
	if err != nil {
		t.Fatalf("doubleLinkedList.PushAt: Expected nil, got %v", err)
	}

	v, err := l.PopAt(0)
	if err != nil {
		t.Fatalf("doubleLinkedList.PopAt: Expected nil, got %v", err)
	}

	if v != 1 {
		t.Fatalf("doubleLinkedList.PopAt: Expected 1, got %v", v)
	}

	if l.Size() != 2 {
		t.Fatalf("doubleLinkedList.PopAt: Expected 2, got %v", l.Size())
	}

	v, err = l.PopAt(1)
	if err != nil {
		t.Fatalf("doubleLinkedList.PopAt: Expected nil, got %v", err)
	}

	if v != 3 {
		t.Fatalf("doubleLinkedList.PopAt: Expected 3, got %v", v)
	}

	if l.Size() != 1 {
		t.Fatalf("doubleLinkedList.PopAt: Expected 1, got %v", l.Size())
	}
}

func TestDoubleLinkedList_PopAt_Empty(t *testing.T) {
	l := structures.NewDoubleLinkedList[int]()

	_, err := l.PopAt(0)
	if err == nil {
		t.Fatalf("doubleLinkedList.PopAt: Expected ErrEmptyDoubleLinkedList, got nil")
	}
}

func TestDoubleLinkedList_PopAt_OutOfRange(t *testing.T) {
	l := structures.NewDoubleLinkedList[int]()

	err := l.PushAt(0, 1)
	if err != nil {
		t.Fatalf("doubleLinkedList.PushAt: Expected nil, got %v", err)
	}

	_, err = l.PopAt(1)
	if err == nil {
		t.Fatalf("doubleLinkedList.PopAt: Expected ErrIndexOutOfRangeInDoubleLinkedList, got nil")
	}
}

func TestDoubleLinkedList_First(t *testing.T) {
	l := structures.NewDoubleLinkedList[int]()

	err := l.PushAt(0, 1)
	if err != nil {
		t.Fatalf("doubleLinkedList.PushAt: Expected nil, got %v", err)
	}

	v, err := l.First()
	if err != nil {
		t.Fatalf("doubleLinkedList.First: Expected nil, got %v", err)
	}

	if v != 1 {
		t.Fatalf("doubleLinkedList.First: Expected 1, got %v", v)
	}
}

func TestDoubleLinkedList_First_Empty(t *testing.T) {
	l := structures.NewDoubleLinkedList[int]()

	_, err := l.First()
	if err == nil {
		t.Fatalf("doubleLinkedList.First: Expected ErrEmptyDoubleLinkedList, got nil")
	}
}

func TestDoubleLinkedList_Last(t *testing.T) {
	l := structures.NewDoubleLinkedList[int]()

	err := l.PushAt(0, 1)
	if err != nil {
		t.Fatalf("doubleLinkedList.PushAt: Expected nil, got %v", err)
	}

	v, err := l.Last()
	if err != nil {
		t.Fatalf("doubleLinkedList.Last: Expected nil, got %v", err)
	}

	if v != 1 {
		t.Fatalf("doubleLinkedList.Last: Expected 1, got %v", v)
	}
}

func TestDoubleLinkedList_Last_Empty(t *testing.T) {
	l := structures.NewDoubleLinkedList[int]()

	_, err := l.Last()
	if err == nil {
		t.Fatalf("doubleLinkedList.Last: Expected ErrEmptyDoubleLinkedList, got nil")
	}
}

func TestDoubleLinkedList_Find(t *testing.T) {
	l := structures.NewDoubleLinkedList[int]()

	err := l.PushAt(0, 1)
	if err != nil {
		t.Fatalf("doubleLinkedList.PushAt: Expected nil, got %v", err)
	}

	err = l.PushAt(1, 2)
	if err != nil {
		t.Fatalf("doubleLinkedList.PushAt: Expected nil, got %v", err)
	}

	err = l.PushAt(2, 3)
	if err != nil {
		t.Fatalf("doubleLinkedList.PushAt: Expected nil, got %v", err)
	}

	if !l.Find(1) {
		t.Fatalf("doubleLinkedList.Find: Expected true, got false")
	}

	if !l.Find(2) {
		t.Fatalf("doubleLinkedList.Find: Expected true, got false")
	}

	if !l.Find(3) {
		t.Fatalf("doubleLinkedList.Find: Expected true, got false")
	}

	if l.Find(4) {
		t.Fatalf("doubleLinkedList.Find: Expected false, got true")
	}
}

func TestDoubleLinkedList_Find_Empty(t *testing.T) {
	l := structures.NewDoubleLinkedList[int]()

	if l.Find(1) {
		t.Fatalf("doubleLinkedList.Find: Expected false, got true")
	}
}

func TestDoubleLinkedList_Size(t *testing.T) {
	l := structures.NewDoubleLinkedList[int]()

	err := l.PushAt(0, 1)
	if err != nil {
		t.Fatalf("doubleLinkedList.PushAt: Expected nil, got %v", err)
	}

	err = l.PushAt(1, 2)
	if err != nil {
		t.Fatalf("doubleLinkedList.PushAt: Expected nil, got %v", err)
	}

	err = l.PushAt(2, 3)
	if err != nil {
		t.Fatalf("doubleLinkedList.PushAt: Expected nil, got %v", err)
	}

	if l.Size() != 3 {
		t.Fatalf("doubleLinkedList.Size: Expected 3, got %v", l.Size())
	}
}

func TestDoubleLinkedList_Size_Empty(t *testing.T) {
	l := structures.NewDoubleLinkedList[int]()

	if l.Size() != 0 {
		t.Fatalf("doubleLinkedList.Size: Expected 0, got %v", l.Size())
	}
}

func TestDoubleLinkedList_GetAt(t *testing.T) {
	l := structures.NewDoubleLinkedList[int]()

	err := l.PushAt(0, 1)
	if err != nil {
		t.Fatalf("doubleLinkedList.PushAt: Expected nil, got %v", err)
	}

	err = l.PushAt(1, 2)
	if err != nil {
		t.Fatalf("doubleLinkedList.PushAt: Expected nil, got %v", err)
	}

	err = l.PushAt(2, 3)
	if err != nil {
		t.Fatalf("doubleLinkedList.PushAt: Expected nil, got %v", err)
	}

	v, err := l.GetAt(0)
	if err != nil {
		t.Fatalf("doubleLinkedList.GetAt: Expected nil, got %v", err)
	}

	if v != 1 {
		t.Fatalf("doubleLinkedList.GetAt: Expected 1, got %v", v)
	}

	v, err = l.GetAt(1)
	if err != nil {
		t.Fatalf("doubleLinkedList.GetAt: Expected nil, got %v", err)
	}

	if v != 2 {
		t.Fatalf("doubleLinkedList.GetAt: Expected 2, got %v", v)
	}

	v, err = l.GetAt(2)
	if err != nil {
		t.Fatalf("doubleLinkedList.GetAt: Expected nil, got %v", err)
	}

	if v != 3 {
		t.Fatalf("doubleLinkedList.GetAt: Expected 3, got %v", v)
	}
}

func TestDoubleLinkedList_GetAt_Empty(t *testing.T) {
	l := structures.NewDoubleLinkedList[int]()

	_, err := l.GetAt(0)
	if err == nil {
		t.Fatalf("doubleLinkedList.GetAt: Expected ErrEmptyDoubleLinkedList, got nil")
	}
}

func TestDoubleLinkedList_GetAt_OutOfRange(t *testing.T) {
	l := structures.NewDoubleLinkedList[int]()

	err := l.PushAt(0, 1)
	if err != nil {
		t.Fatalf("doubleLinkedList.PushAt: Expected nil, got %v", err)
	}

	_, err = l.GetAt(1)
	if err == nil {
		t.Fatalf("doubleLinkedList.GetAt: Expected ErrIndexOutOfRangeInDoubleLinkedList, got nil")
	}
}

func TestDoubleLinkedList_PushAt_OutOfRange(t *testing.T) {
	l := structures.NewDoubleLinkedList[int]()

	err := l.PushAt(1, 1)
	if err == nil {
		t.Fatalf("doubleLinkedList.PushAt: Expected ErrIndexOutOfRangeInDoubleLinkedList, got nil")
	}
}

func TestDoubleLinkedList_PushAt_NegativeIndex(t *testing.T) {
	l := structures.NewDoubleLinkedList[int]()

	err := l.PushAt(-1, 1)
	if err == nil {
		t.Fatalf("doubleLinkedList.PushAt: Expected ErrIndexOutOfRangeInDoubleLinkedList, got nil")
	}
}

func TestDoubleLinkedList_PushAt_Empty(t *testing.T) {
	l := structures.NewDoubleLinkedList[int]()

	err := l.PushAt(0, 1)
	if err != nil {
		t.Fatalf("doubleLinkedList.PushAt: Expected nil, got %v", err)
	}
}

func TestDoubleLinkedList_PushAt_Empty_NegativeIndex(t *testing.T) {
	l := structures.NewDoubleLinkedList[int]()

	err := l.PushAt(-1, 1)
	if err == nil {
		t.Fatalf("doubleLinkedList.PushAt: Expected ErrIndexOutOfRangeInDoubleLinkedList, got nil")
	}
}

func TestDoubleLinkedList_PushAt_Empty_OutOfRange(t *testing.T) {
	l := structures.NewDoubleLinkedList[int]()

	err := l.PushAt(1, 1)
	if err == nil {
		t.Fatalf("doubleLinkedList.PushAt: Expected ErrIndexOutOfRangeInDoubleLinkedList, got nil")
	}
}
