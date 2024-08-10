package dsu

import (
	"testing"
)

func Test(t *testing.T) {
	parent := make([]int, 9)
	rank := make([]int, 9)
	for i := 0; i < len(parent); i++ {
		makeSet(parent, rank, i)
	}

	union(parent, rank, 3, 4)

	x := find(parent, 4)
	y := find(parent, 3)
	if x != y {
		t.Errorf("Different set leaders found 3: %d, 4: %d", y, x)
	}
}
