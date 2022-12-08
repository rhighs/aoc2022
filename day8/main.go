package main

import (
    "os"
    "strings"
    "fmt"
)

func readFile(path string) string {
    buf, err := os.ReadFile(path)
    if err != nil {
        panic("Couldn't read file")
    }
    return string(buf)
}

func parseInput(input string) [][]int {
    var out [][]int
    for _, line := range strings.Split(input, "\n") {
        var row []int
        for _, r := range line {
            row = append(row, int(r - '0'))
        }
        out = append(out, row)
    }
    return out
}

func forEachSide(x, y int, input[][]int, each func(int, int) bool) {
    for xx := x-1; xx >= 0; xx-- {
        if each(xx, y) {
            break
        }
    }
    for xx := x+1; xx < len(input); xx++ {
        if each(xx, y) {
            break
        }
    }
    for yy := y-1; yy >= 0; yy-- {
        if each(x, yy) {
            break
        }
    }
    for yy := y+1; yy < len(input); yy++ {
        if each(x, yy) {
            break
        }
    }
}

func findScenicScore(x, y int, input [][]int) int {
    score := 1
    sideScore := 0

    forEachSide(x, y, input, func(xx, yy int) bool {
        onEdge := xx == len(input[0])-1 || yy == len(input)-1 || xx == 0 || yy == 0
        if input[yy][xx] >= input[y][x] || onEdge {
            score *= sideScore + 1
            sideScore = 0
            return true
        }
        sideScore++
        return false
    })

    return score
}

func p1(input [][]int) int {
    cols := len(input[0])
    rows := len(input)

    visible := cols * rows
    for y := 1; y < rows - 1; y++ {
        for x := 1; x < cols - 1; x++ {
            inv := 0
            forEachSide(x, y, input, func(xx, yy int) bool {
                if input[yy][xx] >= input[y][x] {
                    inv++
                    return true
                }
                return false
            })

            if inv == 4 {
                visible--
            }
        }
    }
    
    return visible
}

func p2(input [][]int) int {
    cols := len(input[0])
    rows := len(input)

    maxScore := 0
    for y := 1; y < rows - 1; y++ {
        for x := 1; x < cols - 1; x++ {
            treeScore := findScenicScore(x, y, input)
            if maxScore < treeScore {
                maxScore = treeScore
            }
        }
    }
    
    return maxScore
}

func main() {
    input := strings.TrimSuffix(readFile("./input.txt"), "\n")
    fmt.Println("p1:", p1(parseInput(input)))
    fmt.Println("p2:", p2(parseInput(input)))
}
