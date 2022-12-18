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

func checkCollision(chamber [][]rune, rock []string, x, y int) bool {
	rockH := len(rock)
	rockW := len(rock[0])

	if rockW+x-1 >= len(chamber[y]) || rockH+y-1 >= len(chamber) {
		return true
	}

	for ry := 0; ry < rockH; ry++ {
		for rx := 0; rx < rockW; rx++ {
			if chamber[y+ry][x+rx] == '#' && rock[ry][rx] == '#' {
				return true
			}
		}
	}

	return false
}

func tryMove(chamber [][]rune, rock []string, x, y int, dir rune) int {
	switch dir {
	case '<':
		if x - 1 >= 0 && !checkCollision(chamber, rock, x-1, y) {
			return x - 1
		}
	case '>':
		if !checkCollision(chamber, rock, x+1, y) {
			return x + 1
		}
	}

	return x
}

func drawRock(chamber [][]rune, rock []string, x, y int) {
	rockH := len(rock)
	rockW := len(rock[0])

	for ry := 0; ry < rockH; ry++ {
		for rx := 0; rx < rockW; rx++ {
			if chamber[y+ry][x+rx] == '.' && rock[ry][rx] == '#' {
				chamber[y+ry][x+rx] = '#'
			}
		}
	}
}

func clearChamber(chamber *[][]rune) {
    top := "......."
    for i := 0; i < len(*chamber); i++ {
        if string((*chamber)[i]) != top {
            top = string((*chamber)[i])
            break
        }
    }
    for i := 0; i < len(*chamber); i++ {
        (*chamber)[i] = []rune(".......")
    }
    (*chamber)[len(*chamber)-1] = []rune(top)
}

func initChamber(height int) (chamber [][]rune) {
	for i := 0; i < height - 1; i++ {
		chamber = append(chamber, []rune("......."))
	}
	chamber = append(chamber, []rune("#######"))
    return
}

type StartingIds struct {
    ri, ii  int
    topRock string
}

type PatternData struct {
    placed, height uint64
}

func p1(input string, maxplacements int) int {
    chamber := initChamber(maxplacements * 4)

	ii := 0
	ri := 0
	placed := 0
	height := len(chamber) - 1

	for placed < maxplacements {
		ri = ri % len(rocks)
		rock := rocks[ri]
		rockH := len(rock)

		x := 2
		y := height - (3 + rockH)
		for ; y+rockH < len(chamber); y++ {
			ii = ii % len(input)
			gasdir := input[ii]
			x = tryMove(chamber, rock, x, y, rune(gasdir))
			ii++
			if checkCollision(chamber, rock, x, y+1) {
				break
			}
		}

		drawRock(chamber, rock, x, y)

		if y < height {
			height = y
		}

		ri++
		placed++
	}

    return len(chamber) - height - 1
}

func p2(input string, maxplacements uint64) uint64 {
    chamber := initChamber(100)

	ii := 0
	ri := 0
	placed := uint64(0)
    patternFound := false
	height := len(chamber) - 1
    patternsHeight := uint64(0)
    incrementalHeight := uint64(0)

    patterns := make(map[StartingIds]PatternData)

    flush := func() (uint64, int) {
        if height <= 8 {
            tmp := uint64(len(chamber) - 1 - height)
            clearChamber(&chamber)
            return incrementalHeight + tmp, len(chamber) - 1
        }
        return incrementalHeight, height
    }

	for {
		ri = ri % len(rocks)
		rock := rocks[ri]
		rockH := len(rock)

        if !patternFound {
            sp := StartingIds { ri, ii, string(chamber[height]) }
            fmt.Println(sp)
            patternFound = patterns[sp].placed != 0
            if !patternFound {
                patterns[sp] = PatternData { placed, incrementalHeight }
            } else {
                patternData := patterns[sp]
                placedDiff := placed - patternData.placed
                heightDiff := incrementalHeight - patternData.height
                q := maxplacements / placedDiff
                patternsHeight = patternData.height + heightDiff * (q - 1)
                placed = patternData.placed + placedDiff * (q - 1)

                fmt.Println(placedDiff, heightDiff, q, patternsHeight, placed, incrementalHeight, patternData.height, patternData)
            }
        }

        if uint64(placed) >= maxplacements  {
            break
        }

		x := 2
		y := height - (3 + rockH)
		for ; y+rockH < len(chamber); y++ {
			ii = ii % len(input)
			gasdir := input[ii]
			ii++

			x = tryMove(chamber, rock, x, y, rune(gasdir))
			if checkCollision(chamber, rock, x, y+1) {
				break
			}
		}

		if y < height {
			height = y
		}

		drawRock(chamber, rock, x, y)

        incrementalHeight, height = flush()

		ri++
		placed++
	}

    return patternsHeight + incrementalHeight
}

func main() {
	input := strings.TrimSpace(readFile("./input2.txt"))
	fmt.Println(p1(input, 2022))
	fmt.Println(p2(input, 1000000000000))
}
