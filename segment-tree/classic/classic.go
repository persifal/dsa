package classic

import (
	"fmt"
)

type tree struct {
	arr []int
	op  func(a, b int) int
}

func (t *tree) String() string {
	return fmt.Sprintf("%v", t.arr)
}

func build(arr []int, op func(a, b int) int) *tree {
	n := len(arr)
	tree := &tree{
		make([]int, 4*n),
		op,
	}

	// root has number 1, 0 and n-1 define source of data to be used
	doBuild(arr, tree, 1, 0, n-1)

	return tree
}

// Build Segment Tree for Sum operation
func doBuild(arr []int, tree *tree, currentNodeTree, left, right int) {
	if left == right {
		tree.arr[currentNodeTree] = arr[left]
	} else {
		mid := (left + right) / 2
		doBuild(arr, tree, currentNodeTree*2, left, mid)
		doBuild(arr, tree, currentNodeTree*2+1, mid+1, right)

		// sum operation
		tree.arr[currentNodeTree] = tree.op(tree.arr[currentNodeTree*2], tree.arr[currentNodeTree*2+1])
	}
}

func query(tree *tree, currentNodeTree, treeLeft, treeRight, left, right int) int {
	if left > right {
		return 0
	}

	// if segments match
	if treeLeft == left && treeRight == right {
		// take current node sum
		return tree.arr[currentNodeTree]
	}

	mid := (treeLeft + treeRight) / 2

	// cut query segments because of cut tree segments
	leftpart := query(tree, currentNodeTree*2, treeLeft, mid, left, min(mid, right))
	rightPart := query(tree, currentNodeTree*2+1, mid+1, treeRight, max(left, mid+1), right)

	return tree.op(leftpart, rightPart)
}

func update(tree *tree, currentNodeTree, treeLeft, treeRight, position, newValue int) {
	if treeLeft == treeRight {
		tree.arr[currentNodeTree] = newValue
	} else {
		mid := (treeLeft + treeRight) / 2
		if position <= mid {
			update(tree, currentNodeTree*2, treeLeft, mid, position, newValue)
		} else {
			update(tree, currentNodeTree*2+1, mid+1, treeRight, position, newValue)
		}

		tree.arr[currentNodeTree] = tree.op(tree.arr[currentNodeTree*2], tree.arr[currentNodeTree*2+1])
	}
}
