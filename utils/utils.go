package utils

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func FileToLines(filePath string) ([]string, error) {
	data, err := os.ReadFile(filePath)
	if err != nil {
		fmt.Errorf("Error reading file")
		return nil, err
	}

	stringData := string(data)
	lines := strings.Split(stringData, "\n")

	return lines, nil
}

func PrintMatrix(warehouse [][]int) {
	for i := range len(warehouse) {
		for j := range len(warehouse[0]) {
			str := strconv.Itoa((warehouse[i][j]))
			spaces := 5 - len(str)
			fmt.Printf(str + strings.Repeat(" ", spaces))
		}
		fmt.Printf("\n")
	}
}

func PrintMatrixRunes(warehouse [][]rune) {
	for i := range len(warehouse) {
		for j := range len(warehouse[0]) {
			fmt.Printf(string(warehouse[i][j]))
		}
		fmt.Printf("\n")
	}
}

func Abs(integer int) int {
	if integer < 0 {
		return -integer
	}
	return integer
}

func Contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}
