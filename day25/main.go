package main

import (
	"fmt"
	"math"
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

func parseInput(input string) []string {
	s := strings.Split(input, "\n")
	for i := range s {
		s[i] = strings.TrimSpace(s[i])
	}
	return s
}

func stringReverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func snafuToDec(snafu string) int {
	n := 0
	for i, s := range snafu {
		switch s {
		case '-':
			n += -1 * int(math.Pow(float64(5), float64(len(snafu)-i-1)))
		case '=':
			n += -2 * int(math.Pow(float64(5), float64(len(snafu)-i-1)))
		case '0':
			n += 0 * int(math.Pow(float64(5), float64(len(snafu)-i-1)))
		case '1':
			n += 1 * int(math.Pow(float64(5), float64(len(snafu)-i-1)))
		case '2':
			n += 2 * int(math.Pow(float64(5), float64(len(snafu)-i-1)))
		}
	}
	return n
}

func decToSnafu(dec int) (out string) {
	var rets []int
	borrow := 0

	for dec != 0 {
		rem := dec % 5
		rem += borrow
		dec /= 5

		if rem >= 3 {
			rem -= 5
			rets = append(rets, rem)
			borrow = 1
		} else {
			rets = append(rets, rem)
			borrow = 0
		}
	}

	for _, r := range rets {
		switch r {
		case -2:
			out += "="
		case -1:
			out += "-"
		case 0:
			out += "0"
		case 1:
			out += "1"
		case 2:
			out += "2"
		}
	}

	return stringReverse(out)
}

func main() {
	input := strings.TrimSpace(readFile("./input.txt"))
	tot := 0
	for _, snafu := range parseInput(input) {
		tot += snafuToDec(snafu)
	}
	fmt.Println(decToSnafu(tot))
}
