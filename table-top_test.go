package main

import "testing"

func TestTableTop(t *testing.T) {
	cases := []struct {
		inX, inY int
	}{
		{5, 5},
		{1, 1},
		{0, 0},
	}
	for _, c := range cases {
		
		got := TableTop{c.inX, c.inY}
		
		if got.Width != c.inX {
			t.Errorf("X(%q) == %q, want %q", c.inX, got, c.inX)
		}
		
		if got.Length != c.inY {
			t.Errorf("Y(%q) == %q, want %q", c.inY, got, c.inY)
		}
	}
}