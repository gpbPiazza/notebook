package graphs

import (
	"fmt"

	"github.com/gpbPiazza/notebook/algo_data/queues"
)

// in graphs vertexes are nodes and
// edges are the link between two verticies
// we say that a vertex is an adjacent of ther vertex when
// they are linked each other directly by edges.
// Also refer to adjacent vertices as neighbors.

// cVertex is vertex from a full connected grapth implementation code
type cVertex struct {
	value     string
	adjacents []*cVertex
}

func NewCVertex(value string) *cVertex {
	return &cVertex{
		value: value,
	}
}

func (cv *cVertex) addAdjancent(vertex *cVertex) {
	if cv.includes(vertex) {
		return
	}
	cv.adjacents = append(cv.adjacents, vertex)

	vertex.addAdjancent(cv)
}

func (cv *cVertex) includes(vertex *cVertex) bool {
	for _, neighbor := range cv.adjacents {
		if neighbor == vertex {
			return true
		}
	}

	return false
}

// depth frist search transverse complete connected grath
func dfs_transverse(vertex *cVertex, visited map[*cVertex]bool) {
	visited[vertex] = true

	fmt.Println(vertex.value)

	for _, neighbor := range vertex.adjacents {
		if visited[neighbor] {
			continue
		}
		dfs_transverse(neighbor, visited)
	}
}

// breth first search transverse to a complete connected grath
func bfs_transverse(vertex *cVertex) {
	visited := make(map[*cVertex]bool)

	currentVertex := vertex
	queue := queues.NewSliceQueue[*cVertex]()
	queue.Enqueue(currentVertex)
	visited[currentVertex] = true

	for queue.Read() != nil {
		currentVertex := queue.Dequeue()
		fmt.Println(currentVertex.value)
		for _, neighbor := range currentVertex.adjacents {
			if !visited[neighbor] {
				queue.Enqueue(neighbor)
				visited[neighbor] = true
			}
		}
	}
}

// Efficiency of a graph search
// O(V + 2E) that V is the number os vertices of the graph and E the
// number of edges in the graph. The number two in E is because when we
// search in a grath we use two iterations per edge to reach the vertices,
// one to check visited and one to visit. Such Big O drop constants we have
// O(V+E).

// bfs VS dfs
// depends in your use case, if you are searching for a vertex close to your current  vertex
// bfs will perform better, if yuou are seraching fora vertex farther from you current vertex
// dfs will perform better. Because dfs will reach the most farther vertecies first and
// bfs will reach the closests ones first.
