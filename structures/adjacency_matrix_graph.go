package structures

import (
	"fmt"
)

// adjacencyMatrixGraph represents a weighted graph implemented using an adjacency matrix.
type adjacencyMatrixGraph[T comparable, W Numeric] struct {
	vertices []T
	matrix   [][]*W
	index    map[T]int
	directed bool
}

// NewAdjacencyMatrixGraph creates a new weighted adjacency matrix graph.
func NewAdjacencyMatrixGraph[T comparable, W Numeric](directed bool) Graph[T, W] {
	return &adjacencyMatrixGraph[T, W]{
		vertices: []T{},
		matrix:   [][]*W{},
		index:    make(map[T]int),
		directed: directed,
	}
}

// AddVertex adds a vertex to the graph.
func (g *adjacencyMatrixGraph[T, W]) AddVertex(vertex T) error {
	if _, exists := g.index[vertex]; exists {
		return fmt.Errorf("%w: %v", ErrVertexAlreadyExists, vertex)
	}
	g.index[vertex] = len(g.vertices)
	g.vertices = append(g.vertices, vertex)
	for i := range g.matrix {
		g.matrix[i] = append(g.matrix[i], nil)
	}
	g.matrix = append(g.matrix, make([]*W, len(g.vertices)))
	return nil
}

// RemoveVertex removes a vertex and all its edges.
func (g *adjacencyMatrixGraph[T, W]) RemoveVertex(vertex T) error {
	idx, exists := g.index[vertex]
	if !exists {
		return ErrVertexNotFound
	}
	// Remove vertex from matrix
	g.matrix = append(g.matrix[:idx], g.matrix[idx+1:]...)
	for i := range g.matrix {
		g.matrix[i] = append(g.matrix[i][:idx], g.matrix[i][idx+1:]...)
	}
	// Remove vertex from list of vertices and update indices
	g.vertices = append(g.vertices[:idx], g.vertices[idx+1:]...)
	delete(g.index, vertex)
	for i := idx; i < len(g.vertices); i++ {
		g.index[g.vertices[i]] = i
	}
	return nil
}

// AddEdge adds a weighted edge between two vertices.
func (g *adjacencyMatrixGraph[T, W]) AddEdge(from, to T, weight W) error {
	fromIdx, fromExists := g.index[from]
	toIdx, toExists := g.index[to]
	if !fromExists || !toExists {
		return ErrVertexNotFound
	}
	g.matrix[fromIdx][toIdx] = &weight
	return nil
}

// RemoveEdge removes the edge between two vertices.
func (g *adjacencyMatrixGraph[T, W]) RemoveEdge(from, to T) error {
	fromIdx, fromExists := g.index[from]
	toIdx, toExists := g.index[to]
	if !fromExists || !toExists {
		return ErrVertexNotFound
	}
	g.matrix[fromIdx][toIdx] = nil
	return nil
}

// HasEdge checks if there is an edge between two vertices.
func (g *adjacencyMatrixGraph[T, W]) HasEdge(from, to T) bool {
	fromIdx, fromExists := g.index[from]
	toIdx, toExists := g.index[to]
	if !fromExists || !toExists {
		return false
	}
	return g.matrix[fromIdx][toIdx] != nil
}

// Neighbors returns the neighbors of a vertex with their weights.
func (g *adjacencyMatrixGraph[T, W]) Neighbors(vertex T) (map[T]W, error) {
	idx, exists := g.index[vertex]
	if !exists {
		return nil, ErrVertexNotFound
	}
	neighbors := make(map[T]W)
	for i, weight := range g.matrix[idx] {
		if weight != nil {
			neighbors[g.vertices[i]] = *weight
		}
	}
	return neighbors, nil
}

// Weight returns the weight of the edge between two vertices.
func (g *adjacencyMatrixGraph[T, W]) Weight(from, to T) (W, error) {
	fromIdx, fromExists := g.index[from]
	toIdx, toExists := g.index[to]
	if !fromExists || !toExists {
		return *new(W), ErrVertexNotFound
	}
	weight := g.matrix[fromIdx][toIdx]
	if weight == nil {
		return *new(W), ErrEdgeNotFound
	}
	return *weight, nil
}

// Vertices returns all vertices in the graph.
func (g *adjacencyMatrixGraph[T, W]) Vertices() []T {
	return g.vertices
}

// Edges returns all edges in the graph with their weights.
func (g *adjacencyMatrixGraph[T, W]) Edges() []Edge[T, W] {
	var edges []Edge[T, W]
	for i, from := range g.vertices {
		for j, to := range g.vertices {
			if g.matrix[i][j] != nil {
				edges = append(edges, Edge[T, W]{From: from, To: to, Weight: *g.matrix[i][j]})
			}
		}
	}
	return edges
}

// Degree returns the out-degree of a vertex.
func (g *adjacencyMatrixGraph[T, W]) Degree(vertex T) (int, error) {
	idx, exists := g.index[vertex]
	if !exists {
		return 0, ErrVertexNotFound
	}
	outDegree := 0
	for _, weight := range g.matrix[idx] {
		if weight != nil {
			outDegree++
		}
	}
	return outDegree, nil
}

// InDegree returns the in-degree of a vertex.
func (g *adjacencyMatrixGraph[T, W]) InDegree(vertex T) (int, error) {
	idx, exists := g.index[vertex]
	if !exists {
		return 0, ErrVertexNotFound
	}
	inDegree := 0
	for i := range g.vertices {
		if g.matrix[i][idx] != nil {
			inDegree++
		}
	}
	return inDegree, nil
}

// Transpose returns the transposed graph (reverses all edges).
func (g *adjacencyMatrixGraph[T, W]) Transpose() Graph[T, W] {
	transposed := NewAdjacencyMatrixGraph[T, W](g.directed)
	for _, vertex := range g.vertices {
		transposed.AddVertex(vertex)
	}
	for i, from := range g.vertices {
		for j, to := range g.vertices {
			if g.matrix[i][j] != nil {
				transposed.AddEdge(to, from, *g.matrix[i][j])
			}
		}
	}
	return transposed
}

// IsDirected returns whether the graph is directed or not.
func (g *adjacencyMatrixGraph[T, W]) IsDirected() bool {
	return g.directed
}

// ShortestPath implements Dijkstra's algorithm to find the shortest path.
func (g *adjacencyMatrixGraph[T, W]) ShortestPath(from, to T) ([]T, error) {
	_, fromExists := g.index[from]
	_, toExists := g.index[to]
	if !fromExists || !toExists {
		return nil, ErrVertexNotFound
	}

	// Initialize distances, previous vertices, and the priority queue.
	distances := make(map[T]W)
	previous := make(map[T]*T)
	pq := NewPriorityQueue[T](func(a, b T) int {
		// Compare distances of a and b for priority queue ordering.
		if distances[a] < distances[b] {
			return -1
		}
		if distances[a] > distances[b] {
			return 1
		}
		return 0
	})

	infinity := NumericMaxValue[W]()
	zero := NumericZeroValue[W]()

	// Set initial distances to infinity.
	for _, vertex := range g.vertices {
		distances[vertex] = infinity
	}
	distances[from] = zero

	// Start with the source vertex in the priority queue.
	pq.Push(from)

	// Dijkstra's algorithm main loop.
	for pq.Size() > 0 {
		current, _ := pq.Pop()
		currentIdx := g.index[current]

		// Early exit if we reached the target vertex.
		if current == to {
			break
		}

		// Iterate over neighbors of the current vertex.
		for neighborIdx, weightPtr := range g.matrix[currentIdx] {
			if weightPtr == nil {
				continue // Skip if there is no edge.
			}

			neighbor := g.vertices[neighborIdx]
			weight := *weightPtr
			newDist := distances[current] + weight

			// Update the distance if a shorter path is found.
			if newDist < distances[neighbor] {
				distances[neighbor] = newDist
				previous[neighbor] = &current
				pq.Push(neighbor)
			}
		}
	}

	// If the destination vertex distance is still infinity, no path exists.
	if distances[to] == infinity {
		return nil, fmt.Errorf("%w: no path from %v to %v", ErrFindingShortestPath, from, to)
	}

	// Reconstruct the path from `to` to `from` using the `previous` map.
	var path []T
	for at := &to; at != nil; at = previous[*at] {
		path = append([]T{*at}, path...)
	}

	return path, nil
}
