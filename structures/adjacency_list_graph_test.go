package structures_test

import (
	"testing"

	"github.com/Jibaru/golang-data-structures/structures"
)

func TestNewAdjacencyListGraph(t *testing.T) {
	graph := structures.NewAdjacencyListGraph[int, float64](false)
	if graph == nil {
		t.Fatal("structures.NewAdjacencyListGraph should not return nil")
	}
}

func TestAdjacencyListGraph_AddVertex(t *testing.T) {
	graph := structures.NewAdjacencyListGraph[int, float64](false)
	err := graph.AddVertex(1)
	if err != nil {
		t.Fatalf("AddVertex should not return an error: %v", err)
	}

	err = graph.AddVertex(1)
	if err == nil {
		t.Fatal("AddVertex should return an error when adding a duplicate vertex")
	}
}

func TestAdjacencyListGraph_AddEdge(t *testing.T) {
	graph := structures.NewAdjacencyListGraph[int, float64](false)
	graph.AddVertex(1)
	graph.AddVertex(2)

	err := graph.AddEdge(1, 2, 3.5)
	if err != nil {
		t.Fatalf("AddEdge should not return an error: %v", err)
	}

	err = graph.AddEdge(1, 3, 2.0)
	if err == nil {
		t.Fatal("AddEdge should return an error when vertices do not exist")
	}
}

func TestAdjacencyListGraph_RemoveVertex(t *testing.T) {
	g := structures.NewAdjacencyListGraph[string, int](true)
	g.AddVertex("A")
	g.AddVertex("B")
	g.AddEdge("A", "B", 10)

	// Test valid vertex removal.
	err := g.RemoveVertex("A")
	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}
	if len(g.Vertices()) != 1 {
		t.Errorf("expected 1 vertex, got %d", len(g.Vertices()))
	}

	// Test removal of non-existing vertex.
	err = g.RemoveVertex("C")
	if err == nil {
		t.Errorf("expected error for non-existing vertex, got nil")
	}
}

func TestAdjacencyListGraph_RemoveEdge(t *testing.T) {
	g := structures.NewAdjacencyListGraph[string, int](true)
	g.AddVertex("A")
	g.AddVertex("B")
	g.AddEdge("A", "B", 10)

	err := g.RemoveEdge("A", "B")
	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}
	if g.HasEdge("A", "B") {
		t.Errorf("expected edge to be removed")
	}

	err = g.RemoveEdge("A", "B")
	if err == nil {
		t.Errorf("expected error for non-existing edge, got nil")
	}
}

func TestAdjacencyListGraph_HasEdge(t *testing.T) {
	g := structures.NewAdjacencyListGraph[string, int](true)
	g.AddVertex("A")
	g.AddVertex("B")
	g.AddEdge("A", "B", 10)

	if !g.HasEdge("A", "B") {
		t.Errorf("expected edge to exist")
	}

	if g.HasEdge("B", "A") {
		t.Errorf("expected edge not to exist")
	}
}

func TestAdjacencyListGraph_Neighbors(t *testing.T) {
	g := structures.NewAdjacencyListGraph[string, int](true)
	g.AddVertex("A")
	g.AddVertex("B")
	g.AddEdge("A", "B", 10)

	neighbors, err := g.Neighbors("A")
	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}
	if len(neighbors) != 1 || neighbors["B"] != 10 {
		t.Errorf("expected neighbor B with weight 10, got %v", neighbors)
	}

	_, err = g.Neighbors("C")
	if err == nil {
		t.Errorf("expected error for non-existing vertex")
	}
}

func TestAdjacencyListGraph_Weight(t *testing.T) {
	g := structures.NewAdjacencyListGraph[string, int](true)
	g.AddVertex("A")
	g.AddVertex("B")
	g.AddEdge("A", "B", 10)

	weight, err := g.Weight("A", "B")
	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}
	if weight != 10 {
		t.Errorf("expected weight 10, got %d", weight)
	}

	_, err = g.Weight("A", "C")
	if err == nil {
		t.Errorf("expected error for non-existing edge")
	}
}

func TestAdjacencyListGraph_Vertices(t *testing.T) {
	g := structures.NewAdjacencyListGraph[string, int](true)
	g.AddVertex("A")
	g.AddVertex("B")

	vertices := g.Vertices()
	if len(vertices) != 2 {
		t.Errorf("expected 2 vertices, got %d", len(vertices))
	}
}

func TestAdjacencyListGraph_Edges(t *testing.T) {
	g := structures.NewAdjacencyListGraph[string, int](true)
	g.AddVertex("A")
	g.AddVertex("B")
	g.AddEdge("A", "B", 10)

	edges := g.Edges()
	if len(edges) != 1 || edges[0].From != "A" || edges[0].To != "B" || edges[0].Weight != 10 {
		t.Errorf("expected edge A -> B with weight 10, got %v", edges)
	}
}

func TestAdjacencyListGraph_Degree(t *testing.T) {
	g := structures.NewAdjacencyListGraph[string, int](true)
	g.AddVertex("A")
	g.AddVertex("B")
	g.AddEdge("A", "B", 10)

	degree, err := g.Degree("A")
	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}
	if degree != 1 {
		t.Errorf("expected degree 1, got %d", degree)
	}

	_, err = g.Degree("C")
	if err == nil {
		t.Errorf("expected error for non-existing vertex")
	}
}

func TestAdjacencyListGraph_InDegree(t *testing.T) {
	g := structures.NewAdjacencyListGraph[string, int](true)
	g.AddVertex("A")
	g.AddVertex("B")
	g.AddEdge("A", "B", 10)

	inDegree, err := g.InDegree("B")
	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}
	if inDegree != 1 {
		t.Errorf("expected in-degree 1, got %d", inDegree)
	}

	_, err = g.InDegree("C")
	if err == nil {
		t.Errorf("expected error for non-existing vertex")
	}
}

func TestAdjacencyListGraph_Transpose(t *testing.T) {
	g := structures.NewAdjacencyListGraph[string, int](true)
	g.AddVertex("A")
	g.AddVertex("B")
	g.AddEdge("A", "B", 10)

	transposed := g.Transpose()
	if !transposed.HasEdge("B", "A") {
		t.Errorf("expected transposed graph to have edge B -> A")
	}
	if transposed.HasEdge("A", "B") {
		t.Errorf("expected transposed graph to not have edge A -> B")
	}
}

func TestAdjacencyListGraph_IsDirected(t *testing.T) {
	g := structures.NewAdjacencyListGraph[string, int](true)
	if !g.IsDirected() {
		t.Errorf("expected graph to be directed")
	}

	undirected := structures.NewAdjacencyListGraph[string, int](false)
	if undirected.IsDirected() {
		t.Errorf("expected graph to be undirected")
	}
}

func TestAdjacencyListGraph_ShortestPath(t *testing.T) {
	g := structures.NewAdjacencyListGraph[string, int](true)
	g.AddVertex("A")
	g.AddVertex("B")
	g.AddVertex("C")
	g.AddEdge("A", "B", 1)
	g.AddEdge("B", "C", 2)

	path, err := g.ShortestPath("A", "C")
	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}
	expectedPath := []string{"A", "B", "C"}
	for i := range path {
		if path[i] != expectedPath[i] {
			t.Errorf("expected %v, got %v", expectedPath, path)
		}
	}

	_, err = g.ShortestPath("A", "D")
	if err == nil {
		t.Errorf("expected error for no path")
	}
}
