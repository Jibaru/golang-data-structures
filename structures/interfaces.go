package structures

import "golang.org/x/exp/constraints"

// Sizer define the size of the structure
type Sizer[T comparable] interface {
	Size() int64
}

// Finder define a way to check if an element exists or not
type Finder[T comparable] interface {
	Find(value T) bool
}

// Stacker define the basic operations of a stack
type Stacker[T comparable] interface {
	Sizer[T]
	Finder[T]
	Pop() (T, error)
	Push(value T)
	Top() (T, error)
	Bottom() (T, error)
}

// Queuer define the basic operations of a queue
type Queuer[T comparable] interface {
	Sizer[T]
	Finder[T]
	Pop() (T, error)
	Push(value T)
	Top() (T, error)
	Bottom() (T, error)
}

// SimpleLinkedLister define the basic operations of a simple linked list
type SimpleLinkedLister[T comparable] interface {
	Sizer[T]
	Finder[T]
	PushAt(index int, value T) error
	PopAt(index int) (T, error)
	First() (T, error)
	Last() (T, error)
	GetAt(index int) (T, error)
}

// DoubleLinkedLister define the basic operations of a double linked list
type DoubleLinkedLister[T comparable] interface {
	Sizer[T]
	Finder[T]
	PushAt(index int, value T) error
	PopAt(index int) (T, error)
	First() (T, error)
	Last() (T, error)
	GetAt(index int) (T, error)
}

// HashMapper define the basic operations of a hash map
type HashMapper[K comparable, V comparable] interface {
	Finder[V]
	Sizer[V]
	PushAt(key K, value V)
	PopAt(key K) (V, error)
	Has(key K) bool
	GetAt(key K) (V, error)
}

// BasicTree define the basic operations of a tree
type BasicTree[T comparable] interface {
	Sizer[T]
	PushAt(in *TreeNode[T], val T) (*TreeNode[T], error)
	DeleteWithDescendants(node *TreeNode[T]) error
	DeleteWithoutDescendants(node *TreeNode[T]) error
	Root() *TreeNode[T]
	Leaves() []*TreeNode[T]
	BreadFirstSearch(value T) bool
	DepthFirstSearch(value T) bool
}

type SearchBinaryTreer[T comparable] interface {
	Finder[T]
	Sizer[T]
	Push(value T)
	Delete(value T) error
	Root() *SearchBinaryTreeNode[T]
}

type PriorityQueuer[T comparable] interface {
	Sizer[T]
	Push(value T)
	Pop() (T, error)
	Top() (T, error)
}

type Numeric interface {
	constraints.Integer | constraints.Float
}

// Graph represents a generic weighted graph with advanced operations
type Graph[T comparable, W Numeric] interface {
	AddVertex(vertex T) error
	RemoveVertex(vertex T) error
	AddEdge(from, to T, weight W) error
	RemoveEdge(from, to T) error
	HasEdge(from, to T) bool
	Neighbors(vertex T) (map[T]W, error)
	GetWeight(from, to T) (W, error)
	GetVertices() []T
	GetEdges() []Edge[T, W]
	Degree(vertex T) (int, error)
	InDegree(vertex T) (int, error)
	Transpose() Graph[T, W]
	IsDirected() bool
	ShortestPath(from, to T) ([]T, error)
}
