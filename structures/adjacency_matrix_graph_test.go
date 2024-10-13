package structures_test

import (
	"errors"
	"testing"

	"github.com/Jibaru/golang-data-structures/structures"
)

func TestNewAdjacencyMatrixGraph(t *testing.T) {
	graph := structures.NewAdjacencyMatrixGraph[int, int](true)
	if graph == nil {
		t.Fatal("structures.structures.NewAdjacencyMatrixGraph should not return nil")
	}
}

func TestAdjacencyMatrixGraph_AddVertex(t *testing.T) {
	graph := structures.NewAdjacencyMatrixGraph[string, int](false)
	err := graph.AddVertex("A")
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	err = graph.AddVertex("A")
	if !errors.Is(err, structures.ErrVertexAlreadyExists) {
		t.Fatalf("expected structures.ErrVertexNotFound, got %v", err)
	}

	vertices := graph.GetVertices()
	if len(vertices) != 1 || vertices[0] != "A" {
		t.Fatalf("expected vertex A to be added, got %v", vertices)
	}
}

func TestAdjacencyMatrixGraph_RemoveVertex(t *testing.T) {
	graph := structures.NewAdjacencyMatrixGraph[string, int](false)
	_ = graph.AddVertex("A")
	_ = graph.AddVertex("B")

	err := graph.RemoveVertex("A")
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	err = graph.RemoveVertex("A")
	if !errors.Is(err, structures.ErrVertexNotFound) {
		t.Fatalf("expected structures.ErrVertexNotFound, got %v", err)
	}

	vertices := graph.GetVertices()
	if len(vertices) != 1 || vertices[0] != "B" {
		t.Fatalf("expected vertex B to remain, got %v", vertices)
	}
}

func TestAdjacencyMatrixGraph_AddEdge(t *testing.T) {
	graph := structures.NewAdjacencyMatrixGraph[string, int](false)
	_ = graph.AddVertex("A")
	_ = graph.AddVertex("B")

	err := graph.AddEdge("A", "B", 5)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	weight, err := graph.GetWeight("A", "B")
	if err != nil || weight != 5 {
		t.Fatalf("expected weight 5, got %v (error: %v)", weight, err)
	}
}

func TestAdjacencyMatrixGraph_RemoveEdge(t *testing.T) {
	graph := structures.NewAdjacencyMatrixGraph[string, int](false)
	_ = graph.AddVertex("A")
	_ = graph.AddVertex("B")
	_ = graph.AddEdge("A", "B", 5)

	err := graph.RemoveEdge("A", "B")
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	_, err = graph.GetWeight("A", "B")
	if !errors.Is(err, structures.ErrEdgeNotFound) {
		t.Fatalf("expected structures.ErrEdgeNotFound, got %v", err)
	}
}

func TestAdjacencyMatrixGraph_HasEdge(t *testing.T) {
	graph := structures.NewAdjacencyMatrixGraph[string, int](false)
	_ = graph.AddVertex("A")
	_ = graph.AddVertex("B")
	_ = graph.AddEdge("A", "B", 5)

	if !graph.HasEdge("A", "B") {
		t.Fatalf("expected edge to exist between A and B")
	}

	if graph.HasEdge("B", "A") {
		t.Fatalf("expected no edge between B and A")
	}
}

func TestAdjacencyMatrixGraph_Neighbors(t *testing.T) {
	graph := structures.NewAdjacencyMatrixGraph[string, int](false)
	_ = graph.AddVertex("A")
	_ = graph.AddVertex("B")
	_ = graph.AddVertex("C")
	_ = graph.AddEdge("A", "B", 5)
	_ = graph.AddEdge("A", "C", 10)

	neighbors, err := graph.Neighbors("A")
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if len(neighbors) != 2 || neighbors["B"] != 5 || neighbors["C"] != 10 {
		t.Fatalf("expected neighbors B(5) and C(10), got %v", neighbors)
	}
}

func TestAdjacencyMatrixGraph_Degree(t *testing.T) {
	graph := structures.NewAdjacencyMatrixGraph[string, int](false)
	_ = graph.AddVertex("A")
	_ = graph.AddVertex("B")
	_ = graph.AddVertex("C")
	_ = graph.AddEdge("A", "B", 5)
	_ = graph.AddEdge("A", "C", 10)

	degree, err := graph.Degree("A")
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if degree != 2 {
		t.Fatalf("expected degree 2 for vertex A, got %d", degree)
	}
}

func TestAdjacencyMatrixGraph_InDegree(t *testing.T) {
	graph := structures.NewAdjacencyMatrixGraph[string, int](false)
	_ = graph.AddVertex("A")
	_ = graph.AddVertex("B")
	_ = graph.AddVertex("C")
	_ = graph.AddEdge("B", "A", 5)
	_ = graph.AddEdge("C", "A", 10)

	inDegree, err := graph.InDegree("A")
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if inDegree != 2 {
		t.Fatalf("expected in-degree 2 for vertex A, got %d", inDegree)
	}
}

func TestAdjacencyMatrixGraph_Transpose(t *testing.T) {
	graph := structures.NewAdjacencyMatrixGraph[string, int](true)
	_ = graph.AddVertex("A")
	_ = graph.AddVertex("B")
	_ = graph.AddEdge("A", "B", 5)

	transposed := graph.Transpose()

	if !transposed.HasEdge("B", "A") || transposed.HasEdge("A", "B") {
		t.Fatalf("expected edge from B to A in transposed graph, got incorrect result")
	}
}

func TestAdjacencyMatrixGraph_ShortestPath(t *testing.T) {
	graph := structures.NewAdjacencyMatrixGraph[string, int](false)
	_ = graph.AddVertex("A")
	_ = graph.AddVertex("B")
	_ = graph.AddVertex("C")
	_ = graph.AddEdge("A", "B", 5)
	_ = graph.AddEdge("B", "C", 10)

	path, err := graph.ShortestPath("A", "C")
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	expectedPath := []string{"A", "B", "C"}
	if len(path) != len(expectedPath) {
		t.Fatalf("expected path length %d, got %d", len(expectedPath), len(path))
	}
	for i := range path {
		if path[i] != expectedPath[i] {
			t.Fatalf("expected %v, got %v at index %d", expectedPath[i], path[i], i)
		}
	}
}
