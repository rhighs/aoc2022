package main

import (
	"fmt"
	"os"
	"strings"
)

type Pos struct {
	x, y int
}

func (p Pos) add(o Pos) Pos {
	return Pos{p.x + o.x, p.y + o.y}
}

type Blizzard struct {
	pos, dir, wrap Pos
}

func readFile(path string) string {
	buf, err := os.ReadFile(path)
	if err != nil {
		panic("Couldn't read file")
	}
	return string(buf)
}

var (
	UP    = Pos{0, -1}
	DOWN  = Pos{0, 1}
	LEFT  = Pos{-1, 0}
	RIGHT = Pos{1, 0}
)

var directions []Pos = []Pos{
	UP,
	DOWN,
	LEFT,
	RIGHT,
}

func parseInput(input string) (field [][]rune, blizzards []Blizzard, start, target Pos) {
	lines := strings.Split(input, "\n")
	start.x = strings.Index(lines[0], "E")
	if start.x == -1 {
		start.x = strings.Index(lines[0], ".")
	}
	target.y = len(lines) - 1
	target.x = strings.Index(lines[target.y], ".")
	for _, line := range lines {
		line = strings.TrimSpace(line)
		field = append(field, []rune(line))
	}
	for y, row := range field {
		for x, c := range row {
			switch c {
			case '^':
				blizzards = append(blizzards, Blizzard{
					Pos{x, y}, UP, Pos{x, len(field) - 2},
				})
			case 'v':
				blizzards = append(blizzards, Blizzard{
					Pos{x, y}, DOWN, Pos{x, 1},
				})
			case '<':
				blizzards = append(blizzards, Blizzard{
					Pos{x, y}, LEFT, Pos{len(field[0]) - 2, y},
				})
			case '>':
				blizzards = append(blizzards, Blizzard{
					Pos{x, y}, RIGHT, Pos{1, y},
				})
			}
		}
	}
	return
}

func inBounds(field [][]rune, pos Pos) bool {
	return pos.x >= 0 && pos.x < len(field[0]) && pos.y >= 0 && pos.y < len(field)
}

func traverseField(field [][]rune, blizzards []Blizzard, start, target Pos) int {
	minutes := 0
	currentStep := make(map[Pos]bool)
	currentStep[start] = true

	for !currentStep[target] {
		//update blizzards
		whereBlizzards := make(map[Pos]bool)
		for i, b := range blizzards {
			bb := b.pos.add(b.dir)
			if inBounds(field, bb) {
				if field[bb.y][bb.x] == '#' {
					blizzards[i].pos = b.wrap
				} else {
					blizzards[i].pos = bb
				}
			}
			whereBlizzards[blizzards[i].pos] = true
		}

		//find new steps
		newStep := make(map[Pos]bool)
		for pos := range currentStep {
			if !(whereBlizzards[pos]) {
				newStep[pos] = true
			}
			for _, d := range directions {
				new := pos.add(d)
				if inBounds(field, new) && field[new.y][new.x] != '#' && !whereBlizzards[new] {
					newStep[new] = true
				}
			}
		}
		currentStep = newStep
		minutes++
	}

	return minutes
}

func p1(input string) int {
	field, blizzards, start, target := parseInput(input)
	return traverseField(field, blizzards, start, target)
}

func p2(input string) int {
	field, blizzards, start, target := parseInput(input)
	return traverseField(field, blizzards, start, target) +
		traverseField(field, blizzards, target, start) +
		traverseField(field, blizzards, start, target)
}

func main() {
	input := strings.TrimSpace(readFile("./input.txt"))
	fmt.Println("p1:", p1(input))
	fmt.Println("p2:", p2(input))
}
