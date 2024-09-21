package structures

import (
	"errors"
)

var (
	ErrPriorityQueueEmpty = errors.New("priority queue is empty")
)

// priorityQueue represents a priority queue with repeated elements allowed
type priorityQueue[T comparable] struct {
	tree    SearchBinaryTreer[T]
	counter map[T]int // Map to count occurrences of each element
}

// NewPriorityQueue creates a new priority queue
func NewPriorityQueue[T comparable](compare func(a, b T) int) PriorityQueuer[T] {
	return &priorityQueue[T]{
		tree:    NewSearchBinaryTree[T](compare),
		counter: make(map[T]int),
	}
}

// Push adds an element to the priority queue
func (pq *priorityQueue[T]) Push(value T) {
	if pq.counter[value] == 0 {
		pq.tree.Push(value) // Only add to tree if it's not already present
	}
	pq.counter[value]++ // Increment the counter
}

// Pop removes and returns the element with the highest priority
func (pq *priorityQueue[T]) Pop() (T, error) {
	if pq.tree.Root() == nil {
		var zero T
		return zero, ErrPriorityQueueEmpty
	}

	minNode := findMinSearchBinaryTreeNode(pq.tree.Root())
	value := minNode.value
	pq.counter[value]--

	if pq.counter[value] == 0 {
		err := pq.tree.Delete(value) // Remove from tree when no more instances are left
		delete(pq.counter, value)    // Remove the value from the counter map
		return value, err
	}

	return value, nil
}

// Top returns the element with the highest priority without removing it
func (pq *priorityQueue[T]) Top() (T, error) {
	if pq.tree.Root() == nil {
		var zero T
		return zero, ErrPriorityQueueEmpty
	}

	minNode := findMinSearchBinaryTreeNode(pq.tree.Root())
	return minNode.value, nil
}

// Size returns the number of elements in the priority queue (including repeated elements)
func (pq *priorityQueue[T]) Size() int64 {
	var size int64
	for _, count := range pq.counter {
		size += int64(count)
	}
	return size
}
