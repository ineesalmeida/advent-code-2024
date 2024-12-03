package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
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
	var goodReports []string
	// Iterate through all reports, add 1 to result when report is safe
	for _, report := range reports {
		levels := splitReports(report)
		safe := checkSafe(levels)
		if safe {
			result++
			goodReports = append(goodReports, report)
		}
	}

	printLines("tmp-305", goodReports)
	fmt.Println(result)
}

// Split report lines into list of integers
func splitReports(report string) []int {
	level_str_list := strings.Split(report, " ")

	var levels []int

	for _, level_str := range level_str_list {
		level, _ := strconv.Atoi(level_str)
		levels = append(levels, level)
	}

	return levels
}

func printLines(filePath string, values []string) error {
	f, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer f.Close()
	for _, value := range values {
		fmt.Fprintln(f, value) // print values to f, one per line
	}
	return nil
}

func checkSafe(levels []int) bool {

	diffs := make([]int, len(levels)-1)
	for i := range levels {
		if i == 0 {
			continue
		}
		diffs[i-1] = levels[i] - levels[i-1]
	}

	return checkDiffs(diffs)
}

func checkDiffs(diffs []int) bool {

	// multiplier := 1
	// for i, diff := range diffs {

	// }

	var increasingDiffs []int
	var decreasingDiffs []int
	var wrongDiffs []int

	for i, diff := range diffs {

		if diff < 0 && diff >= -3 {
			decreasingDiffs = append(decreasingDiffs, i)
		} else if diff > 0 && diff <= 3 {
			increasingDiffs = append(increasingDiffs, i)
		} else {
			wrongDiffs = append(wrongDiffs, i)
		}
	}

	if len(increasingDiffs) == len(diffs) || len(decreasingDiffs) == len(diffs) {
		return true
	}

	if len(increasingDiffs) == len(diffs)-1 {
		// check the decreaseing or wrong
		wrongDiffs = append(wrongDiffs, decreasingDiffs...)
		wrongIndex := wrongDiffs[0]

		if wrongIndex == 0 || wrongIndex == len(diffs)-1 {
			return true
		} else {
			diff := diffs[wrongIndex] + diffs[wrongIndex-1]
			if diff > 0 && diff <= 3 {
				return true
			}
			diff = diffs[wrongIndex] + diffs[wrongIndex+1]
			if diff > 0 && diff <= 3 {
				return true
			}
			return false
		}
	} else if len(decreasingDiffs) == len(diffs)-1 {
		// check the incresing or wrong
		wrongDiffs = append(wrongDiffs, increasingDiffs...)
		wrongIndex := wrongDiffs[0]

		if wrongIndex == 0 || wrongIndex == len(diffs)-1 {
			return true
		} else {
			diff := diffs[wrongIndex] + diffs[wrongIndex-1]
			if diff < 0 && diff >= -3 {
				return true
			}
			diff = diffs[wrongIndex] + diffs[wrongIndex+1]
			if diff < 0 && diff >= -3 {
				return true
			}
			return false
		}
	} else if len(increasingDiffs) == len(diffs)-2 {
		// check the decreaseing or wrong
		wrongDiffs = append(wrongDiffs, decreasingDiffs...)
		wrongIndex1 := wrongDiffs[0]
		wrongIndex2 := wrongDiffs[1]

		if !(wrongIndex1-wrongIndex2 == 1 || wrongIndex1-wrongIndex2 == -1) {
			return false
		}

		diff := diffs[wrongIndex1] + diffs[wrongIndex2]
		if diff > 0 && diff <= 3 {
			return true
		}
		return false
	} else if len(decreasingDiffs) == len(diffs)-2 {
		// check the decreaseing or wrong
		wrongDiffs = append(wrongDiffs, increasingDiffs...)
		wrongIndex1 := wrongDiffs[0]
		wrongIndex2 := wrongDiffs[1]

		if !(wrongIndex1-wrongIndex2 == 1 || wrongIndex1-wrongIndex2 == -1) {
			return false
		}

		diff := diffs[wrongIndex1] + diffs[wrongIndex2]
		if diff < 0 && diff >= -3 {
			return true
		}
		return false
	}
	return false
}

// 1 2 -1 3
// wrong diff is index 2
// wrong index = 2
// diffs[wrongIndex] = -1
// diffs[wrongIndex-1] = 2
// combinedDiff = -1 + 2 = 1

// func checkSign(diffs []int) bool, int {
// 	increase := 0
// 	for _, diff := range diffs {
// 		if diff > 0 {
// 			increase++
// 		} else if diff < 0 {
// 			increase--
// 		} else {
// 			return false
// 		}
// 	}

// 	// 1 1 1 1     ---   4   3 (one is wrong)
// 	// -1 -1 -1 -1 ---  -4  -3  (one is wrong)
// 	return increase == len(diffs) || increase == -len(diffs)
// }

/*
- loop over the list
- diff the right element from the left
- Check if only one of the diffs is different sign from the others
- If more than one is different return false
- If only one is different sign, then add the diff of previous element
 with current and check again
-
*/

// func checkSafe(levels []int) bool {

// 	flag := false

// 	// 5 1 2 3 4
// 	// list[:len(-1)] - list[1:]
// 	// diffs --> -4, 1, 1, 1
// 	// 1 2 5 3 4
// 	// (5-2) + (3-5) = (3 - 2)
// 	// diffs --> 1, 3, -2, 1
// 	// diffs --> 1, 1, 1
// 	// 1 2 5 6 4
// 	// diffs --> 1, 3, 1, -2

// 	//  for index, value...
// 	diffs = levels[:-1] - levels[1:]
// 	checkDiffs(diffs)
// 	for i, _ := range levels {
// 		if i == 0 {
// 			continue
// 		}
// 		diff := levels[i] - levels[i-1]
// 		if !checkDiffs(diff, levels) {
// 			if flag {
// 				return false
// 			}

// 			if !checkDiffs(levels[i]-levels[i-2], levels) {
// 				return false
// 			}

// 			flag = true
// 		}
// 	}
// 	return true
// }
