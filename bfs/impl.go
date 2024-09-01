package bfs

type graph = [][]*edge

type edge struct {
	v      int
	weight int
}

func bfs(g graph, source int) []int {
	graphSize := len(g)
	visited := make([]bool, graphSize)
	distances := make([]int, graphSize)
	distances[source] = 0
	queue := []int{source}
	visited[source] = true
	for len(queue) != 0 {
		u := queue[len(queue)-1]
		queue = queue[:len(queue)-1]
		for _, edge := range g[u] {
			v := edge.v
			if !visited[v] {
				queue = append(queue, v)
				distances[v] = distances[u] + 1
				visited[v] = true
			}
		}
	}

	return distances
}

func newEdge(v, weight int) *edge {
	return &edge{
		v, weight,
	}
}
