package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	inputFile, err := os.Open("./input.txt")
	if err != nil {
		panic(err)
	}

	inputScanner := bufio.NewScanner(inputFile)

	var lastLine []int

	for inputScanner.Scan() {
		line := inputScanner.Text()
		newLine := make([]int, len(line))

		if lastLine == nil {
			for i, char := range line {
				if char == 'S' {
					newLine[i] = 1
				}
			}
			lastLine = newLine
			continue
		}

		for i, char := range line {
			if char == '^' {
				newLine[i-1] += lastLine[i]
				newLine[i+1] += lastLine[i]
			} else {
				newLine[i] += lastLine[i]
			}
		}

		lastLine = newLine
	}

	total := 0
	for _, count := range lastLine {
		total += count
	}

	fmt.Println(total)
}
