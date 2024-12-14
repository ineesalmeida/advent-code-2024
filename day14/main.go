package main

import (
	"advent-code/aoc2024/utils"
	"bufio"
	"fmt"
	"log"
	"os"
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
	q := [4]int{}
	t := 100

	for _, line := range lines {
		val := strings.Split(line, " ")
		p := strings.Split(val[0][2:], ",")
		v := strings.Split(val[1][2:], ",")

		px, _ := strconv.Atoi(p[0])
		py, _ := strconv.Atoi(p[1])
		vx, _ := strconv.Atoi(v[0])
		vy, _ := strconv.Atoi(v[1])

		x := (px + t*vx) % nx
		if x < 0 {
			x += nx
		}
		y := (py + t*vy) % ny
		if y < 0 {
			y += ny
		}

		if x < (nx-1)/2 {
			if y < (ny-1)/2 {
				// top left
				q[0] += 1
			} else if y > (ny-1)/2 {
				// top right
				q[1] += 1
			}
		} else if x > (nx-1)/2 {
			if y < (ny-1)/2 {
				// bottom left
				q[2] += 1
			} else if y > (ny-1)/2 {
				// bottom right
				q[3] += 1
			}
		}

		fmt.Println(p, v, x, y, q)
	}
	result := 1
	for _, qq := range q {
		result *= qq
	}
	return result
}

// nx := 11
// ny := 7
var nx = 101
var ny = 103

func part2(lines []string) any {
	pv := make([][4]int, len(lines))
	counter := 0

	for i, line := range lines {
		val := strings.Split(line, " ")
		p := strings.Split(val[0][2:], ",")
		v := strings.Split(val[1][2:], ",")

		pv[i][0], _ = strconv.Atoi(p[0])
		pv[i][1], _ = strconv.Atoi(p[1])
		pv[i][2], _ = strconv.Atoi(v[0])
		pv[i][3], _ = strconv.Atoi(v[1])
	}

	input := bufio.NewScanner(os.Stdin)
	show := false
	aux := make(map[[2]int]bool)
	for true {
		counter += 1
		oneSecond(pv)

		// is there at least a vertical line?
		show = false
		aux = make(map[[2]int]bool)
		for _, robot := range pv {
			aux[[2]int{robot[0], robot[1]}] = true
			aux2 := 0
			for i := range ny {
				_, ok1 := aux[[2]int{robot[0], robot[1] + i}]
				_, ok2 := aux[[2]int{robot[0], robot[1] - i}]
				if !ok1 && !ok2 {
					break
				}
				if ok1 {
					aux2 += 1
				}
				if ok2 {
					aux2 += 1
				}
				// a verical line of at least 10 in a row
				if aux2 > 10 {
					show = true
					break
				}
			}
			if show {
				break
			}
		}

		if !show {
			continue
		}

		// else print and ask
		printResult(pv)

		fmt.Println("Is it a tree?")
		input.Scan()
		if len(input.Text()) > 0 {
			break
		}
	}
	return counter
}

func printResult(pv [][4]int) {

	// create result matrix
	result := make([]string, ny)
	for i := range result {
		result[i] = strings.Repeat(".", nx)
	}

	// fill it out
	for _, robot := range pv {
		result[robot[1]] = result[robot[1]][:robot[0]] + "#" + result[robot[1]][robot[0]+1:]
	}

	// print it
	for _, line := range result {
		fmt.Println(line)
	}
}

func oneSecond(pv [][4]int) {

	for i, robot := range pv {
		// fmt.Println(robot)
		pv[i][0] += robot[2]
		pv[i][1] += robot[3]
		pv[i][0] = pv[i][0] % nx
		pv[i][1] = pv[i][1] % ny

		if pv[i][0] < 0 {
			pv[i][0] += nx
		}
		if pv[i][1] < 0 {
			pv[i][1] += ny
		}
	}
}
