package main

const (
	rowWin = int(0b0000000000000000000011111)
	colWin = int(0b0000100001000010000100001)
)

var (
	winningMarks = [10]int{
		rowWin,
		rowWin << 5,
		rowWin << 10,
		rowWin << 15,
		rowWin << 20,
		colWin,
		colWin << 5,
		colWin << 10,
		colWin << 15,
		colWin << 20,
	}
)

// bingo has won if winning mark bit mask matches current marks
func (b *bingo) hasWon() bool {
	for _, w := range winningMarks {
		if w&b.marks == w {
			return true
		}
	}
	return false
}

// marks numbers
func (b *bingo) mark(number int) (spaceIndex int, isWinner bool) {
	for i, space := range b.spaces {
		if number == space {
			b.marks |= 1 << i
			return i, b.hasWon()
		}
	}
	return -1, false
}

type bingo struct {
	spaces [25]int
	marks  int
}

func main() {

}
