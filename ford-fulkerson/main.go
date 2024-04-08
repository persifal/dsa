package main

import (
	"log"
)

func main() {
	edges := map[int][]*edge{

		// 0: {{0, 3, 16}, {0, 4, 13}},
		// 2: {{2, 4, 9}, {2, 1, 20}},
		// 3: {{3, 2, 12}},
		// 4: {{4, 3, 4}, {4, 5, 14}},
		// 5: {{5, 2, 7}, {5, 1, 4}},

		0: {{0, 2, 1000000}, {0, 3, 1000000}},
		2: {{2, 3, 1}, {2, 1, 1000000}},
		3: {{3, 1, 1000000}},
	}

	res := maxFlow(edges, 0, 1)
	log.Printf("Max Flow: %d\n", res)
}

type (
	graph = map[int][]*edge
	edge  struct {
		from, to, weight int
	}
)

func maxFlow(g graph, s, t int) int {
	f := 0
	c := 0
	path := findPath(g, s, t)
	for ; path != nil; path = findPath(g, s, t) {
		c++
		mw := minWeight(path)
		update(g, path, mw)
		f += mw
	}
	log.Printf("Iterations: %d\n", c)

	return f
}

func update(g graph, path []*edge, mw int) {
	for i := range path {
		e := path[i]
		w := e.weight
		e.weight = w - mw
		if e.weight == 0 {
			removeEdge(g, e)
		}

		backEdge(g, e.to, e.from, mw)
	}
}

func backEdge(g graph, from, to, weight int) {
	vals := g[from]
	for i := range vals {
		v := vals[i]
		if v.to == to {
			v.weight += weight
			return
		}
	}

	// no edge found. add
	g[from] = append(g[from], &edge{from, to, weight})
}

func removeEdge(g graph, e *edge) {
	vals := g[e.from]
	if len(vals) <= 1 {
		delete(g, e.from)
	} else {
		for i, v := range vals {
			if v.to == e.to {
				// alloc
				g[e.from] = append(vals[:i], vals[i+1:]...)
				break
			}
		}
	}
}

func minWeight(path []*edge) int {
	res := 1 << 31
	for _, e := range path {
		if e.weight < res {
			res = e.weight
		}
	}

	return res
}

func findPath(g graph, s, t int) []*edge {
	visited := make(map[int]bool)
	path := []*edge{}
	if dfs(g, visited, &path, s, t) {
		return path
	}

	return nil
}

func dfs(g graph, visited map[int]bool, path *[]*edge, i, t int) bool {
	visited[i] = true
	edges := g[i]
	for ie := range edges {
		edge := edges[ie]
		*path = append(*path, edge)
		if edge.to == t {
			return true
		}

		if !visited[edge.to] {
			if dfs(g, visited, path, edge.to, t) {
				return true
			}
		}

		*path = (*path)[:len(*path)-1]
	}

	return false
}
