package fenwick

func fenwick(arr []int) []int {
	n := len(arr)
	tree := make([]int, n+1)
	for i := 0; i < n; i++ {
		update(tree, i+1, arr[i])
	}

	return tree
}

func update(tree []int, i, v int) {
	for i <= len(tree) {
		tree[i] += v
		i += i & (-i)
	}
}

// bound from 1 to i
func sum(tree []int, i int) int {
	sum := 0
	for i > 0 {
		sum += tree[i]
		i -= i & (-i)
	}

	return sum
}
