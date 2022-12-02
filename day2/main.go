package main

import (
    "fmt"
    "os"
    "strings"
)

var endings = map[string]int {
    "X": 0,
    "Y": 3,
    "Z": 6,
}

var values = map[string]int {
    "A": 1, 
    "B": 2, 
    "C": 3,
    "X": 1,
    "Y": 2,
    "Z": 3,
}

var scores = map[string]int {
    "XA": 3,
    "YB": 3,
    "ZC": 3,
    "XB": 0,
    "XC": 6,
    "YC": 0,
    "YA": 6,
    "ZA": 0,
    "ZB": 6,
}

type Round struct {
    opponent, suggestion string
}

func readFile(path string) string {
    buf, err := os.ReadFile(path)
    if err != nil {
        panic("Couldn't read file")
    }
    return string(buf)
}

func parseInput(input string) (dst []Round) {
    input = strings.TrimSuffix(input, "\n")
    for _, line := range strings.Split(input, "\n") {
        dst = append(dst, Round {
            opponent: string(line[0]),
            suggestion: string(line[2]),
        })
    }
    return
}

func p1(rounds []Round) int {
    final := 0
    for _, r := range rounds {
        base := values[r.suggestion]
        final += scores[r.suggestion+r.opponent] + base
    }
    return final
}

func p2(rounds []Round) int {
    final := 0
    choices := []string{ "X", "Y", "Z" }
    for _, r := range rounds {
        for _, c := range choices {
            if endings[r.suggestion] == scores[c+r.opponent] {
                final += endings[r.suggestion] + values[c]
            }
        }
    }
    return final
}

func main() {
    rounds := parseInput(readFile("./input.txt"))
    fmt.Println("p1:", p1(rounds))
    fmt.Println("p2:", p2(rounds))
}
