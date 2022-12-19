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

func p2(input string, rayLen int) int {
    cubes := parseInput(input)
    cubeMap := make(map[Cube]bool)

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

    enclosed := func (c Cube) bool {
        dirsFound := make(map[int]bool)
        for offset := 1; offset <= rayLen; offset++ {
            if cubeMap[Cube { c.x + offset, c.y, c.z }] {
                dirsFound[0] = true
            }
            if cubeMap[Cube { c.x - offset, c.y, c.z }] {
                dirsFound[1] = true
            }
            if cubeMap[Cube { c.x, c.y + offset, c.z }] {
                dirsFound[2] = true
            }
            if cubeMap[Cube { c.x, c.y - offset, c.z }] {
                dirsFound[3] = true
            }
            if cubeMap[Cube { c.x, c.y, c.z + offset }] {
                dirsFound[4] = true
            }
            if cubeMap[Cube { c.x, c.y, c.z - offset }] {
                dirsFound[5] = true
            }
        }
        return len(dirsFound) == 10
    }

    for x := minx; x < maxx; x++ {
        for y := miny; y < maxy; y++ {
            for z := minz; z < maxz; z++ {
                c := Cube{x, y, z}
                if !cubeMap[c] && enclosed(c) {
                    cubes = append(cubes, c)
                    cubeMap[c] = true
                }
            }
        }
    }

    return surfaceArea(cubes)
}

func main() {
    input := strings.TrimSpace(readFile("./input.txt"))
    fmt.Println(p1(input))
    fmt.Println(p2(input, 20))
}
