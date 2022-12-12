package main

import (
    "os"
    "fmt"
    "strings"
)

const INFINITY = 999999999

type Vertex struct {
    label, w int
}

type Graph struct {
    edges map[int][]Vertex
    vertices []Vertex
    nvertices int
}

func readFile(path string) string {
    buf, err := os.ReadFile(path)
    if err != nil {
        panic("Couldn't read file")
    }
    return string(buf)
}

func parseInput(input string) (out Graph, start Vertex, end Vertex) {
    lines := strings.Split(input, "\n")

    out.edges = make(map[int][]Vertex)
    out.nvertices = len(lines) * len(lines[0])
    cols := len(lines[0])
    rows := len(lines)

    makeVertex := func(x, y int) Vertex {
        w := rune(lines[y][x])
        if w == 'S' {
            w = 'a'
        }
        if w == 'E' {
            w = 'z'
        }
        return Vertex { x + y * cols, int(w - 'a') }
    }

    addEdge := func (x1, y1, x2, y2 int) {
        w1 := rune(lines[y1][x1])

        v1 := makeVertex(x1, y1)
        v2 := makeVertex(x2, y2)
        if v1.w + 1 >= v2.w {
            if w1 == 'S' {
                start = v1
            }
            if w1 == 'E' {
                end = v1
            }
            out.edges[v1.label] = append(out.edges[v1.label], v2)
        }
    }

    for y := 0; y < rows; y++ {
        for x := 0; x < cols; x++ {
            if x > 0 {
                addEdge(x, y, x-1, y)
            }
            if y > 0 {
                addEdge(x, y, x, y-1)
            }
            if x < len(lines[0]) - 1 {
                addEdge(x, y, x+1, y)
            }
            if y < len(lines) - 1 {
                addEdge(x, y, x, y+1)
            }
            out.vertices = append(out.vertices, makeVertex(x, y))
        }
    }

    return out, start, end
}

func dijkstra(start Vertex, graph Graph) []int {
    visited := make([]bool, graph.nvertices)
    distances := make([]int, graph.nvertices)

    for i := 0; i < len(visited); i++ {
        visited[i] = false
        distances[i] = INFINITY
    }
    distances[start.label] = 0

    for {
        nearest := -1
        min := INFINITY
        for i := 0; i < len(distances); i++ {
            if !visited[i] && distances[i] < min {
                min = distances[i]
                nearest = i
            }
        }

        if nearest == -1 {
            break
        }

        visited[nearest] = true

        for _, adj := range graph.edges[nearest] {
            if distances[adj.label] > distances[nearest] + 1 {
                distances[adj.label] = distances[nearest] + 1
            }
        }
    }

    return distances
}

func p1(input string) int {
    graph, start, end := parseInput(input)
    return dijkstra(start, graph)[end.label]
}

func p2(input string) int {
    min := INFINITY
    graph, _, end := parseInput(input)
    for _, v := range graph.vertices {
        if v.w == 0 {
            pathlen := dijkstra(v, graph)[end.label]
            if pathlen < min {
                min = pathlen
            }
        }
    }
    return min
}

func main() {
    input := strings.TrimSpace(readFile("./input.txt"))
    fmt.Println("p1:", p1(input))
    fmt.Println("p2:", p2(input))
}
