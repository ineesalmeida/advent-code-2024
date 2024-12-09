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

func getEquationParts(line string) (int, []int) {
	var res int
	var nums []int
	lineParts := strings.Split(line, ": ")
	res, _ = strconv.Atoi(lineParts[0])
	for _, val := range strings.Split(lineParts[1], " ") {
		i, _ := strconv.Atoi(val)
		nums = append(nums, i)
	}
	return res, nums
}

func part1(lines []string) any {
	result := 0
	// cache := make(map[string]int)
	for _, line := range lines {
		res, nums := getEquationParts(line)
		result += checkEq(res, nums)
	}
	return result
}

func checkEq(res int, nums []int) int {
	// fmt.Println(res, nums)

	n := len(nums)
	if n == 2 {
		if nums[0]+nums[1] == res || nums[0]*nums[1] == res {
			return res
		} else {
			return 0
		}
	}

	var resultMult int
	var resultSum int

	// Check sums
	resultSum = checkEq(res-nums[n-1], nums[:n-1])

	// Check mult
	if resultSum == 0 && res%nums[n-1] == 0 {
		newRes := res / nums[n-1]
		resultMult = checkEq(newRes, nums[:n-1])
	}

	// If either works
	if resultMult+resultSum > 0 {
		return res
	}
	return 0
}

func part2(lines []string) any {
	result := 0
	// cache := make(map[string]int)
	for _, line := range lines {
		res, nums := getEquationParts(line)
		result += checkEqWithConcat(res, nums)
	}
	return result
}

func concat(n1 int, n2 int) int {
	i, _ := strconv.Atoi(strconv.Itoa(n1) + strconv.Itoa(n2))
	return i
}

func checkEqWithConcat(res int, nums []int) int {
	// fmt.Println(res, nums)

	if len(nums) == 2 {
		if nums[0]+nums[1] == res || nums[0]*nums[1] == res || concat(nums[0], nums[1]) == res {
			// fmt.Println("!!!!!", nums[0]+nums[1], nums[0]*nums[1], concat(nums[0], nums[1]))
			return res
		} else {
			return 0
		}
	}

	// Check sums
	result := checkEqWithConcat(res, append([]int{nums[0] + nums[1]}, nums[2:]...))
	if result == res {
		return res
	}

	// Check mult
	if nums[0] != 0 && nums[1] != 0 && nums[0]*nums[1] <= res {
		result = checkEqWithConcat(res, append([]int{nums[0] * nums[1]}, nums[2:]...))
		if result == res {
			return res
		}
	}

	// Check concat
	concatenated := concat(nums[0], nums[1])
	if concatenated <= res {
		result = checkEqWithConcat(res, append([]int{concatenated}, nums[2:]...))
		if result == res {
			return res
		}
	}
	return 0
}
