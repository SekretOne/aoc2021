package main

import "testing"

func TestExample(t *testing.T) {
	p := Position{0, 0}

	p.Move(&Forward, 5)
	p.Move(&Down, 5)
	p.Move(&Forward, 8)
	p.Move(&Up, 3)
	p.Move(&Down, 8)
	p.Move(&Forward, 2)

	if p.X != 15 || p.Y != 10 {
		t.Errorf("Position = {%d, %d}; want {15, 10}", p.X, p.Y)
	}
}
