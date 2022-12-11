package main

import (
    "os"
    "fmt"
    "strings"
    "strconv"
    "sort"
    "math"
)

type Monkey struct {
    items []int
    op rune
    pow bool
    opvalue, test, t, f, ins int
}

func readFile(path string) string {
    buf, err := os.ReadFile(path)
    if err != nil {
        panic("Couldn't read file")
    }
    return string(buf)
}

func parseInput(input string) (out []Monkey) {
    monkeys := strings.Split(input, "\n\n")
    for _, monkey := range monkeys {
        m := strings.Split(monkey, "\n")
        m[1] = strings.Replace(m[1], "  Starting items: ", "", -1)
        m[2] = strings.Replace(m[2], "  Operation: new = old ", "", -1)
        m[3] = strings.Replace(m[3], "  Test: divisible by ", "", -1)
        m[4] = strings.Replace(m[4], "    If true: throw to monkey ", "", -1)
        m[5] = strings.Replace(m[5], "    If false: throw to monkey ", "", -1)

        mk := Monkey {}
        mk.ins = 0
        mk.pow = false

        strItems := strings.Split(m[1], ", ")
        for _, strItem := range strItems {
            item, _ := strconv.Atoi(strItem)
            mk.items = append(mk.items, int(item))
        }

        test, _ := strconv.Atoi(m[3])
        mk.test = int(test)

        t, _ := strconv.Atoi(m[4])
        f, _ := strconv.Atoi(m[5])
        mk.t = int(t)
        mk.f = int(f)

        mk.op = rune(m[2][0])
        stropvalue := strings.Split(m[2], " ")[1]
        if stropvalue != "old" {
            opvalue, _ := strconv.Atoi(stropvalue)
            mk.opvalue = int(opvalue)
        } else {
            mk.pow = true
        }

        out = append(out, mk)
    }
    return
}

func monkeySim(ms []Monkey, nrounds int, worrydiv int) int {
    var active []int

    mod := 1
    for _, m := range ms {
        mod *= m.test
    }

    for round := 0; round < nrounds; round++ {
        for i := range ms {
            for {
                if len(ms[i].items) == 0 {
                    break
                }

                w := ms[i].items[0]
                ms[i].items = ms[i].items[1:]
                ms[i].ins++

                if ms[i].pow {
                    w *= w
                } else if ms[i].op == '*' {
                    w *= ms[i].opvalue
                } else if ms[i].op == '+' {
                    w += ms[i].opvalue
                }

                if worrydiv != 0 {
                    w = int(math.Floor(float64(w) / float64(worrydiv)))
                }

                if w > mod {
                    w %= mod
                }

                test, t, f := ms[i].test, ms[i].t, ms[i].f

                if w % test == 0 {
                    ms[t].items = append(ms[t].items, w)
                } else {
                    ms[f].items = append(ms[f].items, w)
                }
            }
        }
    }

    for _, m := range ms {
        active = append(active, m.ins)
    }

    sort.Ints(active)
    return active[len(active)-1] * active[len(active)-2]
}

func p1(input string) int {
    mks := parseInput(input)
    return monkeySim(mks, 20, 3.0)
}

func p2(input string) int {
    mks := parseInput(input)
    return monkeySim(mks, 10000, 0.0)
}

func main() {
    input := strings.TrimSpace(readFile("./input.txt"))
    fmt.Println("p1:", p1(input))
    fmt.Println("p2:", p2(input))
}
