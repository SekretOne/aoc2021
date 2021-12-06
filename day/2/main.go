package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

type Position struct {
	X int
	Y int
}

type Direction struct {
	Dx int
	Dy int
}

var (
	Up      = Direction{Dx: 0, Dy: -1}
	Down    = Direction{Dx: 0, Dy: 1}
	Forward = Direction{Dx: 1, Dy: 0}
)

func (p *Position) Move(dir *Direction, dist int) {
	p.X += dir.Dx * dist
	p.Y += dir.Dy * dist
}

type Command struct {
	Dir  *Direction
	Dist int
}

func readInput(fileName string) []Command {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err = file.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	commands := make([]Command, 0, 16)
	var dirStr string
	var dist int
	var dir *Direction
	for {
		_, err := fmt.Fscanf(file, "%s %d\n", &dirStr, &dist)
		if err != nil {
			if err == io.EOF {
				break // stop reading the file
			}
			log.Fatalf("Error while reading file: %v", err)
		}

		switch {
		case dirStr == "forward":
			dir = &Forward
		case dirStr == "down":
			dir = &Down
		case dirStr == "up":
			dir = &Up
		default:
			log.Fatalf("Unrecognized command of '%s'", dirStr)
		}

		commands = append(commands, Command{dir, dist})
	}

	return commands
}

func main() {
	pos := Position{0, 0}
	commands := readInput("day/2/input.txt")
	for _, command := range commands {
		pos.Move(command.Dir, command.Dist)
	}

	fmt.Printf("position{ X: %d, Y: %d }\n", pos.X, pos.Y)
	fmt.Printf("multiplied distance: %d\n", pos.X*pos.Y)
}
