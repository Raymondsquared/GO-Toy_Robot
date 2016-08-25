package main

import "testing"

func TestValidTableTop(t *testing.T) {
	cases := []struct {
		inX, inY int
	}{
		{5, 5},
		{1, 1},
		{10, 1},
		{1, 6},
	}
	for _, c := range cases {
		
		got,e := NewTableTop(c.inX, c.inY)
		
		if e != nil {
			t.Errorf("error thrown")
		}
		if got.Width != c.inX {
			t.Errorf("X(%q) == %q, want %q", got, got.Width, c.inX)
		}
		
		if got.Length != c.inY {
			t.Errorf("Y(%q) == %q, want %q", got, got.Length, c.inY)
		}
	}
}

func TestInvalidTableTop(t *testing.T) {
	cases := []struct {
		inX, inY int
	}{
		{-1, 1},
		{2, -2},
		{-3, -3},
		{0, 0},
	}
	for _, c := range cases {
		got, e := NewTableTop(c.inX, c.inY)
		if e == nil {
			t.Errorf("error should be thrown for invlaid value")
		}
		if got != nil {
			t.Errorf("table top should be null for invlaid value")
		}
	}
}