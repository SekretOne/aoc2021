package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

// readInput takes a file's name and reads a series of integers, one per line.
func readInput(fileName string) []int {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err = file.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	metrics := make([]int, 0, 16)
	var line int
	for {
		_, err := fmt.Fscanln(file, &line) // scan a line
		if err != nil {
			if err == io.EOF {
				break // stop reading the file
			}
			log.Fatal(err)
		}

		metrics = append(metrics, line)
	}

	return metrics
}

// count the number of times a depth measurement increases from the previous measurement.
// context: https://adventofcode.com/2021/day/1
func totalIncreases(metrics []int) int {
	var total = 0
	for i := 0; i < len(metrics)-1; i++ {
		if metrics[i] < metrics[i+1] {
			total++
		}
	}
	return total
}

// sum all the integers in the slice
func sum(addends []int) int {
	s := 0
	for _, val := range addends {
		s += val
	}
	return s
}

// count the number of times a depth measurement increases in a rolling window
// context: https://adventofcode.com/2021/day/1
func totalIncreasesInWindow(metrics []int, window int) int {
	var total = 0
	for i := 0; i < len(metrics)-window; i++ {
		a := sum(metrics[i : i+window])
		b := sum(metrics[i+1 : i+window+1])
		if a < b {
			total++
		}
	}
	return total
}

func main() {
	metrics := readInput("day/1/input.txt")
	total := totalIncreases(metrics)

	println("Total increases (window size of 1)")
	println(total)

	println("Total increases (window size of 3")
	total = totalIncreasesInWindow(metrics, 3)
	println(total)
}
