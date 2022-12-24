package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

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

var directions []Pos = []Pos{
	{1, 0},
	{0, 1},
	{-1, 0},
	{0, -1},
}

func parseInput(input string) ([]string, []any) {
	input = strings.Replace(input, "\r\n", "\n", -1)
	pathAndCommands := strings.Split(input, "\n\n")
	path := strings.Split(pathAndCommands[0], "\n")
	strcommands := pathAndCommands[1]

	var commands []any

	strnum := ""
	for _, c := range strcommands {
		if c >= '0' && c <= '9' {
			strnum += string(c)
			continue
		} else if strnum != "" {
			n, _ := strconv.Atoi(strnum)
			commands = append(commands, n)
			strnum = ""
		}
		commands = append(commands, string(c))
	}

	if strnum != "" {
		n, _ := strconv.Atoi(strnum)
		commands = append(commands, n)
		strnum = ""
	}

	return path, commands
}

func startH(path string) int {
	d := strings.Index(path, ".")
	h := strings.Index(path, "#")
	if h == -1 {
		return d
	}
	return min(d, h)
}

func startV(path []string, from Pos) int {
	for y := from.y; y > 0; y-- {
		if len(path[y-1]) <= from.x || path[y-1][from.x] == ' ' {
			return y
		}
	}
	return 0
}

func endV(path []string, from Pos) int {
	for y := from.y; y < len(path)-1; y++ {
		if len(path[y+1]) <= from.x || path[y+1][from.x] == ' ' {
			return y
		}
	}
	return len(path) - 1
}

func stepH(path string, start, x, dir int) (bool, int) {
	if dir == 0 {
		return false, x
	}

	if x+dir > len(path)-1 {
		if path[start] == '#' {
			return true, x
		} else {
			return false, start
		}
	}

	if x+dir < start {
		if path[len(path)-1] == '#' {
			return true, x
		} else {
			return false, len(path) - 1
		}
	}

	if path[x+dir] == '#' {
		return true, x
	}

	return false, x + dir
}

func stepV(path []string, pos Pos, start, end, dir int) (bool, int) {
	if dir == 0 {
		return false, pos.y
	}

	if pos.y+dir == end+1 {
		if path[start][pos.x] == '#' {
			return true, pos.y
		} else {
			return false, start
		}
	}

	if pos.y+dir == start-1 {
		if path[end][pos.x] == '#' {
			return true, pos.y
		} else {
			return false, end
		}
	}

	if path[pos.y+dir][pos.x] == '#' {
		return true, pos.y
	}

	return false, pos.y + dir
}

func p1(input string) int {
	path, commands := parseInput(input)

	dir := 0
	start := strings.Index(path[0], ".")
	pos := Pos{start, 0}

	for i := 0; i < len(commands); i++ {
		switch commands[i].(type) {
		case int:
			goNext := false
			for m := 0; m < commands[i].(int); m++ {
				if directions[dir].x != 0 {
					goNext, pos.x = stepH(path[pos.y], startH(path[pos.y]), pos.x, directions[dir].x)
				} else {
					goNext, pos.y = stepV(path, pos, startV(path, pos), endV(path, pos), directions[dir].y)
				}
				if goNext {
					break
				}
			}

			if goNext {
				continue
			}
		case string:
			switch commands[i].(string) {
			case "R":
				dir = (dir + 1) % len(directions)
			case "L":
				dir--
				if dir == -1 {
					dir = len(directions) - 1
				}
			}
		}
	}

	return 1000*(pos.y+1) + 4*(pos.x+1) + dir
}

func main() {
	input := readFile("./input.txt")
	fmt.Println("p1:", p1(input))
}
