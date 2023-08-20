package main

import (
	"bufio"
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

func main() {
	file := os.Args[1]
	readFile, err := os.Open(file)
	if err != nil {
		log.Fatalln(err)
	}
	defer readFile.Close()

	scanner := bufio.NewScanner(readFile)

	totalPriority := 0
	for scanner.Scan() {
		topHalf := scanner.Text()[0 : len(scanner.Text())/2]
		bottomHalf := scanner.Text()[len(scanner.Text())/2 : len(scanner.Text())]

		for _, c := range topHalf {
			if incrTotalPriority(c, bottomHalf, &totalPriority) {
				break
			}
		}
	}
	fmt.Println("Part 1:", totalPriority)
}
