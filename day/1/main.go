package main

import (
	"container/list"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	var metrics = list.List{}

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

		metrics.PushBack(line)
	}

	for i := metrics.Front(); i != nil; i = i.Next() {
		fmt.Println(i.Value)
	}
}
