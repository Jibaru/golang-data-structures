package structures

// node element that point to next node and stores the node value
type node[T comparable] struct {
	value T
	next  *node[T]
}

type doubleNode[T comparable] struct {
	value T
	prev  *doubleNode[T]
	next  *doubleNode[T]
}

type TreeNode[T comparable] struct {
	value    T
	children []*TreeNode[T]
}

func (n *TreeNode[T]) Value() T {
	return n.value
}

func (n *TreeNode[T]) Children() []*TreeNode[T] {
	return n.children
}
