package main

import "testing"

func Test(t *testing.T) {
	needle := "sad"
	n := len(needle)
	heystack := "sadbutsad"
	zf := zFunc(needle + "#" + heystack)
	for idx, x := range zf {
		if x == n {
			s := idx - n - 1
			e := idx + x - n - 1
			found := heystack[s:e]
			if found != needle {
				t.Errorf("incorrect found: %s", found)
			}
		}
	}
}
