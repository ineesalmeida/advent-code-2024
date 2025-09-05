package main

import (
	"advent-code/aoc2024/utils"
	"fmt"
	"log"
	"math"
	"time"
)

func main() {
	lines, err := utils.FileToLines("test-input.txt")
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

var nodes = make(map[[2]int]int)

func part1(lines []string) any {
	root, _ := parseInput(lines)
	// nodes[root] = heuristics(root, target, lines)
	return nodes[root]
}

func parseInput(lines []string) ([2]int, [2]int) {
	root := [2]int{}
	target := [2]int{}
	for i, line := range lines {
		for j, node := range line {
			if node == 'S' {
				root = [2]int{i, j}
				nodes[[2]int{i, j}] = 0
				continue
			}
			if node == 'E' {
				target = [2]int{i, j}
			}
			nodes[[2]int{i, j}] = math.MaxInt64
		}
	}
	return root, target
}

func walk(root [2]int, target [2]int, lines []string) int {
	result := 0

	return result
}

func part2(lines []string) any {
	return nil
}
