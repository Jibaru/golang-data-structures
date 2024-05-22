package structures_test

import (
	"testing"

	"github.com/Jibaru/golang-data-structures/structures"
)

func TestNewTree(t *testing.T) {
	tree := structures.NewTree[int](1)

	if tree.Root().Value() != 1 {
		t.Fatalf("NewTree: Expected 1, got %v", tree.Root().Value())
	}
}

func TestTree_PushAt(t *testing.T) {
	tree := structures.NewTree[int](1)

	_, err := tree.PushAt(tree.Root(), 2)
	if err != nil {
		t.Fatalf("Tree.PushAt: Expected nil, got %v", err)
	}

	if tree.Root().Children()[0].Value() != 2 {
		t.Fatalf("Tree.PushAt: Expected 2, got %v", tree.Root().Children()[0].Value())
	}
}

func TestTree_PushAt_Nil(t *testing.T) {
	tree := structures.NewTree[int](1)

	_, err := tree.PushAt(nil, 2)
	if err == nil {
		t.Fatalf("Tree.PushAt: Expected ErrInputNodeNil, got nil")
	}
}

func TestTree_Root(t *testing.T) {
	tree := structures.NewTree[int](1)

	if tree.Root().Value() != 1 {
		t.Fatalf("Tree.Root: Expected 1, got %v", tree.Root().Value())
	}
}

func TestTree_Leaves(t *testing.T) {
	tree := structures.NewTree[int](1)

	leaves := tree.Leaves()
	if len(leaves) != 1 {
		t.Fatalf("Tree.Leaves: Expected 1, got %v", len(leaves))
	}

	if leaves[0].Value() != 1 {
		t.Fatalf("Tree.Leaves: Expected 1, got %v", leaves[0].Value())
	}
}

func TestTree_DeleteWithDescendants(t *testing.T) {
	tree := structures.NewTree[int](1)
	tree.PushAt(tree.Root(), 2)
	tree.PushAt(tree.Root().Children()[0], 3)

	err := tree.DeleteWithDescendants(tree.Root().Children()[0])
	if err != nil {
		t.Fatalf("Tree.DeleteWithDescendants: Expected nil, got %v", err)
	}

	if len(tree.Root().Children()) != 0 {
		t.Fatalf("Tree.DeleteWithDescendants: Expected 0, got %v", len(tree.Root().Children()))
	}
}

func TestTree_DeleteWithDescendants_Nil(t *testing.T) {
	tree := structures.NewTree[int](1)

	err := tree.DeleteWithDescendants(nil)
	if err == nil {
		t.Fatalf("Tree.DeleteWithDescendants: Expected ErrInputNodeNil, got nil")
	}
}

func TestTree_DeleteWithDescendants_NotFound(t *testing.T) {
	tree := structures.NewTree[int](1)
	tree.PushAt(tree.Root(), 2)

	err := tree.DeleteWithDescendants(&structures.TreeNode[int]{})
	if err == nil {
		t.Fatalf("Tree.DeleteWithDescendants: Expected ErrNodeNotFound, got nil")
	}
}

func TestTree_DeleteWithDescendants_Root(t *testing.T) {
	tree := structures.NewTree[int](1)

	err := tree.DeleteWithDescendants(tree.Root())
	if err == nil {
		t.Fatalf("Tree.DeleteWithDescendants: Expected ErrDeleteRoot, got nil")
	}
}

func TestTree_DeleteWithDescendants_Leaves(t *testing.T) {
	tree := structures.NewTree[int](1)
	tree.PushAt(tree.Root(), 2)
	tree.PushAt(tree.Root().Children()[0], 3)

	err := tree.DeleteWithDescendants(tree.Root().Children()[0])
	if err != nil {
		t.Fatalf("Tree.DeleteWithDescendants: Expected nil, got %v", err)
	}

	leaves := tree.Leaves()
	if len(leaves) != 1 {
		t.Fatalf("Tree.DeleteWithDescendants: Expected 1, got %v", len(leaves))
	}

	if leaves[0].Value() != 1 {
		t.Fatalf("Tree.DeleteWithDescendants: Expected 1, got %v", leaves[0].Value())
	}
}

func TestTree_DeleteWithoutDescendants(t *testing.T) {
	tree := structures.NewTree[int](1)
	tree.PushAt(tree.Root(), 2)
	tree.PushAt(tree.Root().Children()[0], 3)

	err := tree.DeleteWithoutDescendants(tree.Root().Children()[0])
	if err != nil {
		t.Fatalf("Tree.DeleteWithoutDescendants: Expected nil, got %v", err)
	}

	if tree.Root().Children()[0].Value() != 3 {
		t.Fatalf("Tree.DeleteWithoutDescendants: Expected 3, got %v", tree.Root().Children()[0].Value())
	}
}

func TestTree_DeleteWithoutDescendants_Nil(t *testing.T) {
	tree := structures.NewTree[int](1)

	err := tree.DeleteWithoutDescendants(nil)
	if err == nil {
		t.Fatalf("Tree.DeleteWithoutDescendants: Expected ErrInputNodeNil, got nil")
	}
}

func TestTree_DeleteWithoutDescendants_NotFound(t *testing.T) {
	tree := structures.NewTree[int](1)
	tree.PushAt(tree.Root(), 2)

	err := tree.DeleteWithoutDescendants(&structures.TreeNode[int]{})
	if err == nil {
		t.Fatalf("Tree.DeleteWithoutDescendants: Expected ErrNodeNotFound, got nil")
	}
}

func TestTree_DeleteWithoutDescendants_Root(t *testing.T) {
	tree := structures.NewTree[int](1)

	err := tree.DeleteWithoutDescendants(tree.Root())
	if err == nil {
		t.Fatalf("Tree.DeleteWithoutDescendants: Expected ErrDeleteRoot, got nil")
	}
}

func TestTree_Size(t *testing.T) {
	tree := structures.NewTree[int](1)
	tree.PushAt(tree.Root(), 2)
	tree.PushAt(tree.Root().Children()[0], 3)

	if tree.Size() != 3 {
		t.Fatalf("Tree.Size: Expected 3, got %v", tree.Size())
	}
}

func TestTree_DepthFirstSearch(t *testing.T) {
	tree := structures.NewTree[int](1)
	tree.PushAt(tree.Root(), 2)
	tree.PushAt(tree.Root().Children()[0], 3)

	if !tree.DepthFirstSearch(3) {
		t.Fatalf("Tree.DepthFirstSearch: Expected true, got false")
	}
}

func TestTree_DepthFirstSearch_NotFound(t *testing.T) {
	tree := structures.NewTree[int](1)
	tree.PushAt(tree.Root(), 2)
	tree.PushAt(tree.Root().Children()[0], 3)

	if tree.DepthFirstSearch(4) {
		t.Fatalf("Tree.DepthFirstSearch: Expected false, got true")
	}
}

func TestTree_BreadFirstSearch(t *testing.T) {
	tree := structures.NewTree[int](1)
	tree.PushAt(tree.Root(), 2)
	tree.PushAt(tree.Root().Children()[0], 3)

	if !tree.BreadFirstSearch(3) {
		t.Fatalf("Tree.BreadFirstSearch: Expected true, got false")
	}
}

func TestTree_BreadFirstSearch_NotFound(t *testing.T) {
	tree := structures.NewTree[int](1)
	tree.PushAt(tree.Root(), 2)
	tree.PushAt(tree.Root().Children()[0], 3)

	if tree.BreadFirstSearch(4) {
		t.Fatalf("Tree.BreadFirstSearch: Expected false, got true")
	}
}
