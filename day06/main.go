package main

import (
	"advent-code/aoc2024/utils"
	"fmt"
	"log"
	"strings"
	"time"
)

func main() {
	lines, err := utils.FileToLines("input.txt")
	if err != nil {
		log.Fatalf("[!] %s\n", err)
	}

	p1start := time.Now()
	part1Answer := part1(lines)
	p1duration := time.Since(p1start)

	p2start := time.Now()
	part2Answer := part2(lines)
	p2duration := time.Since(p2start)

	fmt.Printf("[+] day %s\n> part 1: %d (%s)\n> part 2: %d (%s)\n",
		"02",
		part1Answer, p1duration,
		part2Answer, p2duration,
	)
}

// type Node struct {
// 	Up    *Node
// 	Down  *Node
// 	Left  *Node
// 	Right *Node
// }

// func generateNodeMap(lines []string) (Node, string) {
// 	nodes := make([][]Node, len(lines))
// 	for i, line := range lines {
// 		for j, c := range line {
// 			if string(c) == "#" {
// 				continue
// 			}
// 			if string(c) == "." {
// 				nodes[i][j] = Node{
// 					Up:
// 				}
// 			} else {
// 				// start node

// 			}
// 		}
// 	}
// }

func part1(lines []string) any {
	initialX, initialY, initialDir := getInitialPosition(lines)
	return walk(lines, initialX, initialY, initialDir)
}

func getInitialPosition(lines []string) (int, int, string) {
	dirs := []string{"<", "^", ">", "v"}

	var j int
	for i, line := range lines {
		for _, c := range dirs {
			j = strings.Index(line, c)
			if j != -1 {
				return j, i, c
			}
		}
	}
	return -1, -1, ""
}

func isBlocked(lines []string, x int, y int) bool {
	return string(lines[y][x]) == "#"
}

func isEnd(lines []string, x int, y int) bool {
	return x < 0 || y < 0 || y >= len(lines) || x >= len(lines[0])
}

func ChangeDir(x int, y int, dir string) (int, int, string) {
	switch dir {
	case "^":
		y += 1
		dir = ">"
	case "v":
		y -= 1
		dir = "<"
	case "<":
		x += 1
		dir = "^"
	case ">":
		x -= 1
		dir = "v"
	}
	return x, y, dir
}

func walk(lines []string, x int, y int, dir string) int {
	haveSeenBefore := make(map[string]string)
	haveSeenBefore[string(x)+string(y)] = dir

	counter := 1
	for true {
		if dir == "^" {
			y -= 1
		} else if dir == "v" {
			y += 1
		} else if dir == "<" {
			x -= 1
		} else {
			x += 1
		}
		if isEnd(lines, x, y) {
			break
		}
		if isBlocked(lines, x, y) {
			x, y, dir = ChangeDir(x, y, dir)
		} else {
			val, ok := haveSeenBefore[string(x)+string(y)]
			if !ok {
				counter += 1
			}
			if val == dir {
				// we have ben here before on the same direction ===> loop!
				return -1
			}
			haveSeenBefore[string(x)+string(y)] = dir
			// fmt.Println(x, y, dir)
		}
	}
	return counter
}

func replaceAtIndex(in string, r rune, i int) string {
	out := []rune(in)
	out[i] = r
	return string(out)
}

func part2(lines []string) any {
	initialX, initialY, initialDir := getInitialPosition(lines)
	counter := 0
	tmp := make([]string, len(lines))
	for i, line := range lines {
		for j, c := range line {
			if string(c) == "#" {
				continue
			}
			copy(tmp, lines)
			tmp[i] = replaceAtIndex(tmp[i], rune('#'), j)
			if walk(tmp, initialX, initialY, initialDir) == -1 {
				counter += 1
			}
		}

	}
	return counter
}
