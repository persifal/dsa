package ext2

import (
	"fmt"
)

// поиск подотрезка с максимальной суммой

func main() {

}

type node struct {
	sum, pref, suff, value int
}

func (n *node) String() string {
	return fmt.Sprintf("[sum: %d, pref: %d, suff: %d, value: %d]", n.sum, n.pref, n.suff, n.value)
}

func newNode(val int) *node {
	return &node{
		val,
		max(0, val),
		max(0, val),
		max(0, val),
	}
}

func build(arr []int) []*node {
	n := len(arr)
	tree := make([]*node, 4*n)

	doBuild(arr, tree, 1, 0, n-1)

	return tree
}

func doBuild(arr []int, tree []*node, currentNodeTree, treeLeft, treeRight int) {
	if treeLeft == treeRight {
		tree[currentNodeTree] = newNode(arr[treeLeft])
	} else {
		mid := (treeLeft + treeRight) / 2
		doBuild(arr, tree, currentNodeTree*2, treeLeft, mid)
		doBuild(arr, tree, currentNodeTree*2+1, mid+1, treeRight)

		tree[currentNodeTree] = combine(tree[currentNodeTree*2], tree[currentNodeTree*2+1])
	}
}

func combine(l, r *node) *node {
	return &node{
		l.sum + r.sum,
		max(l.pref, l.sum+r.pref),
		max(r.suff, r.sum+l.suff),
		max(l.suff+r.pref, max(l.value, r.value)),
	}
}

func query(tree []*node, currentNodeTree, treeLeft, treeRight, left, right int) *node {
	if treeLeft == left && treeRight == right {
		return tree[currentNodeTree]
	}

	mid := (treeLeft + treeRight) / 2
	if right <= mid {
		return query(tree, currentNodeTree*2, treeLeft, mid, left, right)
	}

	if left > mid {
		return query(tree, currentNodeTree*2+1, mid+1, treeRight, left, right)
	}

	return combine(
		query(tree, currentNodeTree*2, treeLeft, mid, left, mid),
		query(tree, currentNodeTree*2+1, mid+1, treeRight, mid+1, right),
	)
}

func update(tree []*node, currentNodeTree, treeLeft, treeRight, position, newValue int) {
	if treeLeft == treeRight {
		tree[currentNodeTree] = newNode(newValue)
	} else {
		mid := (treeLeft + treeRight) / 2
		if position <= mid {
			update(tree, currentNodeTree*2, treeLeft, mid, position, newValue)
		} else {
			update(tree, currentNodeTree*2+1, mid+1, treeRight, position, newValue)
		}

		tree[currentNodeTree] = combine(tree[currentNodeTree*2], tree[currentNodeTree*2+1])
	}
}
