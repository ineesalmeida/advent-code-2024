package main

import (
	"advent-code/aoc2024/utils"
	"fmt"
	"log"
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

func part1(lines []string) any {
	result := 0
	for i, line := range lines {
		for j, c := range line {
			if c == rune('9') {
				zeros := make(map[[2]int]bool)
				cache := make(map[[2]int]int)
				walk(lines, i, j, cache, zeros)
				// fmt.Println("Result for this node", i, j, res, len(zeros))
				result += len(zeros)
			}
		}
	}
	return result
}

var nextItemMap = map[string]string{
	"9": "8",
	"8": "7",
	"7": "6",
	"6": "5",
	"5": "4",
	"4": "3",
	"3": "2",
	"2": "1",
	"1": "0",
}
var lastItem string = "0"

func walk(lines []string, i int, j int, cache map[[2]int]int, zeros map[[2]int]bool) int {
	curr := string(lines[i][j])
	// fmt.Println("Checking ", i, j, curr)

	// If we reach the last item, then return 1 for the counter
	if curr == lastItem {
		// fmt.Println("FOUND A ZERO AT ", i, j)
		cache[[2]int{i, j}] = 1
		zeros[[2]int{i, j}] = true
		return 1
	}
	cachedRes, ok := cache[[2]int{i, j}]
	if ok {
		// fmt.Println("FOUND A CACHE AT ", i, j, cachedRes)
		return cachedRes
	}

	totalRes := 0

	if i > 0 && string(lines[i-1][j]) == nextItemMap[curr] {
		res := walk(lines, i-1, j, cache, zeros)
		if res != 0 {
			cache[[2]int{i - 1, j}] = res
			totalRes += res
			// fmt.Println("Walked ", i-1, j, res, totalRes, cache)
		}
	}
	if i < len(lines)-1 && string(lines[i+1][j]) == nextItemMap[curr] {
		res := walk(lines, i+1, j, cache, zeros)
		if res != 0 {
			cache[[2]int{i + 1, j}] = res
			totalRes += res
			// fmt.Println("Walked ", i+1, j, res, totalRes, cache)
		}
	}
	if j > 0 && string(lines[i][j-1]) == nextItemMap[curr] {
		res := walk(lines, i, j-1, cache, zeros)
		if res != 0 {
			cache[[2]int{i, j - 1}] = res
			totalRes += res
			// fmt.Println("Walked ", i, j-1, res, totalRes, cache)
		}
	}
	if j < len(lines[0])-1 && string(lines[i][j+1]) == nextItemMap[curr] {
		res := walk(lines, i, j+1, cache, zeros)
		if res != 0 {
			cache[[2]int{i, j + 1}] = res
			totalRes += res
			// fmt.Println("Walked ", i, j+1, res, totalRes, cache)
		}
	}

	return totalRes
}

func part2(lines []string) any {
	result := 0
	cache := make(map[[2]int]int)
	zeros := make(map[[2]int]bool)
	for i, line := range lines {
		for j, c := range line {
			if c == rune('9') {
				res := walk(lines, i, j, cache, zeros)
				// fmt.Println("Result for this node", i, j, res, len(zeros))
				result += res
			}
		}
	}
	return result
}
