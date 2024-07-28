package classic

import (
	"testing"
)

func TestMax(t *testing.T) {
	arr := []int{3, 2, 5, 4, 11}
	n := len(arr)
	tree := build(arr, func(a, b int) int {
		return max(a, b)
	})
	result := query(tree, 1, 0, n-1, 1, 3)
	if result != 5 {
		t.Errorf("Wrong answer of Query max op. Expected: 5, Given: %d\n", result)
	}

	update(tree, 1, 0, n-1, 3, 12)
	result = query(tree, 1, 0, n-1, 3, 3)
	if result != 12 {
		t.Errorf("Wrong answer of Update max op. Expected: 12, Given: %d\n", result)
	}
}

func TestSum(t *testing.T) {
	arr := []int{3, 2, 5, 4, 11}
	n := len(arr)
	tree := build(arr, func(a, b int) int {
		return a + b
	})
	result := query(tree, 1, 0, n-1, 1, 3)
	if result != 11 {
		t.Errorf("Wrong answer of Query max op. Expected: 11, Given: %d\n", result)
	}

	update(tree, 1, 0, n-1, 4, 0)
	result = query(tree, 1, 0, n-1, 1, 4)
	if result != 11 {
		t.Errorf("Wrong answer of Update max op. Expected: 11, Given: %d\n", result)
	}
}
