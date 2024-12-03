package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func part1() {
	var filePath string
	if len(os.Args) < 1 {
		filePath = "input-ines.txt"
	} else {
		filePath = os.Args[1]
	}

	fmt.Println("Reading file", filePath)
	dat, err := os.ReadFile(filePath)
	if err != nil {
		fmt.Errorf("Error reading file")
		return
	}

	stringData := string(dat)
	reports := strings.Split(stringData, "\n")

	var result int
	// Iterate through all reports, add 1 to result when report is safe
	for _, report := range reports {
		levels := splitReports(report)
		safe := checkSafe(levels)
		if safe {
			result++
		}
	}

	fmt.Println(result)
}

// Split report lines into list of integers
func part1SplitReports(report string) []int {
	level_str_list := strings.Split(report, " ")

	var levels []int

	for _, level_str := range level_str_list {
		level, _ := strconv.Atoi(level_str)
		levels = append(levels, level)
	}

	return levels
}

func part1CheckSafe(levels []int) bool {
	//  for index, value...
	for i, _ := range levels {
		if i == 0 {
			continue
		}
		diff := levels[i] - levels[i-1]
		// Any two adjacent levels differ by at least one and at most three.
		if diff > 3 || diff < -3 || diff == 0 {
			return false
		}
		// The levels are either all increasing or all decreasing.
		if levels[1]-levels[0] > 0 && diff < 0 || levels[1]-levels[0] < 0 && diff > 00 {
			return false
		}
	}
	return true
}
