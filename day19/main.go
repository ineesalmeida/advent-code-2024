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

func part1(lines []string) any {
	parseInput(lines)
	// fmt.Println(availableTowels)
	// fmt.Println(patterns)

	result := 0
	for _, pattern := range patterns {
		if checkPatternPossible(pattern) {
			result += 1
		}
	}

	return result
}

func checkPatternPossible(pattern string) bool {
	// check if a pattern is possible
	// brwrr --> b, br...

	_, ok := availableTowels[pattern]
	if ok {
		// fmt.Println("Found in availale towels", pattern)
		// fmt.Println("CACHE", cache)
		cache[pattern] = true
		return true
	}

	val, ok := cache[pattern]
	if ok {
		// fmt.Println("Found in cache", pattern)
		// fmt.Println("CACHE", cache)
		return val
	}

	for i := range pattern {
		_, ok := availableTowels[pattern[:i+1]]
		if ok {
			if checkPatternPossible(pattern[i+1:]) {
				cache[pattern] = true
				// fmt.Println("Found that is possible ", pattern)
				// fmt.Println("CACHE", cache)
				return true
			}
		} else {
			cache[pattern] = false
		}
	}
	return false
}

var availableTowels = make(map[string]bool)
var patterns = make([]string, 0)
var cache = make(map[string]bool)

func parseInput(lines []string) {
	towels := strings.Split(lines[0], ", ")
	for _, towel := range towels {
		availableTowels[towel] = true
	}

	for _, towel := range lines[2:] {
		patterns = append(patterns, towel)
	}
}

func part2(lines []string) any {
	return nil
}
