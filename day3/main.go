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

func asciiMap(asciiValue int) int {
    if asciiValue > 96 && asciiValue < 123 {
        return asciiValue - 96
    } else if asciiValue > 64 && asciiValue < 91 {
        return 27 + asciiValue - 65
    }
    return 0
}

func matchingPriorities(s1 string, s2 string) [53]int {
    var matches [53]int
    r1 := []rune(s1)
    r2 := []rune(s2)
    for i := 0; i < len(r1); i++ {
        matches[asciiMap(int(r1[i]))] = 1
    }
    for i := 0; i < len(r2); i++ {
        p := asciiMap(int(r2[i]))
        if matches[p] == 1 {
            matches[p] = 2
        }
    }
    for i := 0; i < 53; i++ {
        matches[i] /= 2
    }
    return matches
}

func p1(lines []string) int {
    psum := 0
    for _, sack := range lines {
        halfLen := len(sack)/2
        s1 := sack[0:halfLen]
        s2 := sack[halfLen:len(sack)]
        matches := matchingPriorities(s1, s2)
        for i := 0; i < len(matches); i++ {
            psum += matches[i] * i
        }
    }
    return psum
}

func p2(lines []string) int {
    psum := 0
    for i := 0; i < len(lines) - 3; i += 3 {
        e1 := lines[i]
        e2 := lines[i+1]
        e3 := lines[i+2]
        me1e2 := matchingPriorities(e1, e2)
        me2e3 := matchingPriorities(e2, e3)
        for j := 0; j < len(me1e2); j++ {
            if me1e2[j] + me2e3[j] == 2 {
                psum += j
            }
        }
    }
    return psum
}

func main() {
    lines := strings.Split(readFile("./input.txt"), "\n")
    fmt.Println("p1:", p1(lines))
    fmt.Println("p2:", p2(lines))
}
