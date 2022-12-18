package main

import (
	"fmt"
	"os"
	"strings"
)

func readFile(path string) string {
	buf, err := os.ReadFile(path)
	if err != nil {
		panic("Couldn't read file")
	}
	return string(buf)
}

var rocks [][]string = [][]string{
	{"####"},
	{".#.", "###", ".#."},
	{"..#", "..#", "###"},
	{"#", "#", "#", "#"},
	{"##", "##"},
}

func checkCollision(chamber [][]rune, rock []string, x, y uint64) bool {
	rockH := uint64(len(rock))
	rockW := uint64(len(rock[0]))

	if rockW+x-1 >= uint64(len(chamber[y])) || rockH+y-1 >= uint64(len(chamber)) {
		return true
	}

	for ry := uint64(0); ry < rockH; ry++ {
		for rx := uint64(0); rx < rockW; rx++ {
			if chamber[y+ry][x+rx] == '#' && rock[ry][rx] == '#' {
				return true
			}
		}
	}

	return false
}

func tryMove(chamber [][]rune, rock []string, x, y uint64, dir rune) uint64 {
	switch dir {
	case '<':
		if x > 0 && !checkCollision(chamber, rock, x-1, y) {
			return x - 1
		}
	case '>':
		if !checkCollision(chamber, rock, x+1, y) {
			return x + 1
		}
	}

	return x
}

func drawRock(chamber [][]rune, rock []string, x, y uint64) {
	rockH := uint64(len(rock))
	rockW := uint64(len(rock[0]))

	for ry := uint64(0); ry < rockH; ry++ {
		for rx := uint64(0); rx < rockW; rx++ {
			if chamber[y+ry][x+rx] == '.' && rock[ry][rx] == '#' {
				chamber[y+ry][x+rx] = '#'
			}
		}
	}
}

func clearChamber(chamber [][]rune) [][]rune {
    top := "......."
    for i := 0; i < len(chamber); i++ {
        if string(chamber[i]) != top {
            top = string(chamber[i])
            break
        }
    }
    for i := 0; i < len(chamber); i++ {
        chamber[i] = []rune(".......")
    }
    chamber[len(chamber)-1] = []rune(top)
    return chamber
}

type StartingIds struct {
    ri, ii  uint64
    topRock string
}

type PatternData struct {
    placed, height uint64
}

func problem(input string, maxplacements uint64) uint64 {
	var chamber [][]rune
	for i := uint64(0); i < 512; i++ {
		chamber = append(chamber, []rune("......."))
	}
	chamber = append(chamber, []rune("#######"))

	ii := uint64(0)
	ri := uint64(0)
	placed := uint64(0)
    totalheight := uint64(0)
    heightWithPatterns := uint64(0)
	height := uint64(len(chamber) - 1)

	patterns := make(map[StartingIds]PatternData)
    patternFound := false

	for placed < maxplacements {
		ri = ri % uint64(len(rocks))
		rock := rocks[ri]
		rockH := uint64(len(rock))

        if !patternFound {
            sp := StartingIds{ri, ii, string(chamber[height])}
            if patterns[sp].placed == 0 {
                patterns[sp] = PatternData{placed, height}
            } else {
                fmt.Println("pattern starts at", patterns[sp].placed)
                patternFound = true
                pdiff := uint64(placed - patterns[sp].placed)
                hdiff := uint64(patterns[sp].height - height)
                q := uint64(maxplacements) / pdiff
                q--
                placed += pdiff * q
                heightWithPatterns = hdiff*q + patterns[sp].height
                totalheight += heightWithPatterns
                fmt.Println(q, pdiff, hdiff)
            }
        }

		x := uint64(2)
		y := height - (3 + rockH)
		for ; y+rockH < uint64(len(chamber)); y++ {
			ii = ii % uint64(len(input))
			gasdir := input[ii]
			x = tryMove(chamber, rock, x, y, rune(gasdir))
			if checkCollision(chamber, rock, x, y+1) {
				break
			}
			ii++
		}

		drawRock(chamber, rock, x, y)

		if y < height {
			height = y
		}

        if height <= 8 {
            chamber = clearChamber(chamber)
            totalheight += uint64(len(chamber)) - height - 1
            height = uint64(len(chamber) - 1)
        }

		ri++
		placed++
	}

	return totalheight
}

func main() {
	input := strings.TrimSpace(readFile("./input2.txt"))
	fmt.Println(problem(input, 2022))
	fmt.Println(problem(input, 1000000000000))
}
