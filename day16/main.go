package main

import (
    "os"
    "fmt"
    "strings"
    "strconv"
)

const INF = 99999999

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}

func readFile(path string) string {
    buf, err := os.ReadFile(path)
    if err != nil {
        panic("Couldn't read file")
    }
    return string(buf)
}

type Valve struct {
    rate int
    leadsTo []int
}

func parseInput(input string) (valves []Valve) {
    var connections [][]string
    var rates []int
    names := make(map[string]int)

    input = strings.Replace(input, "Valve ", "", -1)
    input = strings.Replace(input, " tunnels lead to valves ", "", -1)
    input = strings.Replace(input, " tunnel leads to valve ", "", -1)
    input = strings.Replace(input, " has flow rate=", "", -1)

    for i, line := range strings.Split(input, "\n") {
        name := line[0:2]
        sci := strings.Index(line, ";")
        strrate := line[2:sci]
        rate, _ := strconv.Atoi(strrate)
        leadsTo := strings.Split(line[sci+1:], ", ")

        names[name] = i
        connections = append(connections, leadsTo)
        rates = append(rates, rate)
    }

    for i, strLeadsTo := range connections {
        var leadsTo []int
        for _, l := range strLeadsTo {
            leadsTo = append(leadsTo, names[l])
        }
        valve := Valve {
            rates[i],
            leadsTo,
        }
        valves = append(valves, valve)
    }

    return
}

func initialize(valves []Valve) (dist [][]int, adj [][]int) {
    nvalues := len(valves)
    dist = make([][]int, nvalues)
    adj = make([][]int, nvalues)

    for i := range valves {
        dist[i] = make([]int, nvalues)
        adj[i] = make([]int, nvalues)
        for j := range valves {
            if i == j {
                dist[i][j] = 0
            } else {
                dist[i][j] = INF
            }
            adj[i][j] = -1
        }
    }

    for i, valve := range valves {
        for _, j := range valve.leadsTo {
            dist[i][j] = 1
            adj[i][j] = j
        }
    }

    return 
}

func findPath(adj [][]int, i, j int) (path []int) {
    if adj[i][j] == -1 {
        return path
    }

    path = append(path, i)
    for i != j {
        i = adj[i][j]
        path = append(path, i)
    }
    return path
}

func floydWarshall(dist, adj [][]int) ([][]int, [][]int) {
    nvalues := len(dist)

    for k := 0; k < nvalues; k++ {
        for i := 0; i < nvalues; i++ {
            for j := 0; j < nvalues; j++ {
                if dist[i][k] == INF || dist[k][j] == INF {
                    continue
                }

                if dist[i][j] > dist[i][k] + dist[k][j] {
                    dist[i][j] = dist[i][k] + dist[k][j]
                    adj[i][j] = adj[i][k]
                }
            }
        }
    }

    return dist, adj
}

func p1(input string) int {
    valves := parseInput(input)

    dist, adj := initialize(valves)
    dist, adj = floydWarshall(dist, adj)

    return 0
}

func main() {
    input := strings.TrimSpace(readFile("./input.txt"))
    fmt.Println(p1(input))
}

