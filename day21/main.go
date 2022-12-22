package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type MonkeyType uint8

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

const (
	Waiting = iota
	Yelling
)

type Monkey struct {
	name     string
	op       rune
	value    int
	operands []*Monkey
	t        MonkeyType
}

func readFile(path string) string {
	buf, err := os.ReadFile(path)
	if err != nil {
		panic("Couldn't read file")
	}
	return string(buf)
}

func parseMonkey(name string, monkeys map[string]string) *Monkey {
	m := new(Monkey)
	m.name = name
	strexpr := monkeys[name]
	v, err := strconv.Atoi(strexpr)
	if err != nil {
		expr := strings.Split(strexpr, " ")
		m.op = rune(expr[1][0])
		m.operands = append(m.operands, parseMonkey(expr[0], monkeys))
		m.operands = append(m.operands, parseMonkey(expr[2], monkeys))
		m.t = Waiting
	} else {
		m.t = Yelling
		m.value = v
	}
	return m
}

func parseInput(input string) *Monkey {
	monkeys := make(map[string]string)
	for _, line := range strings.Split(input, "\n") {
		line = strings.TrimSpace(line)
		yell := strings.Split(line, ":")
		yell[1] = strings.TrimSpace(yell[1])
		monkeys[yell[0]] = yell[1]
	}
	return parseMonkey("root", monkeys)
}

func getMonkeyWithName(root *Monkey, name string) (bool, *Monkey) {
	if root.name == "humn" {
		return true, root
	}

	for _, m := range root.operands {
		if b, mm := getMonkeyWithName(m, name); b {
			return true, mm
		}
	}

	return false, nil
}

func evalMonkey(monkey *Monkey) int {
	if monkey.t == Yelling {
		return monkey.value
	}

	switch monkey.op {
	case '+':
		return evalMonkey(monkey.operands[0]) + evalMonkey(monkey.operands[1])
	case '-':
		return evalMonkey(monkey.operands[0]) - evalMonkey(monkey.operands[1])
	case '*':
		return evalMonkey(monkey.operands[0]) * evalMonkey(monkey.operands[1])
	case '/':
		return evalMonkey(monkey.operands[0]) / evalMonkey(monkey.operands[1])
	}

	return 0 // never reached
}

func p1(input string) int {
	root := parseInput(input)
	return evalMonkey(root)
}

func p2(input string) int {
	root := parseInput(input)
	_, human := getMonkeyWithName(root, "humn")
	v1 := evalMonkey(root.operands[0])
	v2 := evalMonkey(root.operands[1])
	distance := abs(v1 - v2)
	mul := 1000000000
	for i := 0; v1 != v2; i++ {
		if abs(v1-v2) < distance/10 && mul/10 != 0 {
			distance = abs(v1 - v2)
			mul /= 10
			continue
		}

		human.value += mul
		v1 = evalMonkey(root.operands[0])
		v2 = evalMonkey(root.operands[1])
		if v1 == v2 {
			break
		}
	}

	return human.value
}

func main() {
	input := strings.TrimSpace(readFile("./input.txt"))
	fmt.Println("p1:", p1(input))
	fmt.Println("p2:", p2(input))
}
