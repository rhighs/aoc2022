package main

import (
    "os"
    "fmt"
    "strings"
    "strconv"
)

func abs(n int) int {
    if n < 0 {
        return -n
    }
    return n
}

func readFile(path string) string {
    buf, err := os.ReadFile(path)
    if err != nil {
        panic("Couldn't read file")
    }
    return string(buf)
}

type Cube struct {
    x, y, z int
}
 
func (c Cube) Add(other Cube) Cube {
    return Cube { c.x + other.x, c.y + other.y, c.z + other.z }
}

func parseInput(input string) (out []Cube) {
    for _, line := range strings.Split(input, "\n") {
        line = strings.TrimSpace(line)
        numbers := strings.Split(line, ",")
        x, _ := strconv.Atoi(numbers[0])
        y, _ := strconv.Atoi(numbers[1])
        z, _ := strconv.Atoi(numbers[2])
        out = append(out, Cube { x, y, z })
    }
    return
}

func adjX(c1, c2 Cube) bool {
    if c1.y == c2.y && c1.z == c2.z {
        return abs(c1.x - c2.x) == 1
    }
    return false
}

func adjY(c1, c2 Cube) bool {
    if c1.z == c2.z && c1.x == c2.x {
        return abs(c1.y - c2.y) == 1
    }
    return false
}

func adjZ(c1, c2 Cube) bool {
    if c1.x == c2.x && c1.y == c2.y {
        return abs(c1.z - c2.z) == 1
    }
    return false
}

func adj(c1, c2 Cube) bool {
    return adjX(c1, c2) || adjY(c1, c2) || adjZ(c1, c2)
}

func surfaceArea(cubes []Cube) int {
    tot := len(cubes) * 6
    for i := 0; i < len(cubes); i++ {
        for j := 0; j < len(cubes); j++ {
            if i != j && adj(cubes[i], cubes[j]) {
                tot--
            }
        }
    }
    return tot
}

func p1(input string) int {
    cubes := parseInput(input)
    return surfaceArea(cubes)
}

func p2(input string) int {
    cubes := parseInput(input)
    cubeMap := make(map[Cube]bool)
    steamMap := make(map[Cube]bool)
    var steam []Cube

    for _, cube := range cubes {
        cubeMap[cube] = true
    }

    minx := 99999999
    miny := 99999999
    minz := 99999999
    maxx := 0
    maxy := 0
    maxz := 0
    for _, cube := range cubes {
        if cube.x < minx {
            minx = cube.x
        }
        if cube.y < miny {
            miny = cube.y
        }
        if cube.z < minz {
            minz = cube.z
        }
        if cube.x > maxx {
            maxx = cube.x
        }
        if cube.y > maxy {
            maxy = cube.y
        }
        if cube.z > maxz {
            maxz = cube.z
        }
    }

    minx-=1 
    miny-=1 
    minz-=1 
    maxx+=1 
    maxy+=1 
    maxz+=1 

    var stack []Cube
    stack = append(stack, Cube {0, 0, 0})

    testCubes := []Cube{
        Cube{1, 0, 0},
        Cube{-1, 0, 0},
        Cube{0, 1, 0},
        Cube{0, -1, 0},
        Cube{0, 0, 1},
        Cube{0, 0, -1},
    }

    for len(stack) > 0 {
        cube := stack[0]
        stack = stack[1:]

        for _, t := range testCubes {
            tc := t.Add(cube)
            if !steamMap[tc] && !cubeMap[tc] && !(tc.x<minx || tc.y<miny || tc.z<minz || tc.x>maxx || tc.y>maxy || tc.z>maxz) {
                stack = append(stack, tc)
                steam = append(steam, tc)
                steamMap[tc] = true
            }
        }
    }

    xlen := maxx-minx+1
    ylen := maxy-miny+1
    zlen := maxz-minz+1
    f1 := xlen * ylen * 2
    f2 := zlen * ylen * 2
    f3 := zlen * xlen * 2
    boxSurfaceArea := f1 + f2 + f3
    return surfaceArea(steam) - boxSurfaceArea
}

func main() {
    input := strings.TrimSpace(readFile("./input.txt"))
    fmt.Println("p1:", p1(input))
    fmt.Println("p1:", p2(input))
}
