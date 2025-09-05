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
	result := 0
	for _, line := range lines {
		n, _ := strconv.Atoi(line[:3])
		nPress := PressCode(line)
		fmt.Println("=================", n, nPress)
		result += nPress * n
	}
	return result
}

func PressCode(line string) int {
	start := numKeys['A']

	count := 0
	for _, k := range line {
		moves := GetMoves(start, numKeys[k]) + "A"
		start = numKeys[k]
		fmt.Println(string(k), moves)

		moves2 := PressMoves(moves, moveKeys['A'])
		fmt.Println(string(k), moves2, len(moves2))

		aux := strings.Split(moves2, "A")
		moves3 := ""
		for _, m := range aux[:len(aux)-1] {
			moves3 += PressMoves(m+"A", moveKeys['A'])
		}

		fmt.Println(string(k), moves3, len(moves3))
		count += len(moves3)
	}
	return count
}

var moveKeys = map[rune][2]int{
	//     +---+---+
	//     | ^ | A |
	// +---+---+---+
	// | < | v | > |
	// +---+---+---+
	'^': {0, 1},
	'A': {0, 2},
	'<': {1, 0},
	'v': {1, 1},
	'>': {1, 2},
}

// 0                            2                    9                                  A
// <            A               ^         A          >       ^           ^ A            v    v  v  A
// v   <<  A    >>  ^   A       <    A    >  A       v   A   <    ^   A  A >  A         <vA  A  A  >^A
// <vA <AA >>^A vAA <^A >A      <v<A >>^A vA ^A      <vA >^A <v<A >^A >A A vA ^A <v<A>A>^AAAvA<^A>A

var numKeys = map[rune][2]int{
	// +---+---+---+
	// | 7 | 8 | 9 |
	// +---+---+---+
	// | 4 | 5 | 6 |
	// +---+---+---+
	// | 1 | 2 | 3 |
	// +---+---+---+
	//     | 0 | A |
	//     +---+---+
	'7': {0, 0},
	'8': {0, 1},
	'9': {0, 2},
	'4': {1, 0},
	'5': {1, 1},
	'6': {1, 2},
	'1': {2, 0},
	'2': {2, 1},
	'3': {2, 2},
	'0': {3, 1},
	'A': {3, 2},
}

func PressMoves(moves string, current [2]int) string {
	buttons := make(map[rune]int)
	for _, move := range moves[:len(moves)-1] {
		buttons[move] += 1
	}

	result := ""
	for button, count := range buttons {
		result += GetMoves(current, moveKeys[button]) + strings.Repeat("A", count)
		current = moveKeys[button]
	}

	result += GetMoves(current, moveKeys['A']) + "A"
	return result
}

func GetMoves(start [2]int, target [2]int) string {
	moves := ""
	y, x := target[0]-start[0], target[1]-start[1]

	if x < 0 {
		moves += strings.Repeat("<", -x)
	}
	if y < 0 {
		moves += strings.Repeat("^", -y)
	}
	if y > 0 {
		moves += strings.Repeat("v", y)
	}
	if x > 0 {
		moves += strings.Repeat(">", x)
	}
	return moves
}

func part2(lines []string) any {
	return nil
}
