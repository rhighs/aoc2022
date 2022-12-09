package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Pos struct {
	x, y int
}

func readFile(path string) string {
	buf, err := os.ReadFile(path)
	if err != nil {
		panic("Couldn't read file")
	}
	return string(buf)
}

func sgn(n int) int {
	if n >= 0 {
		return 1
	}
	return -1
}

func abs(n int) int {
	if n >= 0 {
		return n
	}
	return -n
}

func adj(a, b Pos) bool {
	d := Pos{
		x: abs(a.x - b.x),
		y: abs(a.y - b.y),
	}
	return d.x <= 1 && d.y <= 1
}

func __reach(start, target Pos, step func(Pos)) Pos {
	curr := start
	for {
		if adj(curr, target) {
			break
		}
		xdist := target.x - curr.x
		ydist := target.y - curr.y

		if xdist != 0 && ydist != 0 {
			curr.x += sgn(xdist)
			curr.y += sgn(ydist)
		} else if ydist == 0 {
			curr.x += sgn(xdist)
		} else if xdist == 0 {
			curr.y += sgn(ydist)
		}
		step(curr)
	}
	return curr
}

func reach(start, target Pos) Pos {
	return __reach(start, target, func(_ Pos) {})
}

func findPath(start, target Pos) []Pos {
	var path []Pos
	__reach(start, target, func(curr Pos) {
		path = append(path, curr)
	})
	return path
}

func ropeMovement(inputLines []string, nknots int) int {
	knots := make([]Pos, nknots)
	visits := make(map[Pos]bool)
	s := Pos{
		x: 0,
		y: 0,
	}

	for i := range knots {
		knots[i] = s
	}

	visits[knots[0]] = true

	for _, line := range inputLines {
		move := strings.Split(strings.TrimSpace(line), " ")
		value, _ := strconv.Atoi(move[1])

		switch move[0] {
		case "U":
			knots[0].y += value
		case "D":
			knots[0].y -= value
		case "R":
			knots[0].x += value
		case "L":
			knots[0].x -= value
		}

		for i := 1; i < nknots; i++ {
			path := findPath(knots[i], knots[i-1])
			if i == nknots-1 {
				for _, p := range path {
					visits[p] = true
				}
			}
			if len(path) > 0 {
				knots[i] = path[len(path)-1]
			}
		}
	}

	return len(visits)
}

func main() {
	input := strings.TrimSpace(readFile(os.Args[1]))
	lines := strings.Split(input, "\n")
	for i := 2; i < 11; i++ {
		fmt.Println("p1:", i, ropeMovement(lines, i))
	}
}
