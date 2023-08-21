package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"unicode"
)

func incrTotalPriority(c rune, bottomHalf string, totalPriority *int) bool {
	if strings.ContainsRune(bottomHalf, c) {
		if unicode.IsLower(c) {
			*totalPriority += int((c-'a')%26 + 1)
		} else if unicode.IsUpper(c) {
			*totalPriority += int((c-'A')%26 + 27)
		}
		return true
	}
	return false
}

func part1(input []string) int {
	totalPriority := 0
	for _, str := range input {
		topHalf := str[0 : len(str)/2]
		bottomHalf := str[len(str)/2:]
		for _, c := range topHalf {
			if incrTotalPriority(c, bottomHalf, &totalPriority) {
				break
			}
		}
	}
	return totalPriority
}

func part2Helper(bag1 string, bag2 string, bag3 string) int {
	for _, c := range bag1 {
		if strings.ContainsRune(bag2, c) && strings.ContainsRune(bag3, c) {
			if unicode.IsLower(c) {
				return int((c-'a')%26 + 1)
			} else if unicode.IsUpper(c) {
				return int((c-'A')%26 + 27)
			}
		}
	}
	return 0
}

func part2(input []string) int {
	totalPriority := 0
	var groups []string

	for _, str := range input {
		if len(groups) == 3 {
			bag1, bag2, bag3 := groups[0], groups[1], groups[2]
			totalPriority += part2Helper(bag1, bag2, bag3)
			groups = nil
		}
		groups = append(groups, str)
	}
	return totalPriority
}

func main() {
	readFile, err := os.ReadFile("input")
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(string(readFile))
	input := strings.Split(string(readFile), "\n")
	fmt.Println("Part 1:", part1(input))
	fmt.Println("Part 2:", part2(input))
}
