package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func calculateScore(str string) int {
	var res int
	switch str {
	case "A X":
		res = 4
	case "A Y":
		res = 8
	case "A Z":
		res = 3
	case "B X":
		res = 1
	case "B Y":
		res = 5
	case "B Z":
		res = 9
	case "C X":
		res = 7
	case "C Y":
		res = 2
	case "C Z":
		res = 6
	}
	return res
}

func calculateScorePt2(str string) int {
	var res int
	switch str {
	case "A X":
		res = 3
	case "A Y":
		res = 4
	case "A Z":
		res = 8
	case "B X":
		res = 1
	case "B Y":
		res = 5
	case "B Z":
		res = 9
	case "C X":
		res = 2
	case "C Y":
		res = 6
	case "C Z":
		res = 7
	}
	return res
}

func main() {
	file := os.Args[1]
	readFile, err := os.Open(file)
	if err != nil {
		log.Fatalln(err)
	}
	defer readFile.Close()

	scanner := bufio.NewScanner(readFile)

	score := 0
	pt2Score := 0
	for scanner.Scan() {
		score += calculateScore(scanner.Text())
		pt2Score += calculateScorePt2(scanner.Text())
	}

	fmt.Println("Part 1:", score)
	fmt.Println("Part 2:", pt2Score)
}
