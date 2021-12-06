package main

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

func main() {

}
