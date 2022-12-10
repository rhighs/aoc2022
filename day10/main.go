package main

import (
    "fmt"
    "strings"
    "os"
    "strconv"
)

func readFile(path string) string {
    buf, err := os.ReadFile(path)
    if err != nil {
        panic("Couldn't read file")
    }
    return string(buf)
}

func abs(n int) int {
	if n >= 0 {
		return n
	}
	return -n
}

func exec(input string, uptoCycles int, postCycle func(int, int)) {
    lines := strings.Split(input, "\n")
    X := 1
    cycles := 0
    cycle := func() {
        cycles++
        postCycle(X, cycles)
    }
    for _, line := range lines {
        instr := strings.Split(line, " ")
        switch len(instr) {
            case 2:
                cycle()
                cycle()
                v, _ := strconv.Atoi(instr[1])
                X += v
            case 1:
                cycle()
        }

        if cycles > uptoCycles {
            break
        }
    }
}

func p1(input string) int {
    out := 0
    cycle := func(X, cycles int) {
        if (cycles - 20) % 40 == 0 {
            out += X * cycles
        }
    }
    exec(input, 220, cycle)
    return out
}

func p2(input string) string {
    strout := ""
    out := make([][]rune, 6)
    for i := range out {
        out[i] = []rune("........................................")
    }
    cycle := func(X, cycles int) {
        cycles--
        y := cycles / 40
        x := cycles % 40
        if abs(X-x) <= 1 {
            out[y][x] = '#'
        }
    }
    exec(input, 240, cycle)
    for _, r := range out {
        strout += string(r) + "\n"
    }
    return strout
}

func main() {
    input := strings.TrimSpace(readFile("./input.txt"))
    fmt.Println("p1:", p1(input))
    fmt.Println("p2:")
    fmt.Println(p2(input))
}
