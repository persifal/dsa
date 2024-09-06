package quickselect

import (
	"testing"
)

func Test(t *testing.T) {
	arr := []int{5, 2, 8, 3, 11}
	kth := quickSelect(arr, 3)
	if kth != 8 {
		t.Errorf("Quick Select wrong result. Expected: 3, Given: %d\n", kth)
	}
}
