
package main

import (
    "fmt"
    "os"
    "strconv"
    "strings"
)

type stack []rune

func (s stack) Push(value rune) stack {
    return append(s, value)
}

func (s stack) Pop() (stack, rune) {
    l := len(s)
    if l == 0 {
        return nil, 0
    }
    return s[:l-1], s[l-1]
}

type Move struct {
    qty, from, to int
}

type PuzzleInput struct {
    crates []stack
    moves []Move
}

func readFile(path string) string {
    buf, err := os.ReadFile(path)
    if err != nil {
        panic("Couldn't read file")
    }
    return string(buf)
}

func parseInput(raw string) PuzzleInput {
    var pinput PuzzleInput

    input := strings.Split(raw, "\n\n")

    strCrates := input[0]
    strCrates = strings.ReplaceAll(strings.ReplaceAll(strCrates, "[", " "), "]", " ")
    crateLines := strings.Split(strCrates, "\n")
    crateLines = crateLines[:len(crateLines)-1]
    for i, line := range crateLines {
        crateLines[i] = strings.Trim(line, " ")
    }
    // Horrible hacks done here, this is NOT a parser
    for k := 0; k < 2; k++ {
        for i, line := range crateLines {
            var newLine string
            for j := 0; j < len(line); j++ {
                if j & 1 == 0 {
                    newLine += string(line[j])
                }
            }
            crateLines[i] = newLine
        }
    }

    // Parse crates
    pinput.crates = make([]stack, len(crateLines[len(crateLines)-1]))
    for i := len(crateLines) - 1; i >= 0; i-- {
        line := crateLines[i]
        for j := 0; j < len(line); j++ {
            char := rune(line[j])
            if char != ' ' {
                pinput.crates[j] = pinput.crates[j].Push(char)
            }
        }
    }

    //Parse moves
    strMoves := input[1]
    moveLines := strings.Split(strMoves, "\n")
    moveLines = moveLines[:len(moveLines)-1]
    var moves []Move

    for _, line := range moveLines {
        line = strings.Trim(line, " ")
        line = strings.ReplaceAll(line, " ", "")
        line = strings.ReplaceAll(line, "move", "")
        line = strings.ReplaceAll(line, "from", " ")
        line = strings.ReplaceAll(line, "to", " ")
        numbers := strings.Split(line, " ")
        qty, _ := strconv.Atoi(numbers[0])
        from, _ := strconv.Atoi(numbers[1])
        to, _ := strconv.Atoi(numbers[2])
        moves = append(moves, Move {
            qty: qty,
            from: from,
            to: to,
        })
    }

    pinput.moves = moves

    return pinput
}

func topCrates(crates []stack) string {
    var runes []rune
    for _, crate := range crates {
        runes = append(runes, crate[len(crate)-1])
    }
    return string(runes)
}

func p1(input PuzzleInput) string {
    for _, move := range input.moves {
        for i := 0; i < move.qty; i++ {
            if crate, value := input.crates[move.from-1].Pop(); crate != nil {
                input.crates[move.from-1] = crate
                input.crates[move.to-1] = input.crates[move.to-1].Push(value)
            }
        }
    }
    return topCrates(input.crates)
}

func p2(input PuzzleInput) string {
    for _, move := range input.moves {
        var temp stack
        for i := 0; i < move.qty; i++ {
            if crate, value := input.crates[move.from-1].Pop(); crate != nil {
                input.crates[move.from-1] = crate
                temp = temp.Push(value)
            }
        }
        for i := 0; i < move.qty; i++ {
            if crate, value := temp.Pop(); crate != nil {
                temp = crate
                input.crates[move.to-1] = input.crates[move.to-1].Push(value)
            }
        }
    }
    return topCrates(input.crates)
}

func main() {
    input := parseInput(readFile("./input.txt"))
    fmt.Println("p1:", p1(input))
    input = parseInput(readFile("./input.txt"))
    fmt.Println("p2:", p2(input))
}

