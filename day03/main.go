package main

import (
	"advent-code/aoc2024/utils"
	"fmt"
	"log"
	"strconv"
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
	result := 0
	input := strings.Join(lines[:], "")
	muls := strings.Split(input, "mul(")
	for _, mul := range muls[1:] {
		result += getMultiplicationResult(mul)
	}
	return result
}

func getMultiplicationResult(line string) int {
	numbers := strings.Split(line, ",")
	if len(numbers) < 2 {
		return 0
	}

	n1Raw := numbers[0]
	n2Raw := numbers[1]

	// Extra process to get n2Raw
	n2Aux := strings.Split(n2Raw, ")")
	if len(n2Aux) < 2 {
		return 0
	}
	n2Raw = n2Aux[0]

	// check if both are integers
	n1, err1 := strconv.Atoi(n1Raw)
	n2, err2 := strconv.Atoi(n2Raw)

	if err1 != nil || err2 != nil {
		return 0
	}

	return n1 * n2
}

func part2(lines []string) any {
	result := 0
	input := strings.Join(lines[:], "")

	// Check for do()'s and don't()'s
	dos := strings.Split(input, "do()")
	var muls []string
	for _, do := range dos {
		do = strings.Split(do, "don't()")[0]
		muls = append(muls, strings.Split(do, "mul(")...)
	}

	for _, mul := range muls[1:] {
		result += getMultiplicationResult(mul)
	}
	return result
}
