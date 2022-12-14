package main

// This could be optimized a bit. I used a map to speed up lookup for sand grain positions,
// I think there's room for more though. Ranges could be used to insert rock positions in the
// map mentioned above also something similar to quadtrees could be used instead, but idk, this works bye...

import (
    "os"
    "fmt"
    "strings"
    "strconv"
)

const INF = 10000000

type CellT rune
const (
    sand CellT = 'o'
    rock = '#'
    air = '.'
)

func minmax(a, b int) (int, int) {
    if a < b {
        return a, b
    }
    return b, a
}

type Pos struct {
    x, y int
}

func (p Pos) Equal(other Pos) bool {
    return p.x == other.x && p.y == other.y
}

func (p Pos) Add(other Pos) Pos {
    return Pos {
        p.x + other.x,
        p.y + other.y,
    }
}

type Range struct {
    start, end Pos
    vertical bool
}

func (r Range) includes(p Pos) bool {
    if r.vertical {
        ymin, ymax := minmax(r.start.y, r.end.y)
        return p.x == r.start.x && p.y >= ymin && p.y <= ymax
    }
    xmin, xmax := minmax(r.start.x, r.end.x)
    return p.y == r.start.y && p.x >= xmin && p.x <= xmax
}

func readFile(path string) string {
    buf, err := os.ReadFile(path)
    if err != nil {
        panic("Couldn't read file")
    }
    return string(buf)
}

func parsePositions(strPositions []string) []Pos {
    var positions []Pos

    for _, strpos := range strPositions {
        strcoords := strings.Split(strpos, ",")
        x, _ := strconv.Atoi(strcoords[0])
        y, _ := strconv.Atoi(strcoords[1])
        positions = append(positions, Pos { x, y })
    }

    return positions
}

func parseInput(input string) (ranges []Range) {
    for _, line := range strings.Split(input, "\n") {
        strslices := strings.Split(line, " -> ")
        positions := parsePositions(strslices)

        for i := 1; i < len(positions); i++ {
            prev := positions[i-1]
            curr := positions[i]
            ranges = append(ranges, Range {
                prev,
                curr,
                prev.x == curr.x,
            })
        }
    }

    return ranges
}

func fall(from Pos, ranges *[]Range, grains *map[Pos]bool) (Pos, bool) {
    down := Pos { 0, 1 }
    left := Pos { -1, 1 }
    right := Pos { 1, 1 }

    fallStep := func(from, dir Pos) bool {
        maybeInto := from.Add(dir)
        for _, r := range *ranges {
            if r.includes(maybeInto) || (*grains)[maybeInto] {
                return false
            }
        }
        return true
    }

    canFallMore := false
    to := from

    if fallStep(to, down) {
        canFallMore = true
        to = to.Add(down)
    } else if fallStep(to, left) {
        canFallMore = true
        to = to.Add(left)
    } else if fallStep(to, right) {
        canFallMore = true
        to = to.Add(right)
    }

    if !canFallMore {
        (*grains)[to] = true
    }

    return to, canFallMore
}

func fallingSand(ranges []Range, voidThresh int) int {
    source := Pos { 500, 0 }
    sandPos := Pos {}
    deposited := 0

    grains := make(map[Pos]bool)

    intoTheVoid := false
    for !sandPos.Equal(source) {
        sandPos = source
        canFallMore := false

        for {
            sandPos, canFallMore = fall(sandPos, &ranges, &grains)
            intoTheVoid = sandPos.y > voidThresh
            if !canFallMore || intoTheVoid {
                break
            }
        }

        if intoTheVoid {
            break
        }

        deposited++
    }

    return deposited
}

func p1(input string) int {
    return fallingSand(parseInput(input), 200)
}

func p2(input string) int {
    ranges := parseInput(input)
    positions := parsePositions(strings.Split(strings.Replace(input, "\n", " -> ", -1), " -> "))
    maxy := 0
    for _, pos := range positions {
        if pos.y > maxy {
            maxy = pos.y
        }
    }
    maxy += 2
    ranges = append(ranges, Range {
        Pos { -INF, maxy },
        Pos { INF, maxy },
        false,
    })
    return fallingSand(ranges, maxy)
}

func main() {
    input := strings.TrimSpace(readFile("./input.txt"))
    fmt.Println(p1(input))
    fmt.Println(p2(input))
}

