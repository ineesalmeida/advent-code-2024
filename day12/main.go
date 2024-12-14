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

func expand(lines []string, i int, j int, regionsMap [][]int, region int) (int, int) {
	target := lines[i][j]
	area := 0
	perimeter := 0

	stack := [][2]int{{i, j}}
	var curr [2]int

	for len(stack) > 0 {
		curr, stack = stack[0], stack[1:]
		if regionsMap[curr[0]][curr[1]] > 0 {
			// already sorted
			continue
		}
		if lines[curr[0]][curr[1]] == target {
			regionsMap[curr[0]][curr[1]] = region
			area += 1

			// top perimeter
			if curr[0] == 0 || lines[curr[0]-1][curr[1]] != target {
				perimeter += 1
			} else {
				stack = append(stack, [2]int{curr[0] - 1, curr[1]})
			}
			// bottom perimeter
			if curr[0] == len(lines)-1 || lines[curr[0]+1][curr[1]] != target {
				perimeter += 1
			} else {
				stack = append(stack, [2]int{curr[0] + 1, curr[1]})
			}
			// left perimeter
			if curr[1] == 0 || lines[curr[0]][curr[1]-1] != target {
				perimeter += 1
			} else {
				stack = append(stack, [2]int{curr[0], curr[1] - 1})
			}
			// right perimeter
			if curr[1] == len(lines[0])-1 || lines[curr[0]][curr[1]+1] != target {
				perimeter += 1
			} else {
				stack = append(stack, [2]int{curr[0], curr[1] + 1})
			}
		}
	}
	return area, perimeter
}

func part1(lines []string) any {
	// initialize
	var regionsMap = make([][]int, len(lines))
	for i := range regionsMap {
		regionsMap[i] = make([]int, len(lines[0]))
	}

	region := 1
	result := 0

	// get regions
	for i, line := range lines {
		for j, _ := range line {
			if regionsMap[i][j] > 0 {
				continue
			}

			area, perimeter := expand(lines, i, j, regionsMap, region)
			result += area * perimeter
			region += 1
		}
	}

	return result
}

func expand2(lines []string, i int, j int, regionsMap [][]int, region int) (int, int) {
	target := lines[i][j]
	area := 0
	perimeter := 0

	stack := [][2]int{{i, j}}
	var curr [2]int

	for len(stack) > 0 {
		curr, stack = stack[0], stack[1:]
		ci, cj := curr[0], curr[1]
		if regionsMap[ci][cj] > 0 {
			// already sorted
			continue
		}
		if lines[ci][cj] == target {
			regionsMap[ci][cj] = region
			area += 1
			borders := [4]int{}

			// top perimeter
			if ci == 0 || lines[ci-1][cj] != target {
				borders[0] = 1
			} else {
				stack = append(stack, [2]int{ci - 1, cj})
			}

			// bottom perimeter
			if ci == len(lines)-1 || lines[ci+1][cj] != target {
				borders[1] = 1
			} else {
				stack = append(stack, [2]int{ci + 1, cj})
			}

			// right perimeter
			if cj == len(lines[0])-1 || lines[ci][cj+1] != target {
				borders[2] = 1

				// if top or bottom also have a border, then add to perimeter
				if borders[0] == 1 {
					perimeter += 1
				}
				if borders[1] == 1 {
					perimeter += 1
				}
			} else {
				stack = append(stack, [2]int{ci, cj + 1})

				// if top doesn't exist and right doesn exist, check if anti corner
				if borders[0] == 0 {
					if lines[ci-1][cj+1] != target {
						perimeter += 1
					}
				}

				// if bottom doesn't exist and right doesn exist, check if anti corner
				if borders[1] == 0 {
					if lines[ci+1][cj+1] != target {
						perimeter += 1
					}
				}
			}

			// left perimeter
			if cj == 0 || lines[ci][cj-1] != target {
				borders[3] = 1
				if borders[0] == 1 {
					perimeter += 1
				}
				if borders[1] == 1 {
					perimeter += 1
				}
			} else {
				stack = append(stack, [2]int{ci, cj - 1})

				// if top doesn't exist and left doesn exist, check if anti corner
				if borders[0] == 0 {
					if lines[ci-1][cj-1] != target {
						perimeter += 1
					}
				}
				// if bottom doesn't exist and left doesn exist, check if anti corner
				if borders[1] == 0 {
					if lines[ci+1][cj-1] != target {
						perimeter += 1
					}
				}
			}
		}
	}
	return area, perimeter
}

func part2(lines []string) any {
	// initialize
	var regionsMap = make([][]int, len(lines))
	for i := range regionsMap {
		regionsMap[i] = make([]int, len(lines[0]))
	}

	region := 1
	result := 0

	// get regions
	for i, line := range lines {
		for j, _ := range line {
			if regionsMap[i][j] > 0 {
				continue
			}

			area, perimeter := expand2(lines, i, j, regionsMap, region)
			result += area * perimeter
			region += 1
		}
	}

	return result
}
