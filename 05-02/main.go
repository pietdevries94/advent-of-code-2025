package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

type Range struct {
	start int
	end   int
}

func main() {
	inputFile, err := os.Open("./input.txt")
	if err != nil {
		panic(err)
	}
	inputScanner := bufio.NewScanner(inputFile)

	freshRanges := []Range{}

	// When scanning a file, it does it line by line, making it easy to grabbing the instruction
	for inputScanner.Scan() {
		line := inputScanner.Text()

		if line == "" {
			break
		}

		freshRanges = append(freshRanges, RangeStringToRange(line))
	}

	// print the number of zeroes
	fmt.Println(GetNumberOfUniqueIDs(freshRanges))
}

func RangeStringToRange(rangeStr string) Range {
	parts := strings.Split(rangeStr, "-")
	start, _ := strconv.Atoi(parts[0])
	end, _ := strconv.Atoi(parts[1])
	return Range{
		start,
		end,
	}
}

func GetNumberOfUniqueIDs(freshRanges []Range) int {
	slices.SortFunc(freshRanges, func(a, b Range) int {
		return a.start - b.start
	})

	amountSeen := 0
	minValue := 0

	for _, freshRange := range freshRanges {
		if freshRange.end < minValue {
			continue
		}

		start := max(freshRange.start, minValue)
		amountSeen += freshRange.end - start + 1 // because the ranges are inclusive
		minValue = freshRange.end + 1            // to prevent the end to be counted twice
	}

	return amountSeen
}
