package main

import (
	"fmt"
	"os"
	"strings"
)

// A grid is not necessary for this kind of problem as it is not a requirement.
// However, I misunderstood the problem description and used a grid from the begging.
// So i introduced a GRID_SIZE constant in order for it to fit large inputs, note that it might not be
// ideal for very large inputs.
const GRID_SIZE = 256

func readFile(path string) string {
	buf, err := os.ReadFile(path)
	if err != nil {
		fmt.Println("Couldn't read file")
	}
	return string(buf)
}

func parseInput(input string) (out [][]rune, elves int) {
	out = make([][]rune, GRID_SIZE)
	for i := range out {
		out[i] = make([]rune, GRID_SIZE)
		for j := range out[i] {
			out[i][j] = '.'
		}
	}

	strs := strings.Split(input, "\n")
	for y, s := range strs {
		s = strings.TrimSpace(s)
		for x, r := range s {
			if r == '#' {
				elves++
			}
			yy := GRID_SIZE/2 + y - len(strs)/2
			xx := GRID_SIZE/2 + x - len(strs[y])/2
			out[yy][xx] = r
		}
	}
	return out, elves
}

type Pos struct {
	x, y int
}

func (p Pos) add(o Pos) Pos {
	return Pos{p.x + o.x, p.y + o.y}
}

func inBounds(grid [][]rune, p Pos) bool {
	return p.y >= 0 && p.y < len(grid) && p.x >= 0 && p.x < len(grid[p.y])
}

func anyAround(grid [][]rune, pos Pos, what rune) bool {
	around := []Pos{
		pos.add(Pos{-1, -1}),
		pos.add(Pos{0, -1}),
		pos.add(Pos{1, -1}),
		pos.add(Pos{1, 0}),
		pos.add(Pos{1, 1}),
		pos.add(Pos{0, 1}),
		pos.add(Pos{-1, 1}),
		pos.add(Pos{-1, 0}),
	}

	for _, p := range around {
		if inBounds(grid, p) && grid[p.y][p.x] == what {
			return true
		}
	}

	return false
}

func checkAll(grid [][]rune, all []Pos) bool {
	for _, p := range all {
		if inBounds(grid, p) && grid[p.y][p.x] == '#' {
			return false
		}
	}
	return true
}

func northOk(grid [][]rune, pos Pos) bool {
	return checkAll(grid, []Pos{
		pos.add(Pos{-1, -1}),
		pos.add(Pos{0, -1}),
		pos.add(Pos{1, -1}),
	})
}

func southOk(grid [][]rune, pos Pos) bool {
	return checkAll(grid, []Pos{
		pos.add(Pos{-1, 1}),
		pos.add(Pos{0, 1}),
		pos.add(Pos{1, 1}),
	})
}

func westOk(grid [][]rune, pos Pos) bool {
	return checkAll(grid, []Pos{
		pos.add(Pos{-1, -1}),
		pos.add(Pos{-1, 0}),
		pos.add(Pos{-1, 1}),
	})
}

func eastOk(grid [][]rune, pos Pos) bool {
	return checkAll(grid, []Pos{
		pos.add(Pos{1, -1}),
		pos.add(Pos{1, 0}),
		pos.add(Pos{1, 1}),
	})
}

func consider(grid [][]rune, pos Pos, startAt int) Pos {
	cases := []func(Pos) (bool, Pos){
		func(pos Pos) (bool, Pos) {
			if northOk(grid, pos) {
				return true, pos.add(Pos{0, -1})
			}
			return false, pos
		},
		func(pos Pos) (bool, Pos) {
			if southOk(grid, pos) {
				return true, pos.add(Pos{0, 1})
			}
			return false, pos
		},
		func(pos Pos) (bool, Pos) {
			if westOk(grid, pos) {
				return true, pos.add(Pos{-1, 0})
			}
			return false, pos
		},
		func(pos Pos) (bool, Pos) {
			if eastOk(grid, pos) {
				return true, pos.add(Pos{1, 0})
			}
			return false, pos
		},
	}

	for i := 0; i < 4; i++ {
		startAt %= len(cases)
		if b, pnew := cases[startAt](pos); b && inBounds(grid, pnew) {
			return pnew
		}
		startAt++
	}

	return pos
}

func round(grid [][]rune, startAt int) map[Pos][]Pos {
	considerations := make(map[Pos][]Pos)

	for y, row := range grid {
		for x := range row {
			if grid[y][x] == '#' && anyAround(grid, Pos{x, y}, '#') {
				from := Pos{x, y}
				c := consider(grid, from, startAt)
				if !(c.x == from.x && c.y == from.y) {
					considerations[c] = append(considerations[c], from)
				}
			}
		}
	}

	return considerations
}

func rectBounds(grid [][]rune) (int, int, int, int) {
	miny := 9999999
	minx := 9999999
	maxy := 0
	maxx := 0

	for y, row := range grid {
		for x := range row {
			if grid[y][x] == '#' {
				if y < miny {
					miny = y
				}
				if x < minx {
					minx = x
				}
				if y > maxy {
					maxy = y
				}
				if x > maxx {
					maxx = x
				}
			}
		}
	}

	return minx, maxx, miny, maxy
}

func p1(input string) int {
	grid, elves := parseInput(input)
	for i := 0; i < 10; i++ {
		proposals := round(grid, i)
		for k, v := range proposals {
			if len(v) == 1 {
				grid[k.y][k.x] = '#'
				grid[v[0].y][v[0].x] = '.'
			}
		}
	}

	minx, maxx, miny, maxy := rectBounds(grid)
	return (maxx-minx+1)*(maxy-miny+1) - elves
}

func p2(input string) int {
	grid, _ := parseInput(input)
	i := 0
	for {
		proposals := round(grid, i)
		if len(proposals) == 0 {
			return i + 1
		}
		for k, v := range proposals {
			if len(v) == 1 {
				grid[k.y][k.x] = '#'
				grid[v[0].y][v[0].x] = '.'
			}
		}
		i++
	}

	return -1 //never reached
}

func main() {
	input := strings.TrimSpace(readFile("./input.txt"))
	fmt.Println("p1:", p1(input))
	fmt.Println("p2:", p2(input))
}
