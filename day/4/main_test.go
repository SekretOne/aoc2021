package main

import "testing"

func Test_MarkBingo(t *testing.T) {

	b := bingo{
		spaces: [25]int{14, 16, 18, 20, 22, 2, 4, 6, 8, 10, 12, 24, 26, 28, 30, 32, 34, 36, 38, 40, 42, 44, 46, 48, 50},
		marks:  0,
	}

	nums := [...]int{4, 8, 10, 12, 20, 13, 17, 6, 15, 2, 16}

	for _, n := range nums {
		if _, won := b.mark(n); won {
			if n != 2 {
				t.Errorf("expect to win with 2, but got %v", n)
			}
			return
		}
	}

	t.Error("didn't win, when expected to win on 2 being called")
}
