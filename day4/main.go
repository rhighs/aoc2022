package main

import (
    "os"
    "fmt"
    "strings"
    "strconv"
)

func readFile(path string) string {
    buf, err := os.ReadFile(path)
    if err != nil {
        panic("Couldn't read file")
    }
    return string(buf)
}

func parseInput(lines []string) []ElfPair {
    var pairs []ElfPair
    for _, line := range lines {
        strRanges := strings.Split(line, ",")
        strRange1 := strings.Split(strRanges[0], "-")
        strRange2 := strings.Split(strRanges[1], "-")
        min, _ := strconv.Atoi(strRange1[0])
        max, _ := strconv.Atoi(strRange1[1])
        r1 := Range { min, max }
        min, _ = strconv.Atoi(strRange2[0])
        max, _ = strconv.Atoi(strRange2[1])
        r2 := Range { min, max }
        pairs = append(pairs, ElfPair {r1, r2})
    }
    return pairs
}

type Range struct {
    min, max int
}

type ElfPair struct {
    e1, e2 Range
}

func (r1 *Range) includes(r2 *Range) bool {
    return r1.min <= r2.min && r1.max >= r2.max
}

func (r1 *Range) overlaps(r2 *Range) bool {
    r1r2 := (r1.min <= r2.min && r1.max <= r2.max && r1.max >= r2.min) 
    r2r1 := (r2.min <= r1.min && r2.max <= r1.max && r2.max >= r1.min)
    return r1r2 || r2r1 || r1.includes(r2) || r2.includes(r1)
}

func main() {
    inclusions := 0
    overlappings := 0
    lines := strings.Split(strings.TrimSuffix(readFile("./input.txt"), "\n"), "\n")
    for _, pair := range parseInput(lines) {
        if pair.e1.includes(&pair.e2) || pair.e2.includes(&pair.e1) {
            inclusions += 1
        }
        if pair.e1.overlaps(&pair.e2) {
            overlappings += 1
        }
    }
    fmt.Println("p1:", inclusions)
    fmt.Println("p2:", overlappings)
}
