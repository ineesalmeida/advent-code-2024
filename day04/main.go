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
	// Search for XMAS and SAMX horizontally, verically and 2 diagonally
	result := 0

	n := len(lines)
	m := len(lines[0])

	vlines := make([]string, m)
	d1lines := make([]string, n+m)
	d2lines := make([]string, n+m)

	for i, line := range lines {
		// horizontally
		result += countStrincOccurrences(line, "XMAS")
		result += countStrincOccurrences(line, "SAMX")

		// construct vertical lines
		for j, char := range line {
			vlines[j] += string(char)

			// construct diagonal lines
			if j+i < len(d1lines) {
				d1lines[j+i] += string(char)
			}
			if n-j+i >= 0 {
				d2lines[n-j+i] += string(char)
			}
		}
	}

	// vertically
	for _, line := range vlines {
		result += countStrincOccurrences(line, "XMAS")
		result += countStrincOccurrences(line, "SAMX")
	}

	// diagonally 1
	for _, line := range d1lines {
		result += countStrincOccurrences(line, "XMAS")
		result += countStrincOccurrences(line, "SAMX")
	}

	// diagonally 2
	for _, line := range d2lines {
		result += countStrincOccurrences(line, "XMAS")
		result += countStrincOccurrences(line, "SAMX")
	}

	return result
}

func countStrincOccurrences(line string, word string) int {
	return len(strings.Split(line, word)) - 1
}

func part2(lines []string) any {
	// change of plans, search for all the As and then check if there are M and S around it
	result := 0

	A := "A"

	adjacent1 := ""
	adjacent2 := ""
	for i, line := range lines[1 : len(lines)-1] {
		i += 1
		for j, char := range line[1 : len(line)-1] {
			j += 1
			if string(char) == A {
				adjacent1 = string(lines[i-1][j-1]) + A + string(lines[i+1][j+1])
				adjacent2 = string(lines[i-1][j+1]) + A + string(lines[i+1][j-1])

				fmt.Println(adjacent1, adjacent2)

				if (adjacent1 == "MAS" || adjacent1 == "SAM") && (adjacent2 == "MAS" || adjacent2 == "SAM") {
					result += 1
				}
			}
		}
	}
	return result
}
