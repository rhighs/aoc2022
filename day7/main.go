package main

import (
    "fmt"
    "os"
    "strings"
    "strconv"
)

type Dir struct {
    parent *Dir
    subdirs []*Dir
    name string
    files []int
}

func readFile(path string) string {
    buf, err := os.ReadFile(path)
    if err != nil {
        panic("Couldn't read file")
    }
    return string(buf)
}

// Again another horrible input parser, this aoc is getting funnier each day
func parseInput(input string) map[string]*Dir {
    fs := make(map[string]*Dir)
    lines := strings.Split(input, "\n")

    currentDir := "/"
    rootDir := new(Dir)
    rootDir.name = "/"
    rootDir.parent = rootDir
    fs["/"] = rootDir

    for _, line := range lines {
        if line[0:4] == "$ cd" {
            cd := line[5:]
            if cd == ".." {
                currentDir = fs[currentDir].parent.name
            } else if cd != ".." && cd != "/" {
                currentDir += "/" + cd
            } else if cd == "/" {
                currentDir = "/"
            }
        } else  if _, err := strconv.Atoi(string(line[0])); err == nil {
            fileSize, _ := strconv.Atoi(strings.Split(line, " ")[0])
            fs[currentDir].files = append(fs[currentDir].files, fileSize)
        } else if line[0] == 'd' {
            dirName := fs[currentDir].name + "/" + strings.Split(line, " ")[1]
            _, ok := fs[dirName];
            if !ok {
                fs[dirName] = new(Dir)
                fs[dirName].name = dirName
                fs[dirName].parent = fs[currentDir]
            }
            fs[currentDir].subdirs = append(fs[currentDir].subdirs, fs[dirName])
        }
    }

    return fs
}

func recurseDirSize(root *Dir) int {
    size := 0
    for _, dir := range root.subdirs {
        size += recurseDirSize(dir)
    }
    for _, file := range root.files {
        size += file
    }
    return size
}

func main() {
    input := strings.TrimSuffix(readFile("./input.txt"), "\n")
    fs := parseInput(input)

    funkyDirs := 0
    total := 70000000
    needed := 30000000
    var sizes []int

    for key, _ := range fs {
        size := recurseDirSize(fs[key])
        sizes = append(sizes, size)
        if size <= 100000 {
            funkyDirs += size
        }
    }

    toRemove := recurseDirSize(fs["/"]) - (total - needed)
    minValid := 99999999
    for _, size := range sizes {
        if size > toRemove && size < minValid {
            minValid = size
        }
    }

    fmt.Println("p1:", funkyDirs)
    fmt.Println("p2:", minValid)
}
