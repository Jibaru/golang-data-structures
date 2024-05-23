package structures_test

import (
	"testing"

	"github.com/Jibaru/golang-data-structures/structures"
)

func TestNewSearchBinaryTree(t *testing.T) {
	tree := structures.NewSearchBinaryTree[int](func(a, b int) int {
		if a < b {
			return -1
		}
		if a > b {
			return 1
		}
		return 0
	})
	if tree == nil {
		t.Fatal("NewSearchBinaryTree should not return nil")
	}
}

func TestSearchBinaryTree_Push(t *testing.T) {
	tree := structures.NewSearchBinaryTree[int](func(a, b int) int {
		if a < b {
			return -1
		}
		if a > b {
			return 1
		}
		return 0
	})
	tree.Push(5)
	tree.Push(3)
	tree.Push(7)
	tree.Push(1)
	tree.Push(4)
	tree.Push(6)
	tree.Push(8)
}

func TestSearchBinaryTree_Push_Duplicate(t *testing.T) {
	tree := structures.NewSearchBinaryTree[int](func(a, b int) int {
		if a < b {
			return -1
		}
		if a > b {
			return 1
		}
		return 0
	})
	tree.Push(5)
	tree.Push(5)
}

func TestSearchBinaryTree_Delete(t *testing.T) {
	tree := structures.NewSearchBinaryTree[int](func(a, b int) int {
		if a < b {
			return -1
		}
		if a > b {
			return 1
		}
		return 0
	})
	nums := []int{5, 3, 7, 1, 4, 6, 8}

	for _, num := range nums {
		tree.Push(num)
	}

	for _, num := range nums {
		if err := tree.Delete(num); err != nil {
			t.Fatalf("Delete should not return an error for %d", num)
		}
	}
}

func TestSearchBinaryTree_Delete_Empty(t *testing.T) {
	tree := structures.NewSearchBinaryTree[int](func(a, b int) int {
		if a < b {
			return -1
		}
		if a > b {
			return 1
		}
		return 0
	})

	if err := tree.Delete(0); err == nil {
		t.Fatal("Delete should return an error")
	}
}

func TestSearchBinaryTree_Find(t *testing.T) {
	tree := structures.NewSearchBinaryTree[int](func(a, b int) int {
		if a < b {
			return -1
		}
		if a > b {
			return 1
		}
		return 0
	})

	nums := []int{5, 3, 7, 1, 4, 6, 8}

	for _, num := range nums {
		tree.Push(num)
	}

	for _, num := range nums {
		if !tree.Find(num) {
			t.Fatalf("Find should return true for %d", num)
		}
	}
}

func TestSearchBinaryTree_Find_NotFound(t *testing.T) {
	tree := structures.NewSearchBinaryTree[int](func(a, b int) int {
		if a < b {
			return -1
		}
		if a > b {
			return 1
		}
		return 0
	})

	nums := []int{5, 3, 7, 1, 4, 6, 8}

	for _, num := range nums {
		tree.Push(num)
	}

	if tree.Find(0) {
		t.Fatal("Find should return false for 0")
	}
}

func TestSearchBinaryTree_Find_Empty(t *testing.T) {
	tree := structures.NewSearchBinaryTree[int](func(a, b int) int {
		if a < b {
			return -1
		}
		if a > b {
			return 1
		}
		return 0
	})

	if tree.Find(0) {
		t.Fatal("Find should return false for 0")
	}
}

func TestSearchBinaryTree_Size(t *testing.T) {
	tree := structures.NewSearchBinaryTree[int](func(a, b int) int {
		if a < b {
			return -1
		}
		if a > b {
			return 1
		}
		return 0
	})

	if tree.Size() != 0 {
		t.Fatal("Size should be 0")
	}

	tree.Push(5)
	if tree.Size() != 1 {
		t.Fatal("Size should be 1")
	}

	tree.Push(3)
	if tree.Size() != 2 {
		t.Fatal("Size should be 2")
	}

	tree.Push(7)
	if tree.Size() != 3 {
		t.Fatal("Size should be 3")
	}

	tree.Push(1)
	if tree.Size() != 4 {
		t.Fatal("Size should be 4")
	}

	tree.Push(4)
	if tree.Size() != 5 {
		t.Fatal("Size should be 5")
	}

	tree.Push(6)
	if tree.Size() != 6 {
		t.Fatal("Size should be 6")
	}

	tree.Push(8)
	if tree.Size() != 7 {
		t.Fatal("Size should be 7")
	}
}
