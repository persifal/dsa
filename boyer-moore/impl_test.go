package boyer_moore

import (
	"testing"
)

func Test(t *testing.T) {
	// given
	arr := []int{3, 3, 4, 2, 3, 3, 5}

	// when
	dominate := boyerMoore(arr)

	// then
	if dominate != 3 {
		t.Errorf("Error in finding dominate num in arr. Expected: 3, Found: %v", dominate)
	}
}
