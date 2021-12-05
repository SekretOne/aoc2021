package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	var metrics = make([]int, 0, 16)

	file, err := os.Open("day/1/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err = file.Close(); err != nil {
			log.Fatal(err)
		}
	}()

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

	var totalIncreases = 0
	for i := 1; i < len(metrics); i++ {
		if metrics[i-1] < metrics[i] {
			totalIncreases++
		}
	}
	println(totalIncreases)
}
