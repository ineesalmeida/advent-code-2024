package main

import (
	"advent-code/aoc2024/utils"
	"fmt"
	"log"
	"reflect"
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

func part1(lines []string) any {

	result := 0

	for _, line := range lines {
		secret, _ := strconv.Atoi(line)
		for range 2000 {
			secret = evolveSecret(secret)
		}
		result += secret
	}
	return result
}

func evolveSecret(secret int) int {
	// 1
	secret = prune(mix(secret, secret*64))

	// 2
	secret = prune(mix(secret, secret/32))

	// 3
	secret = prune(mix(secret, secret*2048))

	return secret
}

func mix(secret int, n int) int {
	return secret ^ n
}

func prune(n int) int {
	return n % 16777216
}

func part2(lines []string) any {
	changes := make([][]int, 0)
	prices := make([][]int, 0)

	secrets := make([]int, 0)
	for _, line := range lines {
		s, _ := strconv.Atoi(line)
		secrets = append(secrets, s)
		changes = append(changes, []int{})
		prices = append(prices, []int{s % 10})
	}

	// Create secrets
	for range 2000 {
		for i, secret := range secrets {
			secret = evolveSecret(secret)
			price := secret % 10
			changes[i] = append(changes[i], price-secrets[i]%10)
			prices[i] = append(prices[i], price)
			secrets[i] = secret
		}
	}

	// fmt.Println(secrets, changes, prices)

	customerChangeGroups := make([]map[string]int, len(changes))
	for customer, customerChanges := range changes {
		customerChangeGroups[customer] = make(map[string]int)
		for i := range customerChanges {
			if i >= len(customerChanges)-3 {
				break
			}
			// going front to back to get the earliest prices
			j := len(customerChanges) - i
			customerChangeGroups[customer][SlicetoString(customerChanges[j-4:j])] = prices[customer][j]
		}
	}

	// fmt.Println(customerChangeGroups)

	mostBananas := 0

	// arbitrarily only check starting on the firt 100 (we could check all of them, but it's likely that the sequence exists in the firts 100)
	for startCustomer := range customerChangeGroups[:100] {
		fmt.Println(startCustomer)

		for sequence, price := range customerChangeGroups[startCustomer] {

			p := price
			for _, customerChanges := range customerChangeGroups[startCustomer+1:] {
				aux, _ := customerChanges[sequence]
				p += aux
			}

			if p > mostBananas {
				mostBananas = p
				fmt.Println(sequence, mostBananas)
			}
		}
	}

	return mostBananas
}
func SlicetoString(slice []int) string {
	return strings.Trim(strings.Join(strings.Fields(fmt.Sprint(slice)), ","), "[]")
}

func FindSequence(sequence []int, target []int) int {
	for i := range target {
		if i < 3 {
			continue
		}

		if reflect.DeepEqual(target[i-3:i+1], sequence) {
			return i
		}
	}
	return -1
}

func Sum(slice []int) int {
	s := 0
	for _, v := range slice {
		s += v
	}
	return s
}
