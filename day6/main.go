package main

import (
    "fmt"
    "os"
)

func readFile(path string) string {
    buf, err := os.ReadFile(path)
    if err != nil {
        panic("Couldn't read file")
    }
    return string(buf)
}

func checkDistinct(someString string) bool {
    charCounts := make(map[rune]int)
    for _, c := range someString {
        charCounts[c]++
    }
    for _, count := range charCounts {
        if count > 1 {
            return false
        }
    }
    return true
}

func FindPacketStart(input string, packetLen int) int {
    for i := 0; i < len(input) - packetLen; i++ {
        if checkDistinct(input[i:i+packetLen]) {
            return i+packetLen
        }
    }
    return -1
}

func main() {
    input := readFile("./input.txt")
    fmt.Println("p1:", FindPacketStart(input, 4))
    fmt.Println("p2:", FindPacketStart(input, 14))
}
