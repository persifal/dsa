package ext2

import (
	"testing"
)

func TestMaxSumSubArray(t *testing.T) {
	arr := []int{3, 2, -1, 5, -3, 6, 0, 1, -5, 4, 9, -2, -1, 2, 11}
	tree := build(arr)

	n := len(arr)
	result := query(tree, 1, 0, n-1, 1, 5)
	if result.value != 9 {
		t.Errorf("Query wrong result. Expected: 9, Given: %d", result.value)
	}

	update(tree, 1, 0, n-1, 3, 1)
	result = query(tree, 1, 0, n-1, 1, 5)
	if result.value != 6 {
		t.Errorf("Query wrong result. Expected: 6, Given: %d", result.value)
	}
}
