package structures

import "fmt"

// adjacencyListGraph represents a weighted graph implemented using an adjacency list.
type adjacencyListGraph[T comparable, W any] struct {
	adjList  map[T]map[T]W
	directed bool
}

// NewAdjacencyListGraph creates a new weighted adjacency list graph.
func NewAdjacencyListGraph[T comparable, W any](directed bool) Graph[T, W] {
	return &adjacencyListGraph[T, W]{
		adjList:  make(map[T]map[T]W),
		directed: directed,
	}
}

// AddVertex adds a vertex to the graph.
func (g *adjacencyListGraph[T, W]) AddVertex(vertex T) error {
	if _, exists := g.adjList[vertex]; exists {
		return fmt.Errorf("vertex %v already exists", vertex)
	}
	g.adjList[vertex] = make(map[T]W)
	return nil
}

// RemoveVertex removes a vertex and all its edges.
func (g *adjacencyListGraph[T, W]) RemoveVertex(vertex T) error {
	if _, exists := g.adjList[vertex]; !exists {
		return ErrVertexNotFound
	}
	delete(g.adjList, vertex)
	for _, neighbors := range g.adjList {
		delete(neighbors, vertex)
	}
	return nil
}

// AddEdge adds a weighted edge between two vertices.
func (g *adjacencyListGraph[T, W]) AddEdge(from, to T, weight W) error {
	if _, exists := g.adjList[from]; !exists {
		return ErrVertexNotFound
	}
	if _, exists := g.adjList[to]; !exists {
		return ErrVertexNotFound
	}
	g.adjList[from][to] = weight
	return nil
}

// RemoveEdge removes the edge between two vertices.
func (g *adjacencyListGraph[T, W]) RemoveEdge(from, to T) error {
	if _, exists := g.adjList[from]; !exists {
		return ErrVertexNotFound
	}
	if _, exists := g.adjList[from][to]; !exists {
		return ErrEdgeNotFound
	}
	delete(g.adjList[from], to)
	return nil
}

// HasEdge checks if there is an edge between two vertices.
func (g *adjacencyListGraph[T, W]) HasEdge(from, to T) bool {
	_, exists := g.adjList[from][to]
	return exists
}

// Neighbors returns the neighbors of a vertex with their weights.
func (g *adjacencyListGraph[T, W]) Neighbors(vertex T) (map[T]W, error) {
	neighbors, exists := g.adjList[vertex]
	if !exists {
		return nil, ErrVertexNotFound
	}
	return neighbors, nil
}

// GetWeight returns the weight of the edge between two vertices.
func (g *adjacencyListGraph[T, W]) GetWeight(from, to T) (W, error) {
	if _, exists := g.adjList[from]; !exists {
		return *new(W), ErrVertexNotFound
	}
	weight, exists := g.adjList[from][to]
	if !exists {
		return *new(W), ErrEdgeNotFound
	}
	return weight, nil
}

// GetVertices returns all vertices in the graph.
func (g *adjacencyListGraph[T, W]) GetVertices() []T {
	vertices := make([]T, 0, len(g.adjList))
	for vertex := range g.adjList {
		vertices = append(vertices, vertex)
	}
	return vertices
}

// GetEdges returns all edges in the graph with their weights.
func (g *adjacencyListGraph[T, W]) GetEdges() []Edge[T, W] {
	var edges []Edge[T, W]
	for from, neighbors := range g.adjList {
		for to, weight := range neighbors {
			edges = append(edges, Edge[T, W]{From: from, To: to, Weight: weight})
		}
	}
	return edges
}

// Degree returns the out-degree of a vertex.
func (g *adjacencyListGraph[T, W]) Degree(vertex T) (int, error) {
	neighbors, exists := g.adjList[vertex]
	if !exists {
		return 0, ErrVertexNotFound
	}
	return len(neighbors), nil
}

// InDegree returns the in-degree of a vertex.
func (g *adjacencyListGraph[T, W]) InDegree(vertex T) (int, error) {
	if _, exists := g.adjList[vertex]; !exists {
		return 0, ErrVertexNotFound
	}
	inDegree := 0
	for _, neighbors := range g.adjList {
		if _, exists := neighbors[vertex]; exists {
			inDegree++
		}
	}
	return inDegree, nil
}

// Transpose returns the transposed graph (reverses all edges).
func (g *adjacencyListGraph[T, W]) Transpose() Graph[T, W] {
	transposed := NewAdjacencyListGraph[T, W](g.directed)
	for from, neighbors := range g.adjList {
		for to, weight := range neighbors {
			transposed.AddVertex(from)
			transposed.AddVertex(to)
			transposed.AddEdge(to, from, weight)
		}
	}
	return transposed
}

// IsDirected returns whether the graph is directed or not.
func (g *adjacencyListGraph[T, W]) IsDirected() bool {
	return g.directed
}

// ShortestPath implements Dijkstra's algorithm to find the shortest path.
func (g *adjacencyListGraph[T, W]) ShortestPath(from, to T) ([]T, error) {
	// Dijkstra's algorithm (simplified)
	// Requires weights to be comparable, assume W is numeric
	return nil, fmt.Errorf("not implemented")
}
