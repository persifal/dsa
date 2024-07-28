package naive

import (
	"testing"
)

func Test(t *testing.T) {
	arr := []int{3, 2, 5, 4, 11, 15, 16, 17, 18, 19, 20}
	tree := build(arr, 0, len(arr)-1)
	lca := lca(tree, 11, 19)
	if lca.val != 15 {
		t.Errorf("Incorrect LCA. Expected: 15, Given: %d", lca.val)
	}
}
