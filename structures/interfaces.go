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
