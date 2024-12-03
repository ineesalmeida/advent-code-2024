package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	var filePath string
	fmt.Scan(&filePath)

	// filePath := "input-ines.txt"

	fmt.Println("Reading file", filePath)
	dat, err := os.ReadFile(filePath)
	if err != nil {
		fmt.Errorf("Error reading file")
		return
	}

	stringData := string(dat)
	locations := strings.Split(stringData, "\n")

	var locations1 []int
	var locations2 []int

	for _, loc := range locations {
		locSlice := strings.Split(loc, "   ")
		loc1, _ := strconv.Atoi(locSlice[0])
		loc2, _ := strconv.Atoi(locSlice[1])

		locations1 = append(locations1, loc1)
		locations2 = append(locations2, loc2)
	}

	countLocations := countOccurences(locations2)
	var result int

	for _, loc := range locations1 {
		val, ok := countLocations[loc]
		if ok {
			result += loc * val
		}
	}
	fmt.Println(result)
}

// Count the number of occurences of each element within list
func countOccurences(list []int) map[int]int {
	counts := make(map[int]int)
	for _, num := range list {
		counts[num]++
	}
	return counts
}
