package parent_pointer_depth

import "fmt"

type node struct {
	left, right, parent *node
	val                 int
}

func (n *node) String() string {
	return fmt.Sprintf("%d", n.val)
}

func lca(n *node, v1, v2 int) *node {
	return doLca(find(n, v1), find(n, v2))
}

func find(n *node, v int) *node {
	if n == nil {
		return nil
	}

	if n.val == v {
		return n
	}

	l := find(n.left, v)
	if l != nil {
		return l
	}

	r := find(n.right, v)
	if r != nil {
		return r
	}

	return nil
}

func doLca(n1, n2 *node) *node {
	d1 := depth(n1)
	d2 := depth(n2)
	diff := d1 - d2
	if diff < 0 {
		n1, n2 = n2, n1
		diff = -diff
	}

	for diff != 0 {
		n1 = n1.parent
		diff--
	}

	for n1 != nil && n2 != nil {
		if n1 == n2 {
			return n1
		}

		n1 = n1.parent
		n2 = n2.parent
	}

	return nil
}

func depth(n *node) int {
	d := -1
	curr := n
	for curr != nil {
		d++
		curr = curr.parent
	}

	return d
}

// construct binary tree from array
func build(arr []int, from, to int) *node {
	if from > to {
		return nil
	}

	mid := (from + to) / 2
	node := newNode(arr[mid])
	node.left = build(arr, from, mid-1)
	if node.left != nil {
		node.left.parent = node
	}

	node.right = build(arr, mid+1, to)
	if node.right != nil {
		node.right.parent = node
	}

	return node
}

func newNode(val int) *node {
	return &node{
		nil,
		nil,
		nil,
		val,
	}
}
