package main

import (
	"advent-code/aoc2024/utils"
	"fmt"
	"log"
	"sort"
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

	fmt.Printf("[+] day %s\n> part 1: %d (%s)\n> part 2: %s (%s)\n",
		"02",
		part1Answer, p1duration,
		part2Answer, p2duration,
	)
}

var connections = make(map[string][]string)
var triads = make(map[string]bool)

func mapConnections(lines []string) {
	for _, line := range lines {
		computers := strings.Split(line, "-")
		p1, p2 := computers[0], computers[1]
		connections[p1] = append(connections[p1], p2)
		connections[p2] = append(connections[p2], p1)
	}
}

func computeTriads() {
	for p1, p1Connections := range connections {
		// check only the lists that have a connection that starts with t
		if p1[0] != 't' {
			continue
		}

		for _, p2 := range p1Connections {
			p2Connections, _ := connections[p2]
			// p3 is found when a computer in p2Connections exists in p1Connections
			for _, p3 := range p2Connections {
				if utils.Contains(p1Connections, p3) {
					ps := []string{p1, p2, p3}
					sort.Sort(sort.StringSlice(ps))
					triads[strings.Join(ps, ",")] = true
				}
			}
		}
	}
}

func part1(lines []string) any {
	mapConnections(lines)
	computeTriads()
	return len(triads)
}

var groups = make(map[int]map[string]bool)

func part2(lines []string) any {
	// connections and triads are already sorted

	largestGroup := ""
	n := 3
	groups[n] = triads

	for true {
		groups[n+1] = make(map[string]bool)
		for g, _ := range groups[n] {
			group := strings.Split(g, ",")
			GrowGroup(group)
		}
		// no new groups
		if len(groups[n+1]) == 0 {
			break
		}
		n += 1
		// fmt.Println(groups)
	}

	for k, _ := range groups[n] {
		largestGroup = k
		break
	}

	return largestGroup
}

func GrowGroup(group []string) {
	n := len(group)
	for _, p1Connection := range connections[group[0]] {
		newConnection := true
		for _, pXConnextion := range group[1:] {
			if !utils.Contains(connections[pXConnextion], p1Connection) {
				newConnection = false
				break
			}
		}
		if newConnection {
			g := append(group, p1Connection)
			sort.Sort(sort.StringSlice(g))
			groups[n+1][strings.Join(g, ",")] = true
		}
	}
}
