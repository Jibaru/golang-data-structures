package structures

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

// HashMapper define the basic operations of a hash map
type HashMapper[K comparable, V comparable] interface {
	Finder[V]
	Sizer[V]
	PushAt(key K, value V)
	PopAt(key K) (V, error)
	Has(key K) bool
	GetAt(key K) (V, error)
}
