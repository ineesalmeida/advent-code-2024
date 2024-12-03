package utils

import (
	"fmt"
	"os"
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
