package ext1

import (
	"math"
)

// поиск минимума и количества раз, которое он встречается

func build(arr []int) [][]int {
	n := len(arr)
	tree := make([][]int, 4*n)

	doBuild(arr, tree, 1, 0, n-1)

	return tree
}

// Build Segment Tree for min operation
func doBuild(arr []int, tree [][]int, currentNodeTree, left, right int) {
	if left == right {
		tree[currentNodeTree] = []int{arr[left], 1}
	} else {
		mid := (left + right) / 2
		doBuild(arr, tree, currentNodeTree*2, left, mid)
		doBuild(arr, tree, currentNodeTree*2+1, mid+1, right)

		// min operation
		tree[currentNodeTree] = combine(tree[currentNodeTree*2], tree[currentNodeTree*2+1])
	}
}

func combine(a, b []int) []int {
	if a[0] > b[0] {
		return b
	}

	if a[0] < b[0] {
		return a
	}

	return []int{a[0], a[1] + b[1]}
}

func query(tree [][]int, currentNodeTree, treeLeft, treeRight, left, right int) []int {
	if left > right {
		return []int{math.MaxInt32, 0}
	}

	// if segments match
	if treeLeft == left && treeRight == right {
		// take current node sum
		return tree[currentNodeTree]
	}

	mid := (treeLeft + treeRight) / 2

	// cut query segments because of cut tree segments
	leftpart := query(tree, currentNodeTree*2, treeLeft, mid, left, min(mid, right))
	rightPart := query(tree, currentNodeTree*2+1, mid+1, treeRight, max(left, mid+1), right)

	return combine(leftpart, rightPart)
}

func update(tree [][]int, currentNodeTree, treeLeft, treeRight, position, newValue int) {
	if treeLeft == treeRight {
		tree[currentNodeTree] = []int{newValue, 1}
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
