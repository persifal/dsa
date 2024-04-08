package main

import "testing"

func Test(t *testing.T) {
	needle := "sad"
	nl := len(needle)
	heystack := "sadbutsad"
	pref := prefixFunc(needle + "#" + heystack)
	for idx, x := range pref {
		if x == nl {
			e := idx - nl
			s := e - nl
			found := heystack[s:e]
			if found != needle {
				t.Errorf("incorrect found: %s", found)
			}
		}
	}
}
