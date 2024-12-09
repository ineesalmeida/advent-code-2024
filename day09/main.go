package main

import (
	"advent-code/aoc2024/utils"
	"fmt"
	"log"
	"slices"
	"strconv"
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

func checkSum(vals []int) int {
	sum := 0
	for i, val := range vals {
		if val != -1 {
			sum += i * val
		}
	}
	return sum
}

func part1(lines []string) any {
	// single line with ints
	line := make([]int, 0, len(lines[0])*10)
	counter := 0
	for i, val := range lines[0] {
		j, _ := strconv.Atoi(string(val))
		if i%2 == 0 {
			line = append(line, slices.Repeat([]int{counter}, j)...)
			counter += 1
		} else {
			line = append(line, slices.Repeat([]int{-1}, j)...)
		}
	}
	// fmt.Println(line)
	left := 0
	right := len(line) - 1
	for left < right {
		if line[left] == -1 {
			line[left], line[right] = line[right], line[left]
			right -= 1
			for line[right] == -1 {
				right -= 1
			}
		}
		left += 1
	}

	return checkSum(line)
}

func part2(lines []string) any {
	// single line with ints
	line := make([]int, 0, len(lines[0])*10)
	freeSpace := make([][2]int, 0, len(lines[0])*10)
	counter := 0
	for i, val := range lines[0] {
		j, _ := strconv.Atoi(string(val))
		if i%2 == 0 {
			line = append(line, slices.Repeat([]int{counter}, j)...)
			counter += 1
		} else {
			freeSpace = append(freeSpace, [2]int{len(line), j})
			line = append(line, slices.Repeat([]int{-1}, j)...)
		}
	}
	// fmt.Println(line)
	// fmt.Println(len(line))
	// fmt.Println(freeSpace)

	right := len(line) - 1
	left := len(line) - 1
	for right > 0 {
		curr := line[right]
		// fmt.Println("Checking ", curr)
		if curr == 0 {
			break
		}
		for left > 0 && line[left] == curr {
			left -= 1
		}
		auxLine := line[left+1 : right+1]
		for i, v := range freeSpace {
			if v[0] > left {
				break
			}
			if v[1] >= len(auxLine) {
				// move
				// fmt.Println("!!!!!!!!!!!!!!!", freeSpace, v)
				for j := range len(auxLine) {
					line[v[0]+j] = curr
					line[left+1+j] = -1
				}
				freeSpace[i] = [2]int{v[0] + len(auxLine), v[1] - len(auxLine)}
				break
			}
		}
		for left > 0 && line[left] == -1 {
			left -= 1
		}
		right = left
		for left > 0 && line[left] == curr-1 {
			left -= 1
		}
		// fmt.Println(line, left, right)
	}
	// fmt.Println(line)

	return checkSum(line)
}
