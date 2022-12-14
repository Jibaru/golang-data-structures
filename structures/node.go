package structures

// node element that point to next node and stores the node value
type node[T comparable] struct {
	value T
	next  *node[T]
}
