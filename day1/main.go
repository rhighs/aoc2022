package main

import (
    "fmt"
    "os"
    "strconv"
    "strings"
    "sort"
)

func readFile(path string) string {
    buf, err := os.ReadFile(path)
    if err != nil {
        panic("Couldn't read file")
    }
    return string(buf)
}

func listSum(someList []int) int {
    sum := 0
    for _, item := range someList {
        sum += item
    }
    return sum
}

func findMax(someList []int) int {
    m := someList[0]
    for _, value := range someList {
        if value > m {
            m = value
        }
    }
    return m
}

func strIntListsToSumsSorted(strIntLists []string) []int {
    var intList []int
    for _, strIntList := range strIntLists {
        var ints []int 
        for _, strValue := range strings.Split(strIntList, "\n") {
            v, _ := strconv.Atoi(strValue)
            ints = append(ints, v)
        }
        intList = append(intList, listSum(ints))
    }
    sort.Ints(intList)
    return intList
}

func p1(input string) int {
    sums := strIntListsToSumsSorted(strings.Split(input, "\n\n"))
    return sums[len(sums)-1]
}

func p2(input string) int {
    sums := strIntListsToSumsSorted(strings.Split(input, "\n\n"))
    return sums[len(sums)-1] + sums[len(sums)-2] + sums[len(sums)-3]
}

func main() {
    input := readFile("./input.txt")
    fmt.Println("p1:", p1(input))
    fmt.Println("p2:", p2(input))
}
