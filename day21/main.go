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
		numpad := Pad{
			keys:         numKeys,
			forbiddenKey: [2]int{0, 3},
			position:     [2]int{3, 2},
		}
		movepad1 := Pad{
			keys:         moveKeys,
			forbiddenKey: [2]int{0, 0},
			position:     [2]int{0, 2},
		}
		movepad2 := Pad{
			keys:         moveKeys,
			forbiddenKey: [2]int{0, 0},
			position:     [2]int{0, 2},
		}
		moves := numpad.MovesToCode(line)
		moves2 := movepad1.MovesToCode(moves)
		moves3 := movepad2.MovesToCode(moves2)
		nPress := len(moves3)
		fmt.Println("=================", n, nPress)
		result += nPress * n
	}
	return result
}

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

type Pad struct {
	keys         map[rune][2]int
	forbiddenKey [2]int
	position     [2]int
}

func (p *Pad) MovesToCode(code string) string {
	moves := ""
	for _, k := range code {
		moves = moves + p.Move(p.keys[k]) + "A"
	}
	fmt.Println(code, p.position, moves)
	return moves
}

func (p *Pad) Move(target [2]int) string {
	moves := ""
	y, x := target[0]-p.position[0], target[1]-p.position[1]

	lr := ""
	ud := ""
	if x < 0 {
		lr += strings.Repeat("<", -x)
	} else if x > 0 {
		lr += strings.Repeat(">", x)
	}
	if y > 0 {
		ud += strings.Repeat("v", y)
	} else if y < 0 {
		ud += strings.Repeat("^", -y)
	}
	// if there is any left movement, you need to move left before any up/down,
	// except where it passes over the empty square, where it should do up/down movements first,
	// then the left.
	// Otherwise, it should do any up/down movement first then any right movement.
	if x < 0 {
		if target[1] == p.forbiddenKey[1] && p.position[0] == p.forbiddenKey[0] {
			// if going to forbidden column and starting from forbidden row
			moves = moves + ud + lr
		} else {
			moves = moves + lr + ud
		}
	} else {
		if p.position[1] == p.forbiddenKey[1] && target[0] == p.forbiddenKey[0] {
			// if starting from forbidden column and goinf to forbidden row
			moves = moves + lr + ud
		} else {
			moves = moves + ud + lr
		}
	}

	p.position = target
	return moves
}

func part2(lines []string) any {
	return nil
}
