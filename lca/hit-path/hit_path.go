package single_traverse

import "fmt"

type node struct {
	left, right *node
	val         int
}

func (n *node) String() string {
	return fmt.Sprintf("%d", n.val)
}

func lca(n *node, v1, v2 int) *node {
	parentMap := buildParentMap(n)
	var node1 *node
	var node2 *node
	queue := make([]*node, 0)
	queue = append(queue, n)
	for len(queue) != 0 {
		elem := queue[0]
		queue = queue[1:]
		if elem.val == v1 {
			node1 = elem
		}

		if elem.val == v2 {
			node2 = elem
		}

		if elem.left != nil {
			queue = append(queue, elem.left)
		}

		if elem.right != nil {
			queue = append(queue, elem.right)
		}
	}

	path1 := make(map[*node]bool)
	for node1 != nil {
		path1[node1] = true
		node1 = parentMap[node1]
	}

	for node2 != nil {
		if path1[node2] {
			return node2
		}
		node2 = parentMap[node2]
	}

	return nil
}

func buildParentMap(n *node) map[*node]*node {
	m := make(map[*node]*node)
	m[n] = nil
	queue := make([]*node, 0)
	queue = append(queue, n)
	for len(queue) != 0 {
		elem := queue[0]
		queue = queue[1:]
		if elem.left != nil {
			m[elem.left] = elem
			queue = append(queue, elem.left)
		}
		if elem.right != nil {
			m[elem.right] = elem
			queue = append(queue, elem.right)
		}
	}

	return m
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

func newNode(val int) *node {
	return &node{
		nil,
		nil,
		val,
	}
}
