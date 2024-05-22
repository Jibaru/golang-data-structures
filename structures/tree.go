package structures

import (
	"errors"
)

var (
	ErrEmptyTree        = errors.New("empty tree")
	ErrInputNodeNil     = errors.New("input node is nil")
	ErrCannotDeleteRoot = errors.New("cannot delete the root node without its descendants")
	ErrNodeNotFound     = errors.New("node not found")
)

type Tree[T comparable] struct {
	root *TreeNode[T]
}

// NewTree creates a new tree with a root node
func NewTree[T comparable](value T) BasicTree[T] {
	return &Tree[T]{
		root: &TreeNode[T]{value: value},
	}
}

// PushAt insert a new node with value val as a child of in node
func (t *Tree[T]) PushAt(in *TreeNode[T], val T) (*TreeNode[T], error) {
	if in == nil {
		return nil, ErrInputNodeNil
	}
	newTreeNode := &TreeNode[T]{value: val}
	in.children = append(in.children, newTreeNode)
	return newTreeNode, nil
}

// DeleteWithDescendants delete a node and all its subnodes
func (t *Tree[T]) DeleteWithDescendants(node *TreeNode[T]) error {
	if t.root == nil {
		return ErrEmptyTree
	}

	if node == nil {
		return ErrInputNodeNil
	}

	if t.root == node {
		return ErrCannotDeleteRoot
	}

	parent, err := findParent(t.root, node)
	if err != nil {
		return err
	}

	// Remove the node from its parent's children
	for i, child := range parent.children {
		if child == node {
			parent.children = append(parent.children[:i], parent.children[i+1:]...)
			break
		}
	}

	return nil
}

// DeleteWithoutDescendants delete a node and pass its children to its parent
func (t *Tree[T]) DeleteWithoutDescendants(node *TreeNode[T]) error {
	if t.root == nil {
		return ErrEmptyTree
	}
	if node == nil {
		return ErrInputNodeNil
	}

	if t.root == node {
		return ErrCannotDeleteRoot
	}

	parent, err := findParent(t.root, node)
	if err != nil {
		return err
	}

	// Find the index of the node to delete
	var index int
	for i, child := range parent.children {
		if child == node {
			index = i
			break
		}
	}

	// Remove the node and append its children to the parent
	parent.children = append(parent.children[:index], parent.children[index+1:]...)
	parent.children = append(parent.children, node.children...)

	return nil
}

func findParent[T comparable](root, node *TreeNode[T]) (*TreeNode[T], error) {
	if root == nil {
		return nil, ErrInputNodeNil
	}
	for _, child := range root.children {
		if child == node {
			return root, nil
		}
		parent, err := findParent(child, node)
		if err == nil {
			return parent, nil
		}
	}
	return nil, ErrNodeNotFound
}

// Get the root node of the tree
func (t *Tree[T]) Root() *TreeNode[T] {
	return t.root
}

// Leaves returns all leaf nodes in the tree
func (t *Tree[T]) Leaves() []*TreeNode[T] {
	var leaves []*TreeNode[T]
	collectLeaves(t.root, &leaves)
	return leaves
}

func collectLeaves[T comparable](node *TreeNode[T], leaves *[]*TreeNode[T]) {
	if node == nil {
		return
	}
	if len(node.children) == 0 {
		*leaves = append(*leaves, node)
	}
	for _, child := range node.children {
		collectLeaves(child, leaves)
	}
}

// Size returns the total number of nodes in the tree
func (t *Tree[T]) Size() int64 {
	return countTreeNodes(t.root)
}

func countTreeNodes[T comparable](node *TreeNode[T]) int64 {
	if node == nil {
		return 0
	}
	var count int64 = 1
	for _, child := range node.children {
		count += countTreeNodes(child)
	}
	return count
}

// DepthFirstSearch find an element
func (t *Tree[T]) DepthFirstSearch(value T) bool {
	return findDFS(t.root, value)
}

func findDFS[T comparable](node *TreeNode[T], value T) bool {
	if node == nil {
		return false
	}
	if node.value == value {
		return true
	}
	for _, child := range node.children {
		if findDFS(child, value) {
			return true
		}
	}
	return false
}

// BreadFirstSearch find an element
func (t *Tree[T]) BreadFirstSearch(value T) bool {
	if t.root == nil {
		return false
	}
	queue := []*TreeNode[T]{t.root}
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		if node.value == value {
			return true
		}
		queue = append(queue, node.children...)
	}
	return false
}
