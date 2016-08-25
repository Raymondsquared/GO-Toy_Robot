package main

import "testing"

func TestRobot(t *testing.T) {
	cases := []struct {
		inX, inY int
	}{
		{5, 5},
		{1, 1},
		{0, 0},
	}
	for _, c := range cases {
		
		got := Robot{c.inX, c.inY}
		
		if got.X != c.inX {
			t.Errorf("X(%q) == %q, want %q", got, got.X, c.inX)
		}
		
		if got.Y != c.inY {
			t.Errorf("Y(%q) == %q, want %q", got, got.Y, c.inY)
		}
	}
}