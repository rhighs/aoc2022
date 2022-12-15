package main

import (
    "os"
    "fmt"
    "strings"
    "strconv"
)

const INF = 99999999

func readFile(path string) string {
    buf, err := os.ReadFile(path)
    if err != nil {
        panic("Couldn't read file")
    }
    return string(buf)
}

func abs(n int) int {
    if n > 0 {
        return n
    }
    return -n
}

type Pos struct {
    x, y int
}

func (p Pos) Equal(o Pos) bool {
    return p.x == o.x && p.y == o.y
}

func manhattan(a, b Pos) int {
    return abs(a.x - b.x) + abs(a.y - b.y)
}

type Sensor struct {
    pos, beacon Pos
    d int
}

func (s Sensor) Contains(p Pos) bool {
    return s.d >= manhattan(s.pos, p)
}

type Range struct {
    start, end Pos
    vertical bool
}

func parseInput(input string) (out []Sensor) {
    for _, line := range strings.Split(input, "\n") {
        line = strings.Replace(line, "Sensor at", "", -1)
        line = strings.Replace(line, " closest beacon is at", "", -1)
        line = strings.Replace(line, " x=", "", -1)
        line = strings.Replace(line, " y=", "", -1)
        strcoords := strings.Split(line, ":")
        strsensorcoords := strings.Split(strcoords[0], ",")
        strbeaconcoords := strings.Split(strcoords[1], ",")
        xsensor, _ := strconv.Atoi(strsensorcoords[0])
        ysensor, _ := strconv.Atoi(strsensorcoords[1])
        xbeacon, _ := strconv.Atoi(strbeaconcoords[0])
        ybeacon, _ := strconv.Atoi(strbeaconcoords[1])
        sp := Pos { xsensor, ysensor }
        bp := Pos { xbeacon, ybeacon }
        out = append(out, Sensor {
            sp,
            bp,
            manhattan(sp, bp),
        })
    }
    return
}

func p1(input string, ypos int) int {
    sensors := parseInput(input)

    xmin := INF
    xmax := 0
    for _, s := range sensors {
        if s.pos.x - s.d < xmin {
            xmin = s.pos.x - s.d
        }
        if s.pos.x + s.d > xmax {
            xmax = s.pos.x + s.d
        }
    }

    covered := 0
    testPos := Pos { xmin, ypos }
    for x := xmin; x <= xmax; x++ {
        testPos.x = x
        for _, s := range sensors {
            if s.Contains(testPos) && !s.pos.Equal(testPos) && !s.beacon.Equal(testPos) {
                covered++
                break
            }
        }
    }

    return covered
}

func p2(input string) int {
    sensors := parseInput(input)

    curr := Pos { 0, 0 }
    coveringSensor := Sensor {}
    for {
        covered := false
        for _, s := range sensors {
            covered = s.Contains(curr)
            if covered {
                coveringSensor = s
                break
            }
        }

        if !covered {
            break
        } 

        ydist := abs(coveringSensor.pos.y - curr.y)
        xdist := abs(coveringSensor.pos.x - curr.x)
        skip := coveringSensor.d - ydist + xdist + 1
        
        if curr.x + skip > 4000000 {
            curr.x = 0
            curr.y++
        } else {
            curr.x += skip
        }
    }

    return curr.x * 4000000 + curr.y
}

func main() {
    input := strings.TrimSpace(readFile("./input.txt"))
    fmt.Println("p1:", p1(input, 2000000))
    fmt.Println("p2:", p2(input))
}
