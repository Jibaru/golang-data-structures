package structures

// node element that point to next node and stores the node value
type node[T comparable] struct {
	value T
	next  *node[T]
}

// doubleNode element that point to next and previous node and stores the node value
type doubleNode[T comparable] struct {
	value T
	prev  *doubleNode[T]
	next  *doubleNode[T]
}

// TreeNode is a node in a tree that stores a value and a list of children nodes
type TreeNode[T comparable] struct {
	value    T
	children []*TreeNode[T]
}

// Value returns the value of the node
func (n *TreeNode[T]) Value() T {
	return n.value
}

// Children returns a slice of children nodes
func (n *TreeNode[T]) Children() []*TreeNode[T] {
	return n.children
}

// SearchBinaryTreeNode represents a node in the search binary tree
type SearchBinaryTreeNode[T any] struct {
	value T
	left  *SearchBinaryTreeNode[T]
	right *SearchBinaryTreeNode[T]
}
