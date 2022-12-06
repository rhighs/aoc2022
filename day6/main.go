package main

import (
    "fmt"
    "os"
    "sort"
    "strings"
)

func readFile(path string) string {
    buf, err := os.ReadFile(path)
    if err != nil {
        panic("Couldn't read file")
    }
    return string(buf)
}

func checkDistinct(someString string) bool {
    chars := strings.Split(someString, "")
	sort.Strings(chars)
	for i := 1; i < len(chars); i++ {
		if chars[i] == chars[i-1] {
			return false
		}
	}
	return true
}

func findPackeStart(input string, packetLen int) int {
    for i := 0; i < len(input) - packetLen; i++ {
        if checkDistinct(input[i:i+packetLen]) {
            return i+packetLen
        }
    }
    return -1
}

func main() {
    input := readFile("./input.txt")
    fmt.Println("p1:", findPackeStart(input, 4))
    fmt.Println("p2:", findPackeStart(input, 14))
}
