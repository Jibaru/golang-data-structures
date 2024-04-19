package structures_test

import (
	"github.com/Jibaru/golang-data-structures/structures"
	"testing"
)

func TestNewSimpleLinkedList(t *testing.T) {
	l := structures.NewSimpleLinkedList[int]()

	if l.Size() != 0 {
		t.Fatalf("NewSimpleLinkedList: Expected 0, got %v", l.Size())
	}
}

func TestSimpleLinkedList_PushAt(t *testing.T) {
	l := structures.NewSimpleLinkedList[int]()

	err := l.PushAt(0, 1)
	if err != nil {
		t.Fatalf("simpleLinkedList.PushAt: Expected nil, got %v", err)
	}

	err = l.PushAt(1, 2)
	if err != nil {
		t.Fatalf("simpleLinkedList.PushAt: Expected nil, got %v", err)
	}

	err = l.PushAt(2, 3)
	if err != nil {
		t.Fatalf("simpleLinkedList.PushAt: Expected nil, got %v", err)
	}

	if l.Size() != 3 {
		t.Fatalf("simpleLinkedList.PushAt: Expected 3, got %v", l.Size())
	}

	got, err := l.GetAt(0)
	if err != nil {
		t.Fatalf("simpleLinkedList.GetAt: Expected nil, got %v", err)
	}

	if got != 1 {
		t.Fatalf("simpleLinkedList.GetAt: Expected 1, got %v", got)
	}

	got, err = l.GetAt(1)
	if err != nil {
		t.Fatalf("simpleLinkedList.GetAt: Expected nil, got %v", err)
	}

	if got != 2 {
		t.Fatalf("simpleLinkedList.GetAt: Expected 2, got %v", got)
	}

	got, err = l.GetAt(2)
	if err != nil {
		t.Fatalf("simpleLinkedList.GetAt: Expected nil, got %v", err)
	}

	if got != 3 {
		t.Fatalf("simpleLinkedList.GetAt: Expected 3, got %v", got)
	}
}

func TestSimpleLinkedList_PopAt(t *testing.T) {
	l := structures.NewSimpleLinkedList[int]()

	err := l.PushAt(0, 1)
	if err != nil {
		t.Fatalf("simpleLinkedList.PushAt: Expected nil, got %v", err)
	}

	err = l.PushAt(1, 2)
	if err != nil {
		t.Fatalf("simpleLinkedList.PushAt: Expected nil, got %v", err)
	}

	err = l.PushAt(2, 3)
	if err != nil {
		t.Fatalf("simpleLinkedList.PushAt: Expected nil, got %v", err)
	}

	v, err := l.PopAt(0)
	if err != nil {
		t.Fatalf("simpleLinkedList.PopAt: Expected nil, got %v", err)
	}

	if v != 1 {
		t.Fatalf("simpleLinkedList.PopAt: Expected 1, got %v", v)
	}

	if l.Size() != 2 {
		t.Fatalf("simpleLinkedList.PopAt: Expected 2, got %v", l.Size())
	}

	v, err = l.PopAt(1)
	if err != nil {
		t.Fatalf("simpleLinkedList.PopAt: Expected nil, got %v", err)
	}

	if v != 3 {
		t.Fatalf("simpleLinkedList.PopAt: Expected 3, got %v", v)
	}

	if l.Size() != 1 {
		t.Fatalf("simpleLinkedList.PopAt: Expected 1, got %v", l.Size())
	}
}

func TestSimpleLinkedList_PopAt_Empty(t *testing.T) {
	l := structures.NewSimpleLinkedList[int]()

	_, err := l.PopAt(0)
	if err == nil {
		t.Fatalf("simpleLinkedList.PopAt: Expected ErrEmptySimpleLinkedList, got nil")
	}
}

func TestSimpleLinkedList_PopAt_OutOfRange(t *testing.T) {
	l := structures.NewSimpleLinkedList[int]()

	err := l.PushAt(0, 1)
	if err != nil {
		t.Fatalf("simpleLinkedList.PushAt: Expected nil, got %v", err)
	}

	_, err = l.PopAt(1)
	if err == nil {
		t.Fatalf("simpleLinkedList.PopAt: Expected ErrIndexOutOfRangeInSimpleLinkedList, got nil")
	}
}

func TestSimpleLinkedList_First(t *testing.T) {
	l := structures.NewSimpleLinkedList[int]()

	err := l.PushAt(0, 1)
	if err != nil {
		t.Fatalf("simpleLinkedList.PushAt: Expected nil, got %v", err)
	}

	v, err := l.First()
	if err != nil {
		t.Fatalf("simpleLinkedList.First: Expected nil, got %v", err)
	}

	if v != 1 {
		t.Fatalf("simpleLinkedList.First: Expected 1, got %v", v)
	}
}

func TestSimpleLinkedList_First_Empty(t *testing.T) {
	l := structures.NewSimpleLinkedList[int]()

	_, err := l.First()
	if err == nil {
		t.Fatalf("simpleLinkedList.First: Expected ErrEmptySimpleLinkedList, got nil")
	}
}

func TestSimpleLinkedList_Last(t *testing.T) {
	l := structures.NewSimpleLinkedList[int]()

	err := l.PushAt(0, 1)
	if err != nil {
		t.Fatalf("simpleLinkedList.PushAt: Expected nil, got %v", err)
	}

	v, err := l.Last()
	if err != nil {
		t.Fatalf("simpleLinkedList.Last: Expected nil, got %v", err)
	}

	if v != 1 {
		t.Fatalf("simpleLinkedList.Last: Expected 1, got %v", v)
	}
}

func TestSimpleLinkedList_Last_Empty(t *testing.T) {
	l := structures.NewSimpleLinkedList[int]()

	_, err := l.Last()
	if err == nil {
		t.Fatalf("simpleLinkedList.Last: Expected ErrEmptySimpleLinkedList, got nil")
	}
}

func TestSimpleLinkedList_Find(t *testing.T) {
	l := structures.NewSimpleLinkedList[int]()

	err := l.PushAt(0, 1)
	if err != nil {
		t.Fatalf("simpleLinkedList.PushAt: Expected nil, got %v", err)
	}

	err = l.PushAt(1, 2)
	if err != nil {
		t.Fatalf("simpleLinkedList.PushAt: Expected nil, got %v", err)
	}

	err = l.PushAt(2, 3)
	if err != nil {
		t.Fatalf("simpleLinkedList.PushAt: Expected nil, got %v", err)
	}

	if !l.Find(1) {
		t.Fatalf("simpleLinkedList.Find: Expected true, got false")
	}

	if !l.Find(2) {
		t.Fatalf("simpleLinkedList.Find: Expected true, got false")
	}

	if !l.Find(3) {
		t.Fatalf("simpleLinkedList.Find: Expected true, got false")
	}

	if l.Find(4) {
		t.Fatalf("simpleLinkedList.Find: Expected false, got true")
	}
}

func TestSimpleLinkedList_Find_Empty(t *testing.T) {
	l := structures.NewSimpleLinkedList[int]()

	if l.Find(1) {
		t.Fatalf("simpleLinkedList.Find: Expected false, got true")
	}
}

func TestSimpleLinkedList_Size(t *testing.T) {
	l := structures.NewSimpleLinkedList[int]()

	err := l.PushAt(0, 1)
	if err != nil {
		t.Fatalf("simpleLinkedList.PushAt: Expected nil, got %v", err)
	}

	err = l.PushAt(1, 2)
	if err != nil {
		t.Fatalf("simpleLinkedList.PushAt: Expected nil, got %v", err)
	}

	err = l.PushAt(2, 3)
	if err != nil {
		t.Fatalf("simpleLinkedList.PushAt: Expected nil, got %v", err)
	}

	if l.Size() != 3 {
		t.Fatalf("simpleLinkedList.Size: Expected 3, got %v", l.Size())
	}
}

func TestSimpleLinkedList_Size_Empty(t *testing.T) {
	l := structures.NewSimpleLinkedList[int]()

	if l.Size() != 0 {
		t.Fatalf("simpleLinkedList.Size: Expected 0, got %v", l.Size())
	}
}

func TestSimpleLinkedList_GetAt(t *testing.T) {
	l := structures.NewSimpleLinkedList[int]()

	err := l.PushAt(0, 1)
	if err != nil {
		t.Fatalf("simpleLinkedList.PushAt: Expected nil, got %v", err)
	}

	err = l.PushAt(1, 2)
	if err != nil {
		t.Fatalf("simpleLinkedList.PushAt: Expected nil, got %v", err)
	}

	err = l.PushAt(2, 3)
	if err != nil {
		t.Fatalf("simpleLinkedList.PushAt: Expected nil, got %v", err)
	}

	v, err := l.GetAt(0)
	if err != nil {
		t.Fatalf("simpleLinkedList.GetAt: Expected nil, got %v", err)
	}

	if v != 1 {
		t.Fatalf("simpleLinkedList.GetAt: Expected 1, got %v", v)
	}

	v, err = l.GetAt(1)
	if err != nil {
		t.Fatalf("simpleLinkedList.GetAt: Expected nil, got %v", err)
	}

	if v != 2 {
		t.Fatalf("simpleLinkedList.GetAt: Expected 2, got %v", v)
	}

	v, err = l.GetAt(2)
	if err != nil {
		t.Fatalf("simpleLinkedList.GetAt: Expected nil, got %v", err)
	}

	if v != 3 {
		t.Fatalf("simpleLinkedList.GetAt: Expected 3, got %v", v)
	}
}

func TestSimpleLinkedList_GetAt_Empty(t *testing.T) {
	l := structures.NewSimpleLinkedList[int]()

	_, err := l.GetAt(0)
	if err == nil {
		t.Fatalf("simpleLinkedList.GetAt: Expected ErrEmptySimpleLinkedList, got nil")
	}
}

func TestSimpleLinkedList_GetAt_OutOfRange(t *testing.T) {
	l := structures.NewSimpleLinkedList[int]()

	err := l.PushAt(0, 1)
	if err != nil {
		t.Fatalf("simpleLinkedList.PushAt: Expected nil, got %v", err)
	}

	_, err = l.GetAt(1)
	if err == nil {
		t.Fatalf("simpleLinkedList.GetAt: Expected ErrIndexOutOfRangeInSimpleLinkedList, got nil")
	}
}

func TestSimpleLinkedList_PushAt_OutOfRange(t *testing.T) {
	l := structures.NewSimpleLinkedList[int]()

	err := l.PushAt(1, 1)
	if err == nil {
		t.Fatalf("simpleLinkedList.PushAt: Expected ErrIndexOutOfRangeInSimpleLinkedList, got nil")
	}
}

func TestSimpleLinkedList_PushAt_NegativeIndex(t *testing.T) {
	l := structures.NewSimpleLinkedList[int]()

	err := l.PushAt(-1, 1)
	if err == nil {
		t.Fatalf("simpleLinkedList.PushAt: Expected ErrIndexOutOfRangeInSimpleLinkedList, got nil")
	}
}

func TestSimpleLinkedList_PushAt_Empty(t *testing.T) {
	l := structures.NewSimpleLinkedList[int]()

	err := l.PushAt(0, 1)
	if err != nil {
		t.Fatalf("simpleLinkedList.PushAt: Expected nil, got %v", err)
	}
}

func TestSimpleLinkedList_PushAt_Empty_NegativeIndex(t *testing.T) {
	l := structures.NewSimpleLinkedList[int]()

	err := l.PushAt(-1, 1)
	if err == nil {
		t.Fatalf("simpleLinkedList.PushAt: Expected ErrIndexOutOfRangeInSimpleLinkedList, got nil")
	}
}

func TestSimpleLinkedList_PushAt_Empty_OutOfRange(t *testing.T) {
	l := structures.NewSimpleLinkedList[int]()

	err := l.PushAt(1, 1)
	if err == nil {
		t.Fatalf("simpleLinkedList.PushAt: Expected ErrIndexOutOfRangeInSimpleLinkedList, got nil")
	}
}
