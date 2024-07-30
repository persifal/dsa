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
	if n == nil {
		return nil
	}

	if n.val == v1 || n.val == v2 {
		return n
	}

	left := lca(n.left, v1, v2)
	right := lca(n.right, v1, v2)
	if left != nil && right != nil {
		return n
	} else {
		if left != nil {
			return left
		} else {
			return right
		}
	}
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
