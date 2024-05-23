package structures

import "errors"

var (
	ErrPriorityQueueEmpty = errors.New("priority queue is empty")
)

// priorityQueue represents a priority queue
type priorityQueue[T comparable] struct {
	tree SearchBinaryTreer[T]
}

// NewPriorityQueue creates a new priority queue
func NewPriorityQueue[T comparable](compare func(a, b T) int) PriorityQueuer[T] {
	return &priorityQueue[T]{tree: NewSearchBinaryTree[T](compare)}
}

// Push adds an element to the priority queue
func (pq *priorityQueue[T]) Push(value T) {
	pq.tree.Push(value)
}

// Pop removes and returns the element with the highest priority
func (pq *priorityQueue[T]) Pop() (T, error) {
	if pq.tree.Root() == nil {
		var zero T
		return zero, ErrPriorityQueueEmpty
	}

	minNode := findMinSearchBinaryTreeNode(pq.tree.Root())
	err := pq.tree.Delete(minNode.value)
	return minNode.value, err
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

// Size returns the number of elements in the priority queue
func (pq *priorityQueue[T]) Size() int64 {
	return pq.tree.Size()
}
