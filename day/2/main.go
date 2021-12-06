package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

type Submarine struct {
	Forward int
	Depth   int
}

type Submarine2 struct {
	Forward int
	Depth   int
	Aim     int
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

// Move the submarine in the given direction by distance
func (p *Submarine) Move(dir *Direction, dist int) {
	p.Forward += dir.Dx * dist
	p.Depth += dir.Dy * dist
}

// Move the submarine based off direction, distance and the sub's Aim
// Down X increases your aim by X units.
// Up X decreases your aim by X units.
// Forward X does two things:
//   It increases your horizontal position by X units.
//   It increases your depth by your aim multiplied by X.
func (p *Submarine2) Move(dir *Direction, dist int) {
	p.Aim += dir.Dy * dist

	p.Forward += dir.Dx * dist
	p.Depth += dir.Dx * dist * p.Aim
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
	sub1 := Submarine{0, 0}
	commands := readInput("day/2/input.txt")
	for _, command := range commands {
		sub1.Move(command.Dir, command.Dist)
	}

	fmt.Printf("Submarine 1 { Forward: %d, Depth: %d }\n", sub1.Forward, sub1.Depth)
	fmt.Printf("multiplied distance: %d\n", sub1.Forward*sub1.Depth)

	sub2 := Submarine2{
		Forward: 0,
		Depth:   0,
		Aim:     0,
	}

	for _, command := range commands {
		sub2.Move(command.Dir, command.Dist)
	}

	fmt.Printf("\nSubmarine 2 (with aim) { Forward: %d, Depth: %d, Aim: %d }\n", sub2.Forward, sub2.Depth, sub2.Aim)
	fmt.Printf("multiplied distance: %d\n", sub2.Forward*sub2.Depth)

}
