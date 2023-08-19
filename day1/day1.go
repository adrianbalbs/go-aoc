package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main() {
	file := os.Args[1]
	readFile, err := os.Open(file)
	if err != nil {
		fmt.Println(err)
	}

	fileScanner := bufio.NewScanner(readFile)

	var total []int
	curTotal := 0

	for fileScanner.Scan() {
		if fileScanner.Text() == "" {
			total = append(total, curTotal)
			curTotal = 0
		} else {
			num, err := strconv.Atoi(fileScanner.Text())
			if err != nil {
				fmt.Print(err)
			}
			curTotal += num
		}
	}
	sort.Sort(sort.Reverse(sort.IntSlice(total)))
	fmt.Println(total[0])

	readFile.Close()
}
