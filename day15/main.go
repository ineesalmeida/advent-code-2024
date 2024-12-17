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
	robot_pos, warehouse, moves := parseInput(lines)
	// printWarehouse(warehouse)
	// fmt.Println(robot_pos, moves)

	for _, move := range moves {
		// fmt.Println("===============================", i)
		ci, cj := robot_pos[0]+move[0], robot_pos[1]+move[1]
		if warehouse[ci][cj] == '#' {
			// hit a wall
			continue
		} else if warehouse[ci][cj] == '.' {
			// hit nothing, move robot
			robot_pos = [2]int{ci, cj}
		} else if warehouse[ci][cj] == 'O' {
			// hit box, try to move it else say in same place
			aux := 0
			for true {
				bi, bj := ci+move[0]*aux, cj+move[1]*aux
				if warehouse[bi][bj] == '.' {
					// Found a free spot
					warehouse[ci][cj], warehouse[bi][bj] = warehouse[bi][bj], warehouse[ci][cj]
					robot_pos = [2]int{ci, cj}
					break
				} else if warehouse[bi][bj] == '#' {
					// Found a wall before a free spot
					break
				}
				aux += 1
			}
		} else {
			fmt.Errorf("Unexpected rune in map...")
			break
		}
		// fmt.Println(move, robot_pos)
		// printWarehouse(warehouse)
	}

	// calculate result
	result := 0
	for i, line := range warehouse {
		for j, val := range line {
			if val == 'O' {
				result += 100*i + j
			}
		}
	}

	return result
}

func printWarehouse(warehouse [][]rune) {
	for i := range len(warehouse) {
		for j := range len(warehouse[0]) {
			fmt.Printf(string(warehouse[i][j]))
		}
		fmt.Printf("\n")
	}
}

var DIRECTION_MAP = map[rune][2]int{
	'^': {-1, 0},
	'>': {0, +1},
	'v': {+1, 0},
	'<': {0, -1},
}

func parseInput(lines []string) ([2]int, [][]rune, [][2]int) {
	warehouse := make([][]rune, 0)
	moves := make([][2]int, 0)
	initial_pos := [2]int{}
	for i, line := range lines {
		if line == "" {
			continue
		} else if line[0] != '#' {
			for _, m := range line {
				moves = append(moves, DIRECTION_MAP[m])
			}
			continue
		}

		warehouse = append(warehouse, make([]rune, len(lines[0])))
		for j, val := range line {
			if val == '@' {
				initial_pos[0] = i
				initial_pos[1] = j
				warehouse[i][j] = rune('.')
			} else {
				warehouse[i][j] = rune(val)
			}
		}
	}
	// initial robot position
	// if

	return initial_pos, warehouse, moves
}

func part2(lines []string) any {
	return nil
}
