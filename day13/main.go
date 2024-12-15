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
	machines := parseLines(lines)
	// fmt.Println(len(machines), machines)

	price := 0
	for _, machine := range machines {
		priceForMachine := checkMinToWin(machine)
		if priceForMachine > 0 {
			price += priceForMachine
		}
	}
	return price
}

var MAX_PRESS = 100
var PRICE_A = 3
var PRICE_B = 1

func checkMinToWin(machine Machine) int {
	// (Ax, Ay) * Apress + (Bx, By) * Bpress
	// B is cheaper, so start with clicking B until just over the prize, then reduce and add A for the remaining

	a := 0

	var neededB int
	if machine.ButtonB[0] != 0 && machine.ButtonB[1] != 0 {
		neededB = min(machine.Prize[0]/machine.ButtonB[0], machine.Prize[1]/machine.ButtonB[1]) + 1
	} else if machine.ButtonB[0] != 0 {
		neededB = machine.Prize[1]/machine.ButtonB[1] + 1
	} else if machine.ButtonB[1] != 0 {
		neededB = machine.Prize[0]/machine.ButtonB[0] + 1
	}
	b := min(neededB, MAX_PRESS)

	// if we already pressed all Bs available and we still didn't get to the result
	if b == 100 {
		a = min(((machine.Prize[0]-machine.ButtonB[0]*b)/machine.ButtonA[0])+1, ((machine.Prize[0]-machine.ButtonB[0]*b)/machine.ButtonA[0])+1, MAX_PRESS)
	}

	// start position
	p := [2]int{b*machine.ButtonB[0] + a*machine.ButtonA[0], b*machine.ButtonB[1] + a*machine.ButtonA[1]}
	// fmt.Println("Starting at ", a, b, p)

	// max button press did make it
	if a == 100 && b == 100 && (p[0] < machine.Prize[0] || p[1] < machine.Prize[1]) {
		return -1
	}

	for true {
		if p == machine.Prize {
			break
		}

		// try removing one B press at a time, and filling in the rest with A presses
		b -= 1
		if b < 0 {
			return -1
		}

		p[0] -= machine.ButtonB[0]
		p[1] -= machine.ButtonB[1]

		for p[0] < machine.Prize[0] && p[1] < machine.Prize[1] {
			a += 1
			p[0] += machine.ButtonA[0]
			p[1] += machine.ButtonA[1]
		}

		// fmt.Println("Now at ", a, b, p)
	}

	return b*PRICE_B + a*PRICE_A
}

type Machine struct {
	ButtonA [2]int
	ButtonB [2]int
	Prize   [2]int
}

func parseLines(lines []string) []Machine {
	machines := make([]Machine, 0, len(lines)/4)
	machine := Machine{}
	var x int
	var y int
	for _, line := range lines {
		comp := strings.Split(line, " ")
		if comp[0] == "Button" {
			x, _ = strconv.Atoi(comp[2][1 : len(comp[2])-1])
			y, _ = strconv.Atoi(comp[3][1:])
			if comp[1] == "A:" {
				machine.ButtonA = [2]int{x, y}
			} else {
				machine.ButtonB = [2]int{x, y}
			}
		} else if comp[0] == "Prize:" {
			x, _ = strconv.Atoi(comp[1][2 : len(comp[1])-1])
			y, _ = strconv.Atoi(comp[2][2:])
			machine.Prize = [2]int{x, y}
		} else {
			//save
			machines = append(machines, machine)
			machine = Machine{}
		}
	}
	machines = append(machines, machine)
	return machines
}

func part2(lines []string) any {
	machines := parseLines(lines)
	// fmt.Println(len(machines), machines)

	// add the extra zeros to the prize
	for i := range len(machines) {
		machines[i].Prize[0] += 10000000000000
		machines[i].Prize[1] += 10000000000000
	}

	price := 0
	for _, machine := range machines {
		priceForMachine := checkMinToWinWithoutRoof(machine)
		if priceForMachine > 0 {
			price += priceForMachine
		}
	}
	return price
}

func checkMinToWinWithoutRoof(machine Machine) int {
	// (Ax, Ay) * Apress + (Bx, By) * Bpress
	// need to be smarter in this case
	// check if its possible both:
	//      Ax * a + Bx * b == Px
	//      Ay * a + By * b == Py
	// this is a simple math equation:
	//      a = (Px - Bx * b) / Ax
	//      a = (Py - By * b) / Ay
	// which means that
	//      (Px - Bx * b) / Ax = (Py - By * b) / Ay
	//  so
	//      (Px - Bx * b) * Ay = (Py - By * b) * Ax
	//      Ay * Px - Ay * Bx * b = Ax * Py - Ax * By * b
	//      - Ay * Bx * b + Ax * By * b = Ax * Py  - Ay * Px
	//      b * (Ax * By - Ay * Bx) = Ax * Py  - Ay * Px
	//      b  = (Ax * Py  - Ay * Px) / (Ax * By - Ay * Bx)

	b := (machine.ButtonA[0]*machine.Prize[1] - machine.ButtonA[1]*machine.Prize[0]) / (machine.ButtonA[0]*machine.ButtonB[1] - machine.ButtonA[1]*machine.ButtonB[0])
	a := (machine.Prize[0] - machine.ButtonB[0]*b) / machine.ButtonA[0]

	if a*machine.ButtonA[0]+b*machine.ButtonB[0] == machine.Prize[0] && a*machine.ButtonA[1]+b*machine.ButtonB[1] == machine.Prize[1] {
		return b*PRICE_B + a*PRICE_A
	}
	return -1
}
