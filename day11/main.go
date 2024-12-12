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
	stones := strings.Split(lines[0], " ")
	var tmpStones []string
	for range 25 {
		// fmt.Println("=========== ", stones)
		tmpStones = nil
		for _, stone := range stones {
			if stone == "0" {
				tmpStones = append(tmpStones, "1")
			} else if len(stone)%2 == 0 {
				// no need to worry about trailing zeros sfor the first one
				tmpStones = append(tmpStones, stone[:len(stone)/2])
				rightStone, _ := strconv.Atoi(stone[len(stone)/2:])
				tmpStones = append(tmpStones, strconv.Itoa(rightStone))
			} else {
				n, _ := strconv.Atoi(stone)
				n *= 2024
				tmpStones = append(tmpStones, strconv.Itoa(n))
			}
		}
		stones = tmpStones
	}
	return len(stones)
}

var BLINKS = 75
var currentStones = make(map[string]int)

func blink() {
	// fmt.Println("=========== ", currentStones)

	// create copy
	tmpStones := make(map[string]int)
	for k, v := range currentStones {
		tmpStones[k] = v
	}

	for stone, count := range tmpStones {
		if count == 0 {
			continue
		}
		currentStones[stone] -= count
		if stone == "0" {
			currentStones["1"] += count
		} else if len(stone)%2 == 0 {
			// no need to worry about trailing zeros sfor the first one
			leftStone := stone[:len(stone)/2]
			n, _ := strconv.Atoi(stone[len(stone)/2:])
			rightStone := strconv.Itoa(n)

			currentStones[leftStone] += count
			currentStones[rightStone] += count
		} else {
			n, _ := strconv.Atoi(stone)
			n *= 2024
			currentStones[strconv.Itoa(n)] += count
		}
	}

	// sum := 0
	// for stone, count := range currentStones {
	// 	if count > 0 {
	// 		fmt.Println(stone, count)
	// 		sum += count
	// 	}
	// }
	// fmt.Println("TOTAL AFTER 1 BLINK", sum)
}

func part2(lines []string) any {

	// initialize
	for _, stone := range strings.Split(lines[0], " ") {
		currentStones[stone] += 1
	}

	// blink
	for range BLINKS {
		blink()
	}

	// compute result
	result := 0
	for _, count := range currentStones {
		result += count
	}
	return result
}
