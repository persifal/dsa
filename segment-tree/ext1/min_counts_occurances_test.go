package ext1

import (
	"testing"
)

func TestMinAndCountOccurancies(t *testing.T) {
	arr := []int{3, 2, 5, 4, 2, 11}
	n := len(arr)
	tree := build(arr)
	result := query(tree, 1, 0, n-1, 1, 4)
	if result[0] != 2 && result[1] != 2 {
		t.Errorf("Expected: [2 2], Found: %v", result)
	}

	update(tree, 1, 0, n-1, 3, 2)
	result = query(tree, 1, 0, n-1, 2, 4)
	if result[0] != 2 && result[1] != 3 {
		t.Errorf("Expected: [2 3], Found: %v", result)
	}
}
