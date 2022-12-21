package main

import (
	"container/ring"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

func readFile(path string) string {
	buf, err := os.ReadFile(path)
	if err != nil {
		panic("Couldn't read file")
	}
	return string(buf)
}

func parseInput(input string) (out []int) {
	for _, strn := range strings.Split(input, "\n") {
		strn = strings.TrimSpace(strn)
		n, _ := strconv.Atoi(strn)
		out = append(out, n)
	}
	return
}

func intoRing(arr []int) *ring.Ring {
	r := ring.New(len(arr))
	for _, v := range arr {
		r.Value = v
		r = r.Next()
	}
	return r
}

func ringFind(r *ring.Ring, value any) *ring.Ring {
	v, ok := value.(int)
	if ok {
		lr := r.Len()
		for i := 0; i < lr; i++ {
			if r.Value == v {
				return r
			}
			r = r.Next()
		}
	}
	return nil
}

func ringFindAll(r *ring.Ring, value any) (out []*ring.Ring) {
	v, ok := value.(int)
	if ok {
		lr := r.Len()
		for i := 0; i < lr; i++ {
			if r.Value == v {
				out = append(out, r)
			}
			r = r.Next()
		}
	}
    return
}

func printRing(r *ring.Ring) {
	lr := r.Len()
	for i := 0; i < lr-1; i++ {
		fmt.Print(r.Value, ",")
		r = r.Next()
	}
	fmt.Println(r.Value)
}

func p1(input string) int {
	file := parseInput(input)
	decrypted := intoRing(file)

	for _, n := range file {
		subr := ringFind(decrypted, n).Prev()
		removed := subr.Unlink(1)
		subr.Move(n).Link(removed)
	}

	tot := 0
	subr := ringFind(decrypted, 0)
	for i := 0; i < 3; i++ {
		subr = subr.Move(1000)
		v, _ := subr.Value.(int)
		tot += v
	}

	return tot
}

func main() {
	input := strings.TrimSpace(readFile("./input.txt"))
	fmt.Println("p1:", p1(input))
}
