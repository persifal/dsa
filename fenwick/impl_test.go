package fenwick

import "testing"

func Test(t *testing.T) {
	arr := []int{2, 1, 1, 3, 2, 3, 4, 5, 6, 7, 8, 9}
	tree := fenwick(arr)
	result := sum(tree, 6)
	if result != 12 {
		t.Errorf("Error. Expected: 12, Given: %d\n", result)
	}

	update(tree, 3, 21)
	result = sum(tree, 6)
	if result != 33 {
		t.Errorf("Error. Expected: 33, Given: %d\n", result)
	}
}
