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

func chase(predator, prey Pos) Pos {
    if !adj(predator, prey) {
        xdist := prey.x - predator.x
        ydist := prey.y - predator.y

        if xdist != 0 && ydist != 0 {
            predator.x += sgn(xdist)
            predator.y += sgn(ydist)
        } else if ydist == 0 {
            predator.x += sgn(xdist)
        } else if xdist == 0 {
            predator.y += sgn(ydist)
        }
    }
    return predator
}

func ropeMovement(input string, nknots int) int {
	knots := make([]Pos, nknots)
	visits := make(map[Pos]bool)
	for i := range knots {
        knots[i] = Pos {}
	}

	visits[knots[0]] = true

	for _, line := range strings.Split(input, "\n") {
		move := strings.Split(strings.TrimSpace(line), " ")
		value, _ := strconv.Atoi(move[1])
        var tailPath []Pos

        for i := 0; i < value; i++ {
		    switch move[0] {
		    case "U":
		    	knots[0].y++
		    case "D":
		    	knots[0].y--
		    case "R":
		    	knots[0].x++
		    case "L":
		    	knots[0].x--
		    }

		    for i := 1; i < nknots; i++ {
		    	knots[i] = chase(knots[i], knots[i-1])
		    	if i == nknots-1 {
                    tailPath = append(tailPath, knots[i])
		    	}
		    }
        }
        for _, p := range tailPath {
            visits[p] = true
        }
	}

	return len(visits)
}

func main() {
	input := strings.TrimSpace(readFile(os.Args[1]))
	fmt.Println("p1:", ropeMovement(input, 2))
    fmt.Println("p2:", ropeMovement(input, 10))
}
