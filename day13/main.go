package main

import (
    "os"
    "fmt"
    "strings"
    "strconv"
    "sort"
)

type NodeType int
const (
    LIST NodeType = iota
    INT
)

type Node struct {
    t NodeType
    childs []*Node
    value int
}

func readFile(path string) string {
    buf, err := os.ReadFile(path)
    if err != nil {
        panic("Couldn't read file")
    }
    return string(buf)
}

func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}

func matchingBracket(line string, bpos int) int {
    open := 0
    for i := bpos; i < len(line); i++ {
        if line[i] == '[' {
            open++
        }
        if line[i] == ']' {
            open--
        }
        if open == 0 {
            return i
        }
    }
    return len(line)-1
}

func parseTree(line string) *Node {
    node := new(Node)

    for i := 0; i < len(line); i++ {
        if line[i] == '[' {
            resumeat := matchingBracket(line, i)
            child := parseTree(line[i+1:resumeat])
            node.childs = append(node.childs, child)
            i = resumeat
        }

        strvalue := ""
        for ; i < len(line) && line[i] != ',' && line[i] != '[' && line[i] != ']'; i++ {
            strvalue += string(line[i])
        }

        if strvalue != "" {
            intnode := new(Node)
            value, _ := strconv.Atoi(strvalue)
            intnode.t = INT
            intnode.value = value
            intnode.childs = append(intnode.childs, nil)
            node.childs = append(node.childs, intnode)
        }
    }

    node.t = LIST
    return node
}

func compareNodes(n1, n2 *Node) (res int) {
    for i := 0; i < min(len(n1.childs), len(n2.childs)) && res == 0; i++ {
        c1, c2 := n1.childs[i], n2.childs[i]
        switch {
        case c1.t == INT && c2.t == INT:
            res = c2.value - c1.value
        case c1.t == LIST && c2.t == INT:
            node := new(Node)
            node.childs = append(node.childs, c2)
            res = compareNodes(c1, node)
        case c1.t == INT && c2.t == LIST:
            node := new(Node)
            node.childs = append(node.childs, c1)
            res = compareNodes(node, c2)
        default:
            res = compareNodes(c1, c2)
        }
    }

    if res != 0 {
        return res
    }

    return len(n2.childs) - len(n1.childs)
}

func p1(input string) int {
    pairs := strings.Split(input, "\n\n")
    tot := 0
    for i, pair := range pairs {
        p := strings.Split(pair, "\n")
        n1 := parseTree(p[0])
        n2 := parseTree(p[1])
        if compareNodes(n1, n2) > 0 {
            tot += i + 1
        }
    }
    return tot
}

type OrderedPacket struct {
    node *Node
    id int
}

func p2(input string) int {
    packets := strings.Replace(input, "\n\n", "\n", -1)
    strpackets := strings.Split(packets, "\n")
    strpackets = append(strpackets, "[[2]]")
    strpackets = append(strpackets, "[[6]]")

    var ordpackets []OrderedPacket
    for i, strpacket := range strpackets {
        orderedPacket := OrderedPacket {
            parseTree(strpacket),
            i,
        }
        ordpackets = append(ordpackets, orderedPacket)
    }

    sort.Slice(ordpackets, func(i, j int) bool {
        return compareNodes(ordpackets[i].node, ordpackets[j].node) > 0
    })

    tot := 1
    for i, orderedPacket := range ordpackets {
        if len(strpackets) - orderedPacket.id <= 2 {
            tot *= i + 1
        }
    }

    return tot
}

func main() {
    input := strings.TrimSpace(readFile("./input.txt"))
    fmt.Println(p1(input))
    fmt.Println(p2(input))
}
