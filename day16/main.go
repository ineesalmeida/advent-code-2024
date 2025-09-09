package main

import (
	"advent-code/aoc2024/utils"
	"fmt"
	"log"
	"maps"
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

type Direction int

const (
	N Direction = iota + 1
	E
	S
	W
)

type step struct {
	x        int
	y        int
	dir      Direction
	value    int
	prevStep *step
}

func (s *step) isTarget() bool {
	// fmt.Println("TARGET:", target, "X:", s.x, "Y:", s.y)
	return s.x == target[1] && s.y == target[0]
}

func (s *step) code() string {
	return fmt.Sprintf("%v,%v,%d", s.x, s.y, s.dir)
}

type Stack struct {
	items []step
}

var target = [2]int{}

func (s *Stack) Push(data step) {
	s.items = append(s.items, data)
}

func (s *Stack) Pop() (step, bool) {
	if s.IsEmpty() {
		return step{}, false
	}
	top := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return top, true
}

func (s *Stack) IsEmpty() bool {
	if len(s.items) == 0 {
		return true
	}
	return false
}

var seen = make(map[string]int)

func part1(nodes []string) any {

	stack := Stack{}
	parseInput(nodes, &stack)
	// fmt.Println(target)

	result := -1
	for true {
		nextStep, ok := stack.Pop()
		if !ok {
			break
		}
		if result != -1 && nextStep.value >= result {
			// fmt.Println("DEBUG: too large!", nextStep.value)
			continue
		}
		if nextStep.isTarget() {
			// fmt.Println("DEBUG: found result!", nextStep.value)
			result = nextStep.value
			continue
		} else if v, ok := seen[nextStep.code()]; ok && v < nextStep.value {
			// fmt.Println("DEBUG: have seen", nextStep.code())
			continue
		}
		seen[nextStep.code()] = nextStep.value

		nextSteps(nextStep, &stack, nodes)
		// fmt.Println(len(seen), len(stack.items))
	}
	return result
}

func abs(v int) int {
	return max(v, -v)
}

// Check for possible enxt steps
func nextSteps(s step, stack *Stack, nodes []string) {

	if nodes[s.y+1][s.x] != '#' {
		dv := abs(int(s.dir - S)) // 0, 1, 2, 3
		if dv == 3 {
			dv = 1 // if we need to rotate 3 times, we can torate once in the opposite dir
		}
		stack.Push(step{
			x:        s.x,
			y:        s.y + 1,
			dir:      S,
			value:    s.value + dv*1000 + 1,
			prevStep: &s,
		})
	}
	if nodes[s.y-1][s.x] != '#' {
		dv := abs(int(s.dir - N))
		if dv == 3 {
			dv = 1 // if we need to rotate 3 times, we can torate once in the opposite dir
		}
		stack.Push(step{
			x:        s.x,
			y:        s.y - 1,
			dir:      N,
			value:    s.value + dv*1000 + 1,
			prevStep: &s,
		})
	}
	if nodes[s.y][s.x+1] != '#' {
		dv := abs(int(s.dir - E))
		if dv == 3 {
			dv = 1 // if we need to rotate 3 times, we can torate once in the opposite dir
		}
		stack.Push(step{
			x:        s.x + 1,
			y:        s.y,
			dir:      E,
			value:    s.value + dv*1000 + 1,
			prevStep: &s,
		})
	}
	if nodes[s.y][s.x-1] != '#' {
		dv := abs(int(s.dir - W))
		if dv == 3 {
			dv = 1 // if we need to rotate 3 times, we can torate once in the opposite dir
		}
		stack.Push(step{
			x:        s.x - 1,
			y:        s.y,
			dir:      W,
			value:    s.value + dv*1000 + 1,
			prevStep: &s,
		})
	}
}

func parseInput(lines []string, stack *Stack) {
	for i, line := range lines {
		for j, node := range line {
			if node == 'S' {
				stack.Push(step{
					x:     j,
					y:     i,
					dir:   E,
					value: 0,
				})
			} else if node == 'E' {
				target[0] = i
				target[1] = j
			}
		}
	}
}

func part2(nodes []string) any {

	stack := Stack{}
	parseInput(nodes, &stack)
	fmt.Println(target)

	bestSeats := make([]step, 0)
	result := -1
	for true {
		nextStep, ok := stack.Pop()
		if !ok {
			break
		}
		if result != -1 && nextStep.value > result {
			// fmt.Println("DEBUG: too large!", nextStep.value)
			continue
		}
		if nextStep.isTarget() {
			if result == -1 || nextStep.value < result {
				fmt.Println("DEBUG: found better result!", result, nextStep.value)
				result = nextStep.value
				bestSeats = nil
			} else {
				// fmt.Println("DEBUG: found equal result!", result, nextStep.value)
			}
			bestSeats = append(bestSeats, nextStep)
			// fmt.Println("DEBUG: bestSeats:", len(bestSeats))

			continue
		} else if result == nextStep.value {
			continue
		} else if v, ok := seen[nextStep.code()]; ok && v < nextStep.value {
			// fmt.Println("DEBUG: have seen", nextStep.code())
			continue
		}
		seen[nextStep.code()] = nextStep.value

		nextSteps(nextStep, &stack, nodes)
		// fmt.Println(len(seen), len(stack.items))
	}

	seatsMap := make(map[string]int)
	for _, s := range bestSeats {
		maps.Copy(seatsMap, backtrack(s))
	}
	return len(seatsMap)
}

func backtrack(s step) map[string]int {
	seats := make(map[string]int)
	for true {
		seats[fmt.Sprintf("%v,%v", s.x, s.y)]++
		if s.prevStep == nil {
			break
		}
		s = *s.prevStep
	}
	// fmt.Println("DEBUG: seats:", seats)
	return seats
}
