package main

import (
	"advent-code/aoc2024/utils"
	"fmt"
	"log"
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

func mapAntenas(lines []string) map[string][][2]int {
	antennas := make(map[string][][2]int)
	for i, line := range lines {
		for j, c := range line {
			if c != rune('.') {
				antennas[string(c)] = append(antennas[string(c)], [2]int{i, j})
			}
		}
	}
	return antennas
}

func mergeMaps(map1 map[[2]int]bool, map2 map[[2]int]bool) map[[2]int]bool {
	for k, v := range map2 {
		map1[k] = v
	}
	return map1
}

func countAntiNodes(antennas map[string][][2]int, n int, m int) int {
	// fmt.Println(antennas)
	antinodes := make(map[[2]int]bool)

	for _, freqAntennas := range antennas {
		antinodes = mergeMaps(antinodes, getAntiNodes(freqAntennas, n, m))
	}

	return len(antinodes)
}

func getAntiNodes(nodes [][2]int, n int, m int) map[[2]int]bool {
	// fmt.Println("Checking ", nodes)
	antinodes := make(map[[2]int]bool)

	if len(nodes) == 1 {
		return antinodes
	}

	if len(nodes) == 2 {
		diffX := nodes[1][0] - nodes[0][0]
		diffY := nodes[1][1] - nodes[0][1]

		anti1 := [2]int{nodes[0][0] - diffX, nodes[0][1] - diffY}
		anti2 := [2]int{nodes[1][0] + diffX, nodes[1][1] + diffY}
		// fmt.Println(anti1, anti2)

		if anti1[0] >= 0 && anti1[0] < n && anti1[1] >= 0 && anti1[1] < m {
			antinodes[anti1] = true
		}

		if anti2[0] >= 0 && anti2[0] < n && anti2[1] >= 0 && anti2[1] < m {
			antinodes[anti2] = true
		}

		// fmt.Println(antinodes)
		return antinodes
	}

	// check antinodes that first item does with all others, then check all others recursively
	for _, val := range nodes[1:] {
		mergeMaps(antinodes, getAntiNodes([][2]int{nodes[0], val}, n, m))
	}

	// recursion
	return mergeMaps(antinodes, getAntiNodes(nodes[1:], n, m))
}

func part1(lines []string) any {
	antennas := mapAntenas(lines)
	antinodes := make(map[[2]int]bool)

	for _, freqAntennas := range antennas {
		antinodes = mergeMaps(antinodes, getAntiNodes(freqAntennas, len(lines), len(lines[0])))
	}

	return len(antinodes)
}

func getAntiNodesWithFreq(nodes [][2]int, n int, m int) map[[2]int]bool {
	// fmt.Println("Checking ", nodes, n, m)
	antinodes := make(map[[2]int]bool)

	if len(nodes) == 1 {
		return antinodes
	}

	if len(nodes) == 2 {
		diffX := nodes[1][0] - nodes[0][0]
		diffY := nodes[1][1] - nodes[0][1]

		antinodes[nodes[0]] = true
		antinodes[nodes[1]] = true

		var anti [2]int
		mult := 1
		for true {
			anti = [2]int{nodes[0][0] - diffX*mult, nodes[0][1] - diffY*mult}
			if anti[0] < 0 || anti[0] >= n || anti[1] < 0 || anti[1] >= m {
				break
			}
			antinodes[anti] = true
			mult += 1
		}
		mult = 1
		for true {
			anti = [2]int{nodes[1][0] + diffX*mult, nodes[1][1] + diffY*mult}
			if anti[0] < 0 || anti[0] >= n || anti[1] < 0 || anti[1] >= m {
				break
			}
			antinodes[anti] = true
			mult += 1
		}

		// fmt.Println(antinodes)
		return antinodes
	}

	// check antinodes that first item does with all others, then check all others recursively
	for _, val := range nodes[1:] {
		mergeMaps(antinodes, getAntiNodesWithFreq([][2]int{nodes[0], val}, n, m))
	}

	// recursion
	return mergeMaps(antinodes, getAntiNodesWithFreq(nodes[1:], n, m))
}

func part2(lines []string) any {
	antennas := mapAntenas(lines)
	antinodes := make(map[[2]int]bool)

	for _, freqAntennas := range antennas {
		antinodes = mergeMaps(antinodes, getAntiNodesWithFreq(freqAntennas, len(lines), len(lines[0])))
	}
	// fmt.Println(antinodes)

	return len(antinodes)
}

// ##....#....#
// .#.#....0...
// ..#.#0....#.
// ..##...0....
// ....0....#..
// .#...#A....#
// ...#..#.....
// #....#.#....
// ..#.....A...
// ....#....A..
// .#........#.
// ...#......##
