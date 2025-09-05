package main

import (
	"advent-code/aoc2024/utils"
	"fmt"
	"log"
	"reflect"
	"strconv"
	"strings"
	"time"
)

func main() {

	// lines, err := utils.FileToLines("test-input.txt")
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
	orderMap, manuals := parseLines(lines)
	result := 0
	for _, manual := range manuals {
		if checkOrder(manual, orderMap) {
			result += getMidPage(manual)
		}
	}

	return result
}

func parseLines(lines []string) (map[string]int, [][]string) {
	flag := false
	orderIndexes := make(map[string]int)
	manuals := make([][]string, 0, len(lines))
	for _, line := range lines {
		if line == "" {
			flag = true
		} else if !flag {
			ps := strings.Split(line, "|")
			orderIndexes[ps[0]], _ = orderIndexes[ps[0]]
			orderIndexes[ps[1]], _ = orderIndexes[ps[1]]

			if orderIndexes[ps[1]] > orderIndexes[ps[0]] {
				continue
			} else if orderIndexes[ps[1]] == 0 {
				orderIndexes[ps[1]] = orderIndexes[ps[0]] + 1
			} else {
				orderIndexes[ps[1]]
			}

		} else {
			manuals = append(manuals, strings.Split(line, ","))
		}
	}
	return orderIndexes, manuals
}

func checkOrder(manual []string, orderMap map[string][]string) bool {
	// keep track of all pages seen before
	// if a page was supposed to come before another one that has been seen, then return false
	seenBefore := make(map[string]bool)
	for _, page := range manual {
		val, ok := orderMap[page]
		if ok {
			for _, nextPage := range val {
				if seenBefore[nextPage] {
					return false
				}
			}
		}
		seenBefore[page] = true
	}
	return true
}

func getMidPage(manual []string) int {
	midPage := manual[(len(manual)-1)/2]
	midPageInt, _ := strconv.Atoi(midPage)
	return midPageInt
}

func part2(lines []string) any {
	orderMap, manuals := parseLines(lines)
	result := 0
	for _, manual := range manuals {
		if !checkOrder(manual, orderMap) {
			// correct the manual
			correctedManual := correctOrder(manual, orderMap)
			result += getMidPage(correctedManual)
		}
	}

	return result
}

func correctOrder(manual []string, orderMap map[string][]string) []string {
	// keep track of all pages seen before
	// if a page was supposed to come before another one that has been seen, then return false
	seenBefore := make(map[string]int)
	for i, page := range manual {
		val, ok := orderMap[page]
		if ok {
			for _, nextPage := range val {
				index, ok := seenBefore[nextPage]
				if ok {
					swapF := reflect.Swapper(manual)
					swapF(i, index)
					if checkOrder(manual, orderMap) {
						return manual
					} else {
						return correctOrder(manual, orderMap)
					}
				}
			}
		}
		seenBefore[page] = i
	}
	return manual
}
