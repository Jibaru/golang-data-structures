package structures

import (
	"errors"
)

var (
	ErrorSearchBinaryTreeValueNotFound = errors.New("value not found")
)

// searchBinaryTree represents a binary search tree
type searchBinaryTree[T comparable] struct {
	root    *SearchBinaryTreeNode[T]
	compare func(a, b T) int
}

// NewSearchBinaryTree creates a new binary search tree
func NewSearchBinaryTree[T comparable](compare func(a, b T) int) SearchBinaryTreer[T] {
	return &searchBinaryTree[T]{compare: compare}
}

// Push inserts a new value into the binary search tree
func (tree *searchBinaryTree[T]) Push(value T) {
	tree.root = pushSearchBinaryTreeNode(tree.root, value, tree.compare)
}

// pushSearchBinaryTreeNode inserts a new value into the binary search tree
func pushSearchBinaryTreeNode[T comparable](
	node *SearchBinaryTreeNode[T],
	value T,
	compare func(a, b T) int,
) *SearchBinaryTreeNode[T] {
	if node == nil {
		return &SearchBinaryTreeNode[T]{value: value}
	}

	if compare(value, node.value) < 0 {
		node.left = pushSearchBinaryTreeNode(node.left, value, compare)
	} else if compare(value, node.value) > 0 {
		node.right = pushSearchBinaryTreeNode(node.right, value, compare)
	}
	return node
}

// Find finds a value in the binary search tree
func (tree *searchBinaryTree[T]) Find(value T) bool {
	return findNode(tree.root, value, tree.compare)
}

// findNode finds a value in the binary search tree
func findNode[T comparable](
	node *SearchBinaryTreeNode[T],
	value T,
	compare func(a, b T) int,
) bool {
	if node == nil {
		return false
	}

	if compare(value, node.value) < 0 {
		return findNode(node.left, value, compare)
	} else if compare(value, node.value) > 0 {
		return findNode(node.right, value, compare)
	}
	return true
}

// Delete deletes a value from the binary search tree
func (tree *searchBinaryTree[T]) Delete(value T) error {
	var err error
	tree.root, err = deleteSearchBinaryTreeNode(tree.root, value, tree.compare)
	return err
}

// deleteSearchBinaryTreeNode deletes a value from the binary search tree
func deleteSearchBinaryTreeNode[T comparable](
	node *SearchBinaryTreeNode[T],
	value T,
	compare func(a, b T) int,
) (*SearchBinaryTreeNode[T], error) {
	if node == nil {
		return nil, ErrorSearchBinaryTreeValueNotFound
	}

	if compare(value, node.value) < 0 {
		node.left, _ = deleteSearchBinaryTreeNode(node.left, value, compare)
	} else if compare(value, node.value) > 0 {
		node.right, _ = deleteSearchBinaryTreeNode(node.right, value, compare)
	} else {
		// Node with only one child or no child
		if node.left == nil {
			return node.right, nil
		} else if node.right == nil {
			return node.left, nil
		}

		// Node with two children: Get the inorder successor (smallest in the right subtree)
		minValueNode := findMinSearchBinaryTreeNode(node.right)
		node.value = minValueNode.value
		node.right, _ = deleteSearchBinaryTreeNode(node.right, minValueNode.value, compare)
	}
	return node, nil
}

// findMinSearchBinaryTreeNode finds the node with the minimum value in the binary search tree
func findMinSearchBinaryTreeNode[T comparable](node *SearchBinaryTreeNode[T]) *SearchBinaryTreeNode[T] {
	current := node
	for current != nil && current.left != nil {
		current = current.left
	}
	return current
}

// Size returns the size of the binary search tree
func (tree *searchBinaryTree[T]) Size() int64 {
	return size(tree.root)
}

// size returns the size of the binary search tree
func size[T comparable](node *SearchBinaryTreeNode[T]) int64 {
	if node == nil {
		return 0
	}
	return 1 + size(node.left) + size(node.right)
}

func (tree *searchBinaryTree[T]) Root() *SearchBinaryTreeNode[T] {
	return tree.root
}
