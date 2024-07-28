package naive

import "fmt"

type node struct {
	left, right *node
	val         int
}

func (n *node) String() string {
	return fmt.Sprintf("%d", n.val)
}

func newNode(val int) *node {
	return &node{
		nil,
		nil,
		val,
	}
}

func lca(tree *node, v1, v2 int) *node {
	path1 := []*node{}
	dfs(tree, &path1, v1)
	path2 := []*node{}
	dfs(tree, &path2, v2)
	i := 0
	for path1[i].val == path2[i].val {
		i++
	}

	if i <= 0 {
		return nil
	}

	return path1[i-1]
}

func dfs(n *node, path *[]*node, v int) bool {
	if n == nil {
		return false
	}

	if n.val == v {
		return true
	}

	*path = append(*path, n)
	if dfs(n.left, path, v) || dfs(n.right, path, v) {
		return true
	}

	*path = (*path)[:len(*path)-1]

	return false
}

// construct binary tree from array
func build(arr []int, from, to int) *node {
	if from > to {
		return nil
	}

	mid := (from + to) / 2
	node := newNode(arr[mid])
	node.left = build(arr, from, mid-1)
	node.right = build(arr, mid+1, to)

	return node
}
