package main

import (
	"container/ring"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type RingKey struct {
	i, n int
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

func intoHashedRings(arr []int) (map[RingKey]*ring.Ring, RingKey) {
	out := make(map[RingKey]*ring.Ring)
	r := intoRing(arr)
	zk := RingKey{}
	for i := 0; i < r.Len(); i++ {
		n := r.Value.(int)
		if n == 0 {
			zk = RingKey{i, n}
		}
		out[RingKey{i, n}] = r
		r = r.Next()
	}
	return out, zk
}

func decrypt(file []int, mixes int, decryptionKey int) int {
    if decryptionKey != 0 {
	    for i := range file {
		    file[i] *= decryptionKey
	    }
    } 

	hashedRings, zk := intoHashedRings(file)
	l, hl := len(file)-1, len(file)/2

	for k := 0; k < mixes; k++ {
		for i, n := range file {
			subr := hashedRings[RingKey{i, n}].Prev()
			removed := subr.Unlink(1)

			//https://github.com/python/cpython/blob/85dd6cb6df996b1197266d1a50ecc9187a91e481/Modules/_collectionsmodule.c#L764
			if (n > hl) || (n < -hl) {
				n %= l
				if n > hl {
					n -= l
				} else if n < -hl {
					n += l
				}
			}

			subr.Move(n).Link(removed)
		}
	}

	s := 0
	subr := hashedRings[zk]
	for i := 0; i < 3; i++ {
		subr = subr.Move(1000)
		v, _ := subr.Value.(int)
		s += v
	}

	return s
}

func main() {
    file := parseInput(strings.TrimSpace(readFile("./input.txt")))
	fmt.Println("p1:", decrypt(file, 1, 0))
	fmt.Println("p2:", decrypt(file, 10, 811589153))
}
