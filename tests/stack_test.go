package tests

import (
	"testing"

	"github.com/Jibaru/golang-data-structures/structures"
)

func TestStackPushIsTop(t *testing.T) {
	expected := 19
	stack := structures.NewStack[int]()

	stack.Push(17)
	stack.Push(18)
	stack.Push(19)

	got, err := stack.Top()
	if err != nil {
		printError("Top()", t, expected, err.Error())
	}

	if got != expected {
		printError("Top()", t, expected, got)
	}
}

func TestStackPop(t *testing.T) {
	expected := 19
	stack := structures.NewStack[int]()

	stack.Push(17)
	stack.Push(18)
	stack.Push(19)

	got, err := stack.Pop()
	if err != nil {
		printError("Pop()", t, expected, err.Error())
	}

	if got != expected {
		printError("Pop()", t, expected, got)
	}
}

func TestStackSize(t *testing.T) {
	var expected int64 = 3
	stack := structures.NewStack[int]()

	stack.Push(17)
	stack.Push(18)
	stack.Push(19)

	got := stack.Size()

	if got != expected {
		printError("Size()", t, expected, got)
	}
}

func TestStackNotFound(t *testing.T) {
	expected := false
	stack := structures.NewStack[int]()

	stack.Push(17)
	stack.Push(18)
	stack.Push(19)

	got := stack.Find(5)

	if got != expected {
		printError("Find(5)", t, expected, got)
	}
}

func TestStackPopErrEmptyStack(t *testing.T) {
	expected := structures.ErrEmptyStack
	stack := structures.NewStack[int]()

	_, err := stack.Pop()
	if err != nil && err != expected {
		printError("Pop()", t, expected, err.Error())
	}
}

func TestStackTopErrEmptyStack(t *testing.T) {
	expected := structures.ErrEmptyStack
	stack := structures.NewStack[int]()

	_, err := stack.Top()
	if err != nil && err != expected {
		printError("Top()", t, expected, err.Error())
	}
}

func TestStackBottomErrEmptyStack(t *testing.T) {
	expected := structures.ErrEmptyStack
	stack := structures.NewStack[int]()

	_, err := stack.Top()
	if err != nil && err != expected {
		printError("Bottom()", t, expected, err.Error())
	}
}
