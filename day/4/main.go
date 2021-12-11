package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

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

// marks numbers and returns the index of the space that was marked (or -1 if not) and whether this was winning
func (b *bingo) mark(number int) (spaceIndex int, isWinner bool) {
	for i, space := range b.spaces {
		if number == space {
			b.marks |= 1 << i
			return i, b.hasWon()
		}
	}
	return -1, false
}

// scores the bingo game (sum of unused numbers times the winning number)
func score(b *bingo, winningNumber int) int {
	sumOfUnmarked := 0
	for i := 0; i < 25; i++ {
		if 1<<i&b.marks == 0 { //if unmarked
			sumOfUnmarked += b.spaces[i]
		}
	}
	return winningNumber * sumOfUnmarked
}

func readInput(fileName string) (numbers []int, bingos []*bingo) {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatalln(err)
	}
	defer func() {
		if err = file.Close(); err != nil {
			log.Fatalln(err)
		}
	}()

	var rawNumbers string
	if _, err := fmt.Fscanln(file, &rawNumbers); err != nil {
		log.Fatalln(err)
	}
	splits := strings.Split(rawNumbers, ",")
	numbers = make([]int, 0, len(splits))

	for _, s := range splits {
		n, err := strconv.Atoi(s)
		if err != nil {
			log.Fatalf("unable to convert into number array, value: %v\n", s)
		}
		numbers = append(numbers, n)
	}

	bingos = make([]*bingo, 0, 4)

	for {
		b := bingo{
			spaces: [25]int{},
			marks:  0,
		}
		// read the 25 spaces
		if _, err := fmt.Fscanln(file); err != nil {
			log.Fatalf("unexpected format; is not a bingo card")
		}

		for i := 0; i < 25; i += 5 {
			if _, err := fmt.Fscanln(file, &b.spaces[i], &b.spaces[i+1], &b.spaces[i+2], &b.spaces[i+3], &b.spaces[i+4]); err != nil {
				if err == io.EOF {
					return numbers, bingos
				}
				log.Fatalf("unable to read number %v: %v\n", i, err)
			}
		}
		bingos = append(bingos, &b)
	}
}

type bingo struct {
	spaces [25]int
	marks  int
}

func main() {
	numbers, cards := readInput("day/4/input.txt")

	for _, n := range numbers {
		for _, card := range cards {
			if _, won := card.mark(n); won {

				s := score(card, n)
				fmt.Printf("card: {\n  %v,\n%025b,\n}\n", card.spaces, card.marks)
				fmt.Printf("winning score: %v\n", s)
				return
			}
		}
	}
	fmt.Println("Nothing won?")
}
