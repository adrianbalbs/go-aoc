package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func part1Helper(line string) (start1, end1, start2, end2 int) {
	idGroups := strings.Split(line, ",")
	groups := make([]int, 0)
	for _, group := range idGroups {
		groupNums := strings.Split(group, "-")
		for _, num := range groupNums {
			res, err := strconv.Atoi(num)
			if err != nil {
				log.Fatalln(err)
			}
			groups = append(groups, res)
		}
	}
	start1, end1, start2, end2 = groups[0], groups[1], groups[2], groups[3]
	return
}

func isContained(start1, end1, start2, end2 int) int {
	if (start1 >= start2 && end1 <= end2) || (start2 >= start1 && end2 <= end1) {
		return 1
	}
	return 0
}

func part1(input []string) int {
	totalContained := 0
	for _, str := range input {
		if str != "" {
			start1, end1, start2, end2 := part1Helper(str)
			totalContained += isContained(start1, end1, start2, end2)
		}
	}
	return totalContained
}

func main() {
	readFile, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatalln(err)
		return
	}
	input := strings.Split(string(readFile), "\n")
	fmt.Println("Part 1:", part1(input))
}
