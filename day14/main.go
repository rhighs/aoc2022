package main

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

func (r Range) asPositions() (out []Pos) {
    if r.vertical {
        ymin, ymax := minmax(r.start.y, r.end.y)
        for y := ymin; y <= ymax; y++ {
            out = append(out, Pos {r.start.x, y})
        }
        return out
    }
    xmin, xmax := minmax(r.start.x, r.end.x)
    for x := xmin; x <= xmax; x++ {
        out = append(out, Pos {x, r.start.y})
    }
    return out
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

func fall(from Pos, grains *map[Pos]bool) (Pos, bool) {
    down := Pos { 0, 1 }
    left := Pos { -1, 1 }
    right := Pos { 1, 1 }

    to := from
    canFallMore := true

    switch {
    case !(*grains)[from.Add(Pos { 0, 1 })]:
        to = to.Add(down)
    case !(*grains)[from.Add(Pos { -1, 1 })]:
        to = to.Add(left)
    case !(*grains)[from.Add(Pos { 1, 1 })]:
        to = to.Add(right)
    default:
        canFallMore = false
    }

    if !canFallMore {
        (*grains)[to] = true
    }

    return to, canFallMore
}

func fallingSand(ranges []Range, voidThresh int) int {
    deposited := 0
    sandPos := Pos {}
    source := Pos { 500, 0 }
    grains := make(map[Pos]bool)

    for _, r := range ranges {
        for _, p := range r.asPositions() {
            grains[p] = true
        }
    }

    intoTheVoid := false;
    for !(sandPos.x == source.x && sandPos.y == source.y) && !intoTheVoid {
        sandPos = source
        canFallMore := true

        for canFallMore && !intoTheVoid {
            sandPos, canFallMore = fall(sandPos, &grains)
            intoTheVoid = sandPos.y > voidThresh
        }

        deposited++
    }

    if intoTheVoid {
        deposited--
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

