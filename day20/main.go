package main

import (
	"advent-code/aoc2024/utils"
	"fmt"
	"log"
	"time"
)

func main() {
	lines, err := utils.FileToLines("test-input.txt")
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
	goThroughRaceTrack(lines)

	result := 0

	for i, line := range lines {
		// skip the first and last line
		if i == 0 || i == len(lines)-1 {
			continue
		}
		for j, c := range line {
			// skip the first and last column
			if j == 0 || j == len(line)-1 {
				continue
			}
			// skip any non walls
			if c != '#' {
				continue
			}

			// found  wall, check if there is any in between that would skip a few pico seconds
			if track[i+1][j] != 0 && track[i-1][j] != 0 {
				// check shortcut
				// debug[i][j] = '#'
				savings := track[i+1][j] - track[i-1][j]
				if savings < 0 {
					savings = -savings
				}
				savings -= 2

				if savings >= SAVING_THREADHOLD {
					// fmt.Println("Saving", i, j, savings)
					result += 1
				}
			}
			if track[i][j+1] != 0 && track[i][j-1] != 0 {
				// debug[i][j] = '#'
				// check shortcut
				savings := track[i][j+1] - track[i][j-1]
				if savings < 0 {
					savings = -savings
				}
				savings -= 2

				if savings >= SAVING_THREADHOLD {
					fmt.Println("Saving", i, j, savings)
					result += 1
				}
			}
		}
	}

	return result
}

var start [2]int
var end [2]int
var track [][]int

func findStart(lines []string) {
	for i, line := range lines {
		for j, c := range line {
			if c == 'S' {
				start = [2]int{i, j}
				if end[0] != 0 {
					return
				}
			} else if c == 'E' {
				end = [2]int{i, j}
				if start[0] != 0 {
					return
				}
			}
		}
	}
}

func goThroughRaceTrack(lines []string) {
	findStart(lines)

	// start track
	for range lines {
		track = append(track, make([]int, len(lines[0])))
	}

	i, j := start[0], start[1]
	var dir [2]int
	if lines[i+1][j] == '.' {
		dir[0] = 1
	} else if lines[i-1][j] == '.' {
		dir[0] = -1
	} else if lines[i][j+1] == '.' {
		dir[1] = 1
	} else if lines[i][j-1] == '.' {
		dir[1] = -1
	}

	counter := 1
	for true {
		track[i][j] = counter
		counter += 1

		if lines[i+dir[0]][j+dir[1]] == '#' {
			// check to the side
			if dir[1] == 0 {
				// check +/- j
				if lines[i][j-1] == '.' {
					dir[0] = 0
					dir[1] = -1
				} else {
					dir[0] = 0
					dir[1] = 1
				}
			} else {
				// check +/- i
				if lines[i-1][j] == '.' {
					dir[0] = -1
					dir[1] = 0
				} else {
					dir[0] = 1
					dir[1] = 0
				}
			}
		}

		i += dir[0]
		j += dir[1]

		if i == end[0] && j == end[1] {
			track[i][j] = counter
			break
		}
	}
	utils.PrintMatrix(track)
}

var SAVING_THREADHOLD = 76
var RADIUS = 20

func part2(lines []string) any {

	result := 0

	for i, row := range track {
		// skip the first and last line
		if i == 0 || i == len(track)-1 {
			continue
		}
		for j, c := range row {
			// skip the first and last column
			if j == 0 || j == len(row)-1 {
				continue
			}
			// skip any non walls
			if c == 0 {
				continue
			}

			// check all around
			ii := 0
			jj := 0
			// 1, 19
			// 2, 18
			for ii < RADIUS {
				if i-ii < 0 || i+ii >= len(track) {
					ii += 1
					continue
				}
				for jj < RADIUS-ii {
					if j-jj < 0 || j+jj >= len(track[0]) {
						jj += 1
						continue
					}

					if track[i+ii][j+jj]-track[i][j]-ii-jj >= SAVING_THREADHOLD {
						fmt.Println("here 1", i, j, ii, jj)
						result += 1
					}
					if track[i-ii][j-jj]-track[i][j]-ii-jj >= SAVING_THREADHOLD {
						fmt.Println("here 2", i, j, ii, jj)
						result += 1
					}
					if track[i+ii][j-jj]-track[i][j]-ii-jj >= SAVING_THREADHOLD {
						fmt.Println("here 3", i, j, ii, jj)
						result += 1
					}
					if track[i-ii][j+jj]-track[i][j]-ii-jj >= SAVING_THREADHOLD {
						fmt.Println("here 4", i, j, ii, jj)
						result += 1
					}

					jj += 1
				}
				ii += 1
			}
		}
	}

	return result
}

func Abs(integer int) int {
	if integer < 0 {
		return -integer
	}
	return integer
}
