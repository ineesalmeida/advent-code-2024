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
	initial_pos := parseLines(lines)
	return walk(initial_pos, lines, 0)
}

// directions
// EAST 1
// DOWN 2
// WEST 3
// TOP  4

func parseLines(lines []string) [3]int {
	initial_pos := [3]int{}
	for i, line := range lines {
		for j, val := range line {
			if val == 'S' {
				initial_pos = [3]int{i, j, 1}
				return initial_pos
			}
		}
	}
	return initial_pos
}

func walk(pos [3]int, lines []string, n_moves int) int {
	// fmt.Println("Checking", pos)
	// fmt.Println("Checking", curr_path)
	if n_moves > 500 {
		// seen it before! it's a loop
		return -1
	}
	n_moves += 1

	// input := bufio.NewScanner(os.Stdin)
	// _ = input.Scan()

	fi, fj, fd := pos[0], pos[1], pos[2]
	li, lj, ld := pos[0], pos[1], pos[2]
	ri, rj, rd := pos[0], pos[1], pos[2]

	if pos[2] == 1 { // EAST
		fj += 1
		li -= 1
		ri += 1
	} else if pos[2] == 2 { //SOUTH
		fi += 1
		lj += 1
		rj -= 1
	} else if pos[2] == 3 { // WEST
		fj -= 1
		li += 1
		ri -= 1
	} else { //NORTH
		fi -= 1
		lj -= 1
		rj += 1
	}

	ld -= 1
	if ld == 0 {
		ld = 4
	}
	rd += 1
	if rd == 5 {
		rd = 1
	}

	if lines[fi][fj] == 'E' {
		// reached target
		return 1
	} else if lines[li][lj] == 'E' || lines[ri][rj] == 'E' {
		return 1001
	}

	lowest_path := -1

	if lines[fi][fj] != '#' {
		// if not a wall, go front
		fw := walk([3]int{fi, fj, fd}, lines, n_moves)
		if fw != -1 {
			lowest_path = fw + 1
			lowest_path = fw + 1
		}
	}
	if lines[ri][rj] != '#' {
		// if not a wall, try going left
		rw := walk([3]int{ri, rj, rd}, lines, n_moves)
		if rw != -1 && (lowest_path == -1 || rw+1001 <= lowest_path) {
			lowest_path = rw + 1001
		}
	}
	if lines[li][lj] != '#' {
		// if not a wall, try going left
		lw := walk([3]int{li, lj, ld}, lines, n_moves)
		if lw != -1 && (lowest_path == -1 || lw+1001 <= lowest_path) {
			lowest_path = lw + 1001
		}
	}
	// fmt.Println(lowest_path)
	return lowest_path
}

// func copyMap(og map[[2]int]bool) map[[2]int]bool {
// 	new := make(map[[2]int]bool)
// 	for k, v := range og {
// 		new[k] = v
// 	}
// 	return new
// }

func part2(lines []string) any {
	return nil
}
