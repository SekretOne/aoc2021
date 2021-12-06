package main

import "testing"

func TestExample(t *testing.T) {
	p := Submarine{0, 0}

	p.Move(&Forward, 5)
	p.Move(&Down, 5)
	p.Move(&Forward, 8)
	p.Move(&Up, 3)
	p.Move(&Down, 8)
	p.Move(&Forward, 2)

	if p.Forward != 15 || p.Depth != 10 {
		t.Errorf("Submarine = {%d, %d}; want {15, 10}", p.Forward, p.Depth)
	}
}
