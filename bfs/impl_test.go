package bfs

import (
	"slices"
	"testing"
)

func Test(t *testing.T) {
	graph := build()
	distances := bfs(graph, 0)
	expected := []int{0, 1, 1, 2, 3}
	if slices.Compare(distances, expected) != 0 {
		t.Errorf("Error in finding distances. Expected: %v, Found: %v", expected, distances)
	}
}

func build() graph {
	return graph{
		0: []*edge{
			newEdge(1, 4),
			newEdge(2, 1),
		},
		1: []*edge{
			newEdge(3, 1),
		},
		2: []*edge{
			newEdge(1, 2),
			newEdge(3, 5),
		},
		3: []*edge{
			newEdge(4, 3),
		},
		4: []*edge{},
	}
}
