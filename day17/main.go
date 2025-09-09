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

func parseLines(lines []string) (int, int, int, []int) {
	a := strings.Split(lines[0], ": ")[1]
	b := strings.Split(lines[1], ": ")[1]
	c := strings.Split(lines[2], ": ")[1]
	p := strings.Split(strings.Split(lines[4], ": ")[1], ",")

	A, _ := strconv.Atoi(a)
	B, _ := strconv.Atoi(b)
	C, _ := strconv.Atoi(c)

	P := make([]int, len(p))
	for i := range p {
		P[i], _ = strconv.Atoi(p[i])
	}

	return A, B, C, P
}

type Computer struct {
	A, B, C         int
	program, output []int
}

func (c Computer) combo(operand int) int {
	switch operand {
	case 4:
		return c.A
	case 5:
		return c.B
	case 6:
		return c.C
	case 7:
		panic("invalid combo operand")
	default:
		return operand
	}
}

func (c *Computer) execute(code, operand int) {
	// Case 3 is only used to loop back to start
	switch code {
	case 0:
		c.A = c.A >> c.combo(operand)
	case 1:
		c.B = c.B ^ operand
	case 2:
		c.B = c.combo(operand) % 8
	case 4:
		c.B = c.B ^ c.C
	case 5:
		c.output = append(c.output, c.combo(operand)%8)
	case 6:
		c.B = c.A >> c.combo(operand)
	case 7:
		c.C = c.A >> c.combo(operand)
	}
}

func (c Computer) Output() string {
	outStr := make([]string, len(c.output))
	for i, val := range c.output {
		outStr[i] = strconv.Itoa(val)
	}
	return strings.Join(outStr, ",")
}

func part1(lines []string) any {
	a, b, c, p := parseLines(lines)
	computer := Computer{
		A:       a,
		B:       b,
		C:       c,
		program: p,
	}
	fmt.Println("Computer:", computer)

	for i := 0; i < len(computer.program); i += 2 {
		code, operand := computer.program[i], computer.program[i+1]
		fmt.Println("DEBUG:", code, operand)
		fmt.Println("DEBUG:", computer)
		computer.execute(code, operand)
	}

	fmt.Println(computer.Output())
	return nil
}

func part2(lines []string) any {
	return nil
}
