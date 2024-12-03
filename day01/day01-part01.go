package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func part1() {

	// var filePath string
	// fmt.Scan(&filePath)

	filePath := "input-ines.txt"

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

	sort.Ints(locations1)
	sort.Ints(locations2)

	var totalDistance int
	var distance int

	for i := range locations {
		distance = locations1[i] - locations2[i]
		if distance < 0 {
			distance = -distance
		}
		totalDistance += distance
	}
	fmt.Println(totalDistance)
}

// ATTEMPT 1 - failed because inputs differ for people

// const url string = "https://adventofcode.com/2024/day/1/input"

// func main() {

// 	resp, err := http.Get(url)
// 	if err != nil {
// 		fmt.Errorf("Error getting response")
// 		return
// 	}

// 	defer resp.Body.Close()
// 	body, err := io.ReadAll(resp.Body)
// 	stringData := string(body)

// 	// var locations string
// 	// fmt.Scan(&locations)

// 	fmt.Println(stringData)
// }

/*

resp, err := http.Get("...")
check(err) // does some error handling
*/

// fmt.Scan()
