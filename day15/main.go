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
	return initial_pos, warehouse, moves
}

func part2(lines []string) any {
	x, y, warehouse, moves := parseInputPart2(lines)
	printWarehouse(warehouse)
	fmt.Println(x, y)

	isWall := func(xx, yy int) bool {
		return warehouse[xx][yy] == rune('#')
	}
	isFree := func(xx, yy int) bool {
		return warehouse[xx][yy] == rune('.')
	}
	isBox := func(xx, yy int) bool {
		return warehouse[xx][yy] == rune('[') || warehouse[xx][yy] == rune(']')
	}
	var moveBoxHor func(xx, yy int, dir int) bool
	moveBoxHor = func(xx, yy int, dir int) bool {
		if isWall(xx, yy+2*dir) {
			return false
		}
		if isFree(xx, yy+2*dir) {
			warehouse[xx][yy+2*dir], warehouse[xx][yy+dir], warehouse[xx][yy] = warehouse[xx][yy+dir], warehouse[xx][yy], rune('.')
			return true
		}
		if isBox(xx, yy+2*dir) {
			ok := moveBoxHor(xx, yy+2*dir, dir)
			if ok {
				warehouse[xx][yy+2*dir], warehouse[xx][yy+dir], warehouse[xx][yy] = warehouse[xx][yy+dir], warehouse[xx][yy], rune('.')
				return true
			}
			return false
		}
		return false
	}
	var canMoveBoxVer func(xx, yy int, dir int) bool
	canMoveBoxVer = func(xx, yy int, dir int) bool {
		var yy_l, yy_r int
		if warehouse[xx][yy] == rune('[') {
			yy_l, yy_r = yy, yy+1
		} else {
			yy_l, yy_r = yy-1, yy
		}
		if isWall(xx+dir, yy_l) || isWall(xx+dir, yy_r) {
			return false
		}
		if isFree(xx+dir, yy_l) && isFree(xx+dir, yy_r) {
			return true
		}
		l_box, r_box := isBox(xx+dir, yy_l), isBox(xx+dir, yy_r)
		if l_box {
			ok := canMoveBoxVer(xx+dir, yy_l, dir)
			if !ok {
				return false
			}
		}
		if r_box {
			ok := canMoveBoxVer(xx+dir, yy_r, dir)
			if !ok {
				return false
			}
		}
		return true
	}
	var moveBoxVer func(xx, yy int, dir int, original bool) bool
	moveBoxVer = func(xx, yy int, dir int, original bool) bool {
		var yy_l, yy_r int
		if warehouse[xx][yy] == rune('[') {
			yy_l, yy_r = yy, yy+1
		} else {
			yy_l, yy_r = yy-1, yy
		}
		if isWall(xx+dir, yy_l) || isWall(xx+dir, yy_r) {
			return false
		}
		if isFree(xx+dir, yy_l) && isFree(xx+dir, yy_r) {
			warehouse[xx+dir][yy_l], warehouse[xx+dir][yy_r], warehouse[xx][yy_l], warehouse[xx][yy_r] = warehouse[xx][yy_l], warehouse[xx][yy_r], rune('.'), rune('.')
			return true
		}
		l_box, r_box := isBox(xx+dir, yy_l), isBox(xx+dir, yy_r)
		// if it's the same box, set right oneto null
		if l_box && warehouse[xx+dir][yy_r] == rune(']') {
			r_box = false
		}
		ok_l, ok_r := true, true
		if l_box {
			ok_l = canMoveBoxVer(xx+dir, yy_l, dir)
		}
		if r_box {
			ok_r = canMoveBoxVer(xx+dir, yy_r, dir)
		}

		if ok_l && ok_r && original {
			if l_box {
				moveBoxVer(xx+dir, yy_l, dir, true)
			}
			if r_box {
				moveBoxVer(xx+dir, yy_r, dir, true)
			}
			//[][]..
			//.[].
			//....
			//[]@.
			warehouse[xx+dir][yy_l], warehouse[xx+dir][yy_r], warehouse[xx][yy_l], warehouse[xx][yy_r] = warehouse[xx][yy_l], warehouse[xx][yy_r], rune('.'), rune('.')
			return true
		}
		return false
	}

	moveRobot := func(dir rune) {
		switch dir {
		case rune('<'):
			if isWall(x, y-1) {
				return
			}
			if isFree(x, y-1) {
				y--
				return
			}
			if !isBox(x, y-1) {
				fmt.Println("ERROR: unexpected rune")
			}
			ok := moveBoxHor(x, y-1, -1)
			if ok {
				y--
			}
		case rune('>'):
			if isWall(x, y+1) {
				return
			}
			if isFree(x, y+1) {
				y++
				return
			}
			if !isBox(x, y+1) {
				fmt.Println("ERROR: unexpected rune")
			}
			ok := moveBoxHor(x, y+1, 1)
			if ok {
				y++
			}
		case rune('^'):
			if isWall(x-1, y) {
				return
			}
			if isFree(x-1, y) {
				x--
				return
			}
			if !isBox(x-1, y) {
				fmt.Println("ERROR: unexpected rune")
			}
			ok := moveBoxVer(x-1, y, -1, true)
			if ok {
				x--
			}
		case rune('v'):
			if isWall(x+1, y) {
				return
			}
			if isFree(x+1, y) {
				x++
				return
			}
			if !isBox(x+1, y) {
				fmt.Println("ERROR: unexpected rune")
			}
			ok := moveBoxVer(x+1, y, 1, true)
			if ok {
				x++
			}
		}
	}

	for _, move := range moves {
		moveRobot(move)
		// printWarehouse(warehouse)
	}

	printWarehouse(warehouse)
	// calculate result
	result := 0
	for i, line := range warehouse {
		for j, val := range line {
			if val == '[' {
				result += 100*i + j
			}
		}
	}

	return result
}

func parseInputPart2(lines []string) (int, int, [][]rune, []rune) {
	warehouse := make([][]rune, 0)
	moves := make([]rune, 0)
	var x, y int
	for i, line := range lines {
		if line == "" {
			continue
		} else if line[0] != '#' {
			for _, m := range line {
				moves = append(moves, m)
			}
			continue
		}

		warehouse = append(warehouse, make([]rune, len(lines[0])*2))
		for j, val := range line {
			if val == '@' {
				x = i
				y = 2 * j
				warehouse[i][2*j] = rune('.')   // 0 2 4...
				warehouse[i][2*j+1] = rune('.') // 1 3 5...
			} else if val == 'O' {
				warehouse[i][2*j] = rune('[')
				warehouse[i][2*j+1] = rune(']')
			} else {
				warehouse[i][2*j] = rune(val)
				warehouse[i][2*j+1] = rune(val)
			}
		}
	}

	return x, y, warehouse, moves
}
