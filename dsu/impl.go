package dsu

func makeSet(parent, rank []int, x int) {
	parent[x] = x
	rank[x] = 1
}

func find(parent []int, x int) int {
	setId := parent[x]
	if setId != x {
		parent[x] = find(parent, setId)
		return parent[x]
	}

	return setId
}

func union(parent, rank []int, x, y int) {
	xset := find(parent, x)
	yset := find(parent, y)
	if xset == yset {
		return
	}

	if rank[xset] < rank[yset] {
		xset, yset = yset, xset
	}

	parent[yset] = xset
	if rank[xset] == rank[yset] {
		rank[xset]++
	}
}
